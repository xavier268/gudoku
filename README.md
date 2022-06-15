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

    // random source
	rand := rand.New(rand.NewSource(time.Now().UnixMicro()))

    // build a puzzle and its solution
	puzzle, solution := sdk.Build(rand)

    // display both
	fmt.Println("Puzzle (replace all zeroes !) :")
	puzzle.Dump()
	fmt.Println("Solution :")
	fmt.Println(solution)

}


```
