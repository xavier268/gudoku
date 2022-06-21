[![Go Reference](https://pkg.go.dev/badge/github.com/xavier268/gudoku.svg)](https://pkg.go.dev/github.com/xavier268/gudoku)
[![Go Report Card](https://goreportcard.com/badge/github.com/xavier268/gudoku)](https://goreportcard.com/report/github.com/xavier268/gudoku)

# gudoku
Sudoku builder/solver

# How to use from the command line

```
# to generate 5 puzzles, with solutions, with a low difficulty of 15
go run . -s -c=5 -d=15 -o=easy.txt

# to generate 12 puzzle, without solution, with the maximum difficulty, in less than 10 seconds
go run . -c=12 -d=100 -o=difficult.txt -t=10s

# to get help
go run . -h

```
# How to use the library directly :

```

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xavier268/gudoku/sdk"
)

func main() {

	rand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	puzzle, solution := sdk.BuildRandom(rand, 0)

	puzzle.Dump("Puzzle :")
	solution.Dump("Solution :")

	fmt.Println("Shuffling ...")
	s := sdk.NewShuffler(rand)
	s.Shuffle(puzzle, solution)

	puzzle.Dump("Puzzle shuffled:")
	solution.Dump("Solution shuffled :")

}



```
