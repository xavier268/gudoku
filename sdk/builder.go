package sdk

import (
	"fmt"
	"math/rand"
)

// Build build a puzzle and its solution from a random seed.
func Build(rand *rand.Rand) (puzzle, solution *Table) {

	t := NewTable()

	for i := 1; i < 10; i++ {
		a := rand.Intn(9 * 9)
		t.Set(a, i)
	}
	if !t.Solve() {
		t.Dump()
		fmt.Println("Unsolvable random starting position, trying another one ...")
		return Build(rand)
	}
	solution = t.Clone()
	puzzle = t.Clone()
	//fmt.Println(solution)

	// remove positions, starting from a random position, until the puzzle starts to have multiple solutions
	a := rand.Intn(9 * 9)
	for {
		try := puzzle.Clone()

		for a = (a + 1) % 81; try.Get(a) == 0; a = (a + 1) % 81 {
		} // loop until non zero reached
		try.Set(a, 0)

		//try.Dump() // debug

		if try.SolveBack() {
			if try.Equal(solution) { // still unique ..
				puzzle.Set(a, 0) // register  puzzle modification
				continue
			} else { // not unique anymore
				return puzzle, solution
			}

		} else { // could not solve
			panic("internal logic error 2")
		}

	}

}
