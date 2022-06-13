package sdk

import (
	"fmt"
)

type Table struct {
	// table of values
	tab [3][3][3][3]int
}

func NewTable() *Table {
	var t Table
	return &t
}

func (t *Table) Clone() *Table {
	var tt Table
	tt.Walk(func(i, j, k, l int) bool {
		tt.Set(i, j, k, l, t.Get(i, j, k, l))
		return false
	})
	return &tt
}

func (t *Table) Equal(tt *Table) bool {
	return !t.Walk(
		func(a, b, c, d int) bool {
			return t.Get(a, b, c, d) != tt.Get(a, b, c, d)
		})
}

func (t *Table) Get(a, b, c, d int) int {
	return t.tab[a][b][c][d]
}

func (t *Table) Set(a, b, c, d int, val int) {
	if val < 0 || val > 9 {
		fmt.Println("Trying to set invalid value : ", val)
		panic("trying to set invalid value")
	}
	t.tab[a][b][c][d] = val
}

// TODO - probably wrong computation ?
func (t *Table) Free(a, b, c, d int, val int) bool {
	if t.Get(a, b, c, d) != 0 {
		return false
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.Get(a, b, i, j) == val ||
				t.Get(a, i, c, j) == val ||
				t.Get(a, i, j, d) == val ||
				t.Get(i, b, c, j) == val ||
				t.Get(i, b, j, d) == val ||
				t.Get(i, j, c, d) == val {
				return false
			}
		}
	}
	return true
}

// Walk the table, applying the function.
// If function returns true, the walk is stopped.
func (t *Table) Walk(wf func(a, b, c, d int) (stop bool)) (stopped bool) {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if wf(i, j, k, l) {
						return true
					}
				}
			}
		}
	}
	return false

}
