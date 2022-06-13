package sdk

import (
	"math/rand"
)

// Build build a puzzle and its solution from a random seed.
func Build(rand *rand.Rand) (puzzle, solution *Table) {

	t := NewTable()

	for i := 1; i < 10; i++ {
		a, b, c, d := rand.Intn(3), rand.Intn(3), rand.Intn(3), rand.Intn(3)
		t.Set(a, b, c, d, i)
	}
	if !t.Solve() {
		panic("internal logic error")
	}
	solution = t.Clone()
	puzzle = t.Clone()
	//fmt.Println(solution)

	// remove positions, starting from a random position, until the puzzle starts to have multiple solutions
	a, b, c, d := rand.Intn(3), rand.Intn(3), rand.Intn(3), rand.Intn(3)
	for {
		try := puzzle.Clone()

		for a, b, c, d = inc(a, b, c, d); try.Get(a, b, c, d) == 0; a, b, c, d = inc(a, b, c, d) {
		} // loop until non zero reached
		try.Set(a, b, c, d, 0)

		//try.Dump() // debug

		if try.SolveBack() {
			if try.Equal(solution) { // still unique ..
				puzzle.Set(a, b, c, d, 0) // register  puzzle modification
				continue
			} else { // not unique anymore
				return puzzle, solution
			}

		} else { // could not solve
			panic("internal logic error 2")
		}

	}

}

// incremnt the position, infinitely ...
func inc(a, b, c, d int) (int, int, int, int) {
	a++
	if a == 3 {
		a = 0
		b++
		if b == 3 {
			b = 0
			c++
			if c == 3 {
				c = 0
				d++
				if d == 3 {
					d = 0
				}
			}

		}
	}
	return a, b, c, d

}
