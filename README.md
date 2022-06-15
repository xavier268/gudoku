# gudoku
Sudoku builder/solver

# How to use :

```

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xavier268/gudoku/sdk"
)

func main() {

	rand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	puzzle, solution := sdk.BuildRandom(rand)

	puzzle.Dump("Puzzle :")
	solution.Dump("Solution :")

	fmt.Println("Shuffling ...")
	s := sdk.NewShuffler(rand)
	s.Shuffle(puzzle, solution)

	puzzle.Dump("Puzzle shuffled:")
	solution.Dump("Solution shuffled :")

}



```
