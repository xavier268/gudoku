package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xavier268/gudoku/sdk"
)

func main() {

	rand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	puzzle, solution := sdk.Build(rand)

	fmt.Println("Puzzle (replace all zeroes !) :")
	puzzle.Dump()
	fmt.Println("Solution :")
	fmt.Println(solution)

}
