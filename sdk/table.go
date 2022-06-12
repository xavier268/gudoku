package sdk

import (
	"fmt"
)

type Table interface {
	// Get the value at coordinates
	Get(a, b, c, d int) int
	// Set val to the coordinates
	Set(a, b, c, dint, val int)
	// Clone a table
	Clone() Table
	// Free checks if the given coordinates is empty AND could receive value val
	Free(a, b, c, d int, val int) bool
	Print()
	String() string
	Equal(t Table) bool
}

type table struct {
	// table of values
	tab [3][3][3][3]int
}

func NewTable() Table {
	var t table
	return &t
}

func (t *table) Clone() Table {
	var tt table
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					tt.Set(i, j, k, l, t.Get(i, j, k, l))
				}
			}
		}
	}
	return &tt
}

func (t *table) Equal(tt Table) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if tt.Get(i, j, k, l) != t.Get(i, j, k, l) {
						return false
					}
				}
			}
		}
	}
	return true
}

func (t *table) Get(a, b, c, d int) int {
	return t.tab[a][b][c][d]
}

func (t *table) Set(a, b, c, d int, val int) {
	if val < 0 || val > 9 {
		fmt.Println("Trying to set invalid value : ", val)
		panic("trying to set invalid value")
	}
	t.tab[a][b][c][d] = val
}

func (t *table) Free(a, b, c, d int, val int) bool {
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
