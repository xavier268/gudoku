package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/xavier268/gudoku/sdk"
)

/*
Command line flags :
-h
	display help
-c int
	shorthand for -count (default 10)
-count int
	number of sudoku to generate (default 10)
-d int
	shorthand for -difficulty (default 81)
-difficulty int
	maximum allowed difficulty (number of blank values) (default 81)
-o string
	shorthand for -output (default "sudokus.txt")
-output string
	file name to output solutions (default "sudokus.txt")
-s
	shorthand for -solutions
-solutions
	if true, solutions are also generated
-t duration
	shorthand for -timeout
-timeout duration
	maximum time allocated (0 for no limit).
	Duration string is a signed sequence of decimal numbers with optional fraction and unit suffix, like '100ms', '2.3h' or '4h35m'.
*/

var flagCount int
var flagMaxDifficulty int
var flagMaxTime time.Duration
var flagWithSolutions bool
var flagOutputFile string
var flagVerbose bool

func init() {
	flag.IntVar(&flagCount, "count", 10, "number of sudoku to generate")
	flag.IntVar(&flagCount, "c", 10, "shorthand for -count")

	flag.IntVar(&flagMaxDifficulty, "difficulty", 9*9, "maximum allowed difficulty (number of blank values)")
	flag.IntVar(&flagMaxDifficulty, "d", 9*9, "shorthand for -difficulty")

	flag.DurationVar(&flagMaxTime, "timeout", 0, "maximum time allocated (0 for no limit).\nDuration string is a signed sequence of decimal numbers with optional fraction and unit suffix, like '100ms', '2.3h' or '4h35m'.")
	flag.DurationVar(&flagMaxTime, "t", 0, "shorthand for -timeout")

	flag.BoolVar(&flagWithSolutions, "solutions", false, "if true, solutions are also generated")
	flag.BoolVar(&flagWithSolutions, "s", false, "shorthand for -solutions")

	flag.BoolVar(&flagVerbose, "v", false, "print more detailed (verbose) information ")

	flag.StringVar(&flagOutputFile, "output", "sudokus.txt", "file name to output solutions")
	flag.StringVar(&flagOutputFile, "o", "sudokus.txt", "shorthand for -output")

}

func main() {

	// set up
	flag.Parse()
	var ctx context.Context
	var cancel context.CancelFunc
	out := make(chan pair, 10)
	of, err := os.Create(flagOutputFile)
	if err != nil {
		panic(err)
	}
	defer of.Close()
	if flagMaxTime > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), flagMaxTime)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	fmt.Fprintln(of, "Auto-generated sudokus\n", time.Now())
	fmt.Fprintf(of, "\nCLI Version : %s\nSdk Version : %s\n%s\n", VERSION, sdk.VERSION, sdk.COPYRIGHT)

	if flagVerbose {

		verb := "\n"
		verb += fmt.Sprintf("CLI Version : %s\nSdk Version : %s\n%s\n", VERSION, sdk.VERSION, sdk.COPYRIGHT)
		verb += fmt.Sprintf("\tCount         \t%d\n", flagCount)
		verb += fmt.Sprintf("\tDifficulty    \t%d\n", flagMaxDifficulty)
		verb += fmt.Sprintf("\tTimeout       \t%v\n", flagMaxTime)
		verb += fmt.Sprintf("\tWith solutions\t%v\n", flagWithSolutions)
		verb += fmt.Sprintf("\tVerbose       \t%v\n", flagVerbose)
		verb += fmt.Sprintf("\tSaved file    \t%s\n", flagOutputFile)

		fmt.Println(verb)
		fmt.Fprintln(of, verb)
	}

	// launch generating goroutines
	fmt.Printf("Launching %d goroutines\n", runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go genTab(ctx, out)
	}

	// read up to the requestestd count or timeout ...
	for i := 1; i <= flagCount; i++ {

		ti := time.Now()

		select {

		case <-ctx.Done():
			fmt.Fprintf(of, "\n... Time out after %v \n", flagMaxTime)
			fmt.Printf("\n... Time out after %v \n", flagMaxTime)
			return

		case pp := <-out:

			fmt.Fprintln(of, "\nPuzzle N° ", i, "\n", pp.pzl.StringDot())

			if flagVerbose {
				verb := fmt.Sprintf("There are %d values and %d blanks,\tGenerated in %v\n", 9*9-pp.pzl.Difficulty(), pp.pzl.Difficulty(), time.Since(ti))
				fmt.Fprint(of, verb)
				fmt.Printf("%d\t%s", i, verb)
			}
			if flagWithSolutions {
				fmt.Fprintln(of, "\nSolution to puzzle N° ", i, "\n", pp.sol.String())
			}
		}

	}
}

type pair struct{ pzl, sol *sdk.Table }

// generate a couple of puzzle and solution and send it into the out channel
func genTab(ctx context.Context, out chan pair) {

	rd := rand.New(rand.NewSource(time.Now().Unix() + int64(rand.Uint64()))) // ensure each goroutine has a different seed
	sh := sdk.NewShuffler(rd)

	if flagVerbose {
		fmt.Println("===== starting table generator goroutine =========")
	}

	for {
		select {
		case <-ctx.Done():
			if flagVerbose {
				fmt.Println("===== stopping table generator goroutine =========")
			}
			return
		default:
			p, s := sdk.BuildRandom(rd, 9*9-flagMaxDifficulty)
			sh.Shuffle(p, s)
			sh.Reset()
			out <- pair{p, s}
		}
	}

}
