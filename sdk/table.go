package sdk

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type Table struct {
	// table of values
	tab [3][3][3][3]int
	n   int // nb of non zero values
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
	old := t.Get(a, b, c, d)
	if old == val {
		return
	}
	if old == 0 {
		t.n++
	} else {
		if val == 0 {
			t.n--
		}
		if old == val {
			return
		}
	}

	if val < 0 || val > 9 {
		fmt.Println("Trying to set invalid value : ", val)
		panic("trying to set invalid value")
	}
	t.tab[a][b][c][d] = val
}

// Valid checks if current table is a valid, possibly incomplete, sudoku.
func (t *Table) Valid() bool {
	var ml, mc, mg int // mask lines, col, group
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ml, mc, mg = 0, 0, 0
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					// lines
					v := t.Get(i, j, k, l)
					if v > 0 {
						b := 1 << v
						if b&ml != 0 {
							return false
						} else {
							ml = ml | b
						}
					}
					// col
					v = t.Get(k, l, i, j)
					if v > 0 {
						b := 1 << v
						if b&mc != 0 {
							return false
						} else {
							mc = mc | b
						}
					}
					// groups
					v = t.Get(i, k, j, l)
					if v > 0 {
						b := 1 << v
						if b&mg != 0 {
							return false
						} else {
							mg = mg | b
						}
					}

				}
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

// To string
func (t *Table) String() string {
	return t.WalkString(
		func(i, j, k, l int) string {
			return fmt.Sprintf("%3d", t.Get(i, j, k, l))
		})
}

// WalkString generate a string while walking the table, adding separators and newlines.
func (t *Table) WalkString(wf func(i, j, k, l int) string) string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "-----------------------------------------------------------------------")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {

					fmt.Fprint(&sb, wf(i, j, k, l))
				}
				if k != 2 {
					fmt.Fprint(&sb, " | ")
				}

			}
			fmt.Fprintln(&sb)
		}
		fmt.Fprintln(&sb, "-----------------------------------------------------------------------")
	}
	return sb.String()
}

func (t *Table) Print() {
	fmt.Print(t.String())
}

func (t *Table) Dump() {
	t.Print()
	fmt.Printf("There are %d non-zero values and %d zero values\n", t.n, 9*9-t.n)
	if t.Valid() {
		fmt.Println("The grid is VALID")
	} else {
		fmt.Println("The grid is INVALID")
	}
}

// Scan from io.Reader, replacing current table content.
func (t *Table) Scan(r io.Reader) {
	buf, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					// remove non digits
					for len(buf) != 0 && (buf[0] < '0' || buf[0] > '9') {
						buf = buf[1:]
					}
					// stop if no more digits
					if len(buf) == 0 {
						return
					}
					// read and use one digit
					t.Set(i, j, k, l, int(buf[0]-'0'))
					buf = buf[1:]
				}
			}
		}
	}
}

// RandValue provides a random value between 1 and 9 included.
func RandValue() int {
	return rand.Intn(8) + 1
}
