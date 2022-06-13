package sdk

import (
	"context"
	"time"
)

// Solve attempts to solve provided table, returning true on success, or false if failed.
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

// Solves attempts to solve provided table for up to n solutions, sending the solution to the channel.
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
