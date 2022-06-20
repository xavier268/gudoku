package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
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

func init() {
	flag.IntVar(&flagCount, "count", 10, "number of sudoku to generate")
	flag.IntVar(&flagCount, "c", 10, "shorthand for -count")

	flag.IntVar(&flagMaxDifficulty, "difficulty", 9*9, "maximum allowed difficulty (number of blank values)")
	flag.IntVar(&flagMaxDifficulty, "d", 9*9, "shorthand for -difficulty")

	flag.DurationVar(&flagMaxTime, "timeout", 0, "maximum time allocated (0 for no limit).\nDuration string is a signed sequence of decimal numbers with optional fraction and unit suffix, like '100ms', '2.3h' or '4h35m'.")
	flag.DurationVar(&flagMaxTime, "t", 0, "shorthand for -timeout")

	flag.BoolVar(&flagWithSolutions, "solutions", false, "if true, solutions are also generated")
	flag.BoolVar(&flagWithSolutions, "s", false, "shorthand for -solutions")

	flag.StringVar(&flagOutputFile, "output", "sudokus.txt", "file name to output solutions")
	flag.StringVar(&flagOutputFile, "o", "sudokus.txt", "shorthand for -output")

}

func main() {

	flag.Parse()
	var ctx context.Context
	var cancel context.CancelFunc
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	sh := sdk.NewShuffler(rd)

	if flagMaxTime > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), flagMaxTime)
		defer cancel()
	} else {
		ctx = context.Background()
	}

	of, err := os.Create(flagOutputFile)
	if err != nil {
		panic(err)
	}
	defer of.Close()

	fmt.Fprintln(of, "Generated sudoku - ", time.Now())

	for i := 1; i <= flagCount; i++ {

		select {

		case <-ctx.Done():
			fmt.Fprintf(of, "\nTime out after %v \n", flagMaxTime)
			return

		default:
			p, s := sdk.BuildRandom(rd, 9*9-flagMaxDifficulty)
			sh.Shuffle(p, s)
			sh.Reset()

			fmt.Fprintln(of, "\nPuzzle N° ", i, "\n", p.StringDot())
			if flagWithSolutions {
				fmt.Fprintln(of, "\nSolution to puzzle N° ", i, "\n", s.String())
			}
		}
	}
}
