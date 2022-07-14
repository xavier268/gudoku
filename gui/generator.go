package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xavier268/gudoku/sdk"
)

type pair struct {
	puzzle, solution *sdk.Table
}

// gen will continuously maintain a couple of precomputed puzzles, for better reactivity when new is called.
func gen(rand *rand.Rand, ch chan pair) {

	for {
		ti := time.Now()
		puzzle, solution := sdk.BuildRandom(rand, 81-flagMaxDifficulty)
		s := sdk.NewShuffler(rand)
		s.Shuffle(puzzle, solution)
		dur := time.Since(ti)
		ch <- pair{puzzle, solution}
		if flagVerbose {
			fmt.Println("Precomputed a new puzzle in ", dur)
		}
	}

}
