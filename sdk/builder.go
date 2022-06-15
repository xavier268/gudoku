package sdk

import (
	"fmt"
	"math/rand"
)

// BuildRandom build a puzzle and its solution from a random seed.
func BuildRandom(rd *rand.Rand) (puzzle, solution *Table) {

	t := NewTable()

	// initialize random table ...
	for i := 1; i < 10; i++ {
		a := rd.Intn(9 * 9)
		t.Set(a, i)
	}
	if !t.Solve() {
		t.Dump()
		fmt.Println("Unsolvable random starting position, trying another one ...")
		return BuildRandom(rd)
	}

	solution = t.Clone()
	puzzle = BuildFromSolution(solution)
	return puzzle, solution

}

// BuildFromSolution attempts to find a minimal puzzle for the given solution
func BuildFromSolution(solution *Table) (puzzle *Table) {
	if solution.n != 9*9 || !solution.Valid() {
		solution.Dump()
		panic("Attempting to build from a solution that is not a solution ...")
	}

	puzzle = solution.Clone() // best puzzle so far
	try := puzzle.Clone()     // try is a temporary table, modified when solving

	for a := 0; a < 9*9; a++ {

		if try.Get(a) == 0 {
			continue // ignore zero values ..
		}
		// try erasing a value
		try.Set(a, 0)

		if try.SolveBack() {
			if try.Equal(solution) { // still unique ..
				puzzle.Set(a, 0) // update puzzle
				//fmt.Printf("%3d", puzzle.n) // show progress, DEBUG
				try = puzzle.Clone() // try from there now
				a = 0                //  reset loop from scratch
				continue
			} else { // not unique anymore, too far : restore and continue exploring loop
				try = puzzle.Clone() // do not update puzzle ...
				continue
			}
		} else { // could not solve
			panic("internal logic error 2")
		}
	}
	return puzzle
}
