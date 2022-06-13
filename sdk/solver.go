package sdk

import (
	"context"
	"time"
)

// Solve attempts to solve provided table, returning true on success, or false if failed.
// The table contains a solution, if found.
func (t *Table) Solve() bool {

	// fmt.Println("Solving for ", t.n)
	// t.Dump()

	if t.Valid() && t.n == 9*9 {
		return true // done !
	}

	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			for c := 0; c < 3; c++ {
				for d := 0; d < 3; d++ {
					v := t.Get(a, b, c, d)
					if v == 0 {
						for i := 1; i <= 9; i++ {
							t.Set(a, b, c, d, i)
							if t.Valid() && t.Solve() {
								return true
							} // else iterate values
						}
						t.Set(a, b, c, d, 0) // restore state before returning
						return false
					}

				}
			}
		}
	}
	return false
}

// SolveBack is the same as Solve, but working from the other end of the solution space.
// Used to ensure unicity of a solution.
func (t *Table) SolveBack() bool {

	if t.Valid() && t.n == 9*9 {
		return true // done !
	}

	for a := 2; a >= 0; a-- {
		for b := 2; b >= 0; b-- {
			for c := 2; c >= 0; c-- {
				for d := 2; d >= 0; d-- {
					v := t.Get(a, b, c, d)
					if v == 0 {
						for i := 9; i > 0; i-- {
							t.Set(a, b, c, d, i)
							if t.Valid() && t.SolveBack() {
								return true
							} // else iterate values
						}
						t.Set(a, b, c, d, 0) // restore state before returning
						return false
					}

				}
			}
		}
	}
	return false
}

// Solven attempts to solve provided table with the provided context, sending the solution to the channel.
func (t *Table) Solven(ctx context.Context, out chan *Table) {

	if ctx.Err() != nil {
		return
	}

	if t.Valid() && t.n == 9*9 {
		out <- t.Clone()
	}

	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			for c := 0; c < 3; c++ {
				for d := 0; d < 3; d++ {
					v := t.Get(a, b, c, d)
					if v == 0 {
						for i := 1; i <= 9; i++ {
							if ctx.Err() != nil {
								return
							}
							t.Set(a, b, c, d, i)
							if t.Valid() {
								t.Solven(ctx, out)
							} // else iterate values
						}
						t.Set(a, b, c, d, 0) // restore state before returning
						return
					}

				}
			}
		}
	}
}

// SolveSlice returns a slice of solutions that could be generarated within the specified duration.
func (t *Table) SolveSlice(duration time.Duration) []*Table {

	var sol []*Table
	out := make(chan *Table, 10)
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	go t.Solven(ctx, out)

	for {
		select {
		case ttt := <-out:
			sol = append(sol, ttt)
		case <-ctx.Done():
			return sol
		}
	}
}
