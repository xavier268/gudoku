package sdk

import (
	"fmt"
	"io"
	"strings"
)

type Table struct {
	// table of values
	tab [9 * 9]int
	n   int // nb of non zero values
}

func NewTable() *Table {
	var t Table

	return &t
}

func (t *Table) Clone() *Table {
	var tt Table
	for a := 0; a < 9*9; a++ {
		tt.Set(a, t.Get(a))
	}
	return &tt
}

func (t *Table) Equal(tt *Table) bool {
	for a := 0; a < 9*9; a++ {
		if t.Get(a) != tt.Get(a) {
			return false
		}
	}
	return true
}

func (t *Table) Get(a int) int {
	return t.tab[a]
}

func (t *Table) Set(a int, val int) {
	old := t.Get(a)
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
	t.tab[a] = val
}

// Valid checks if current table is a valid, possibly incomplete, sudoku.
func (t *Table) Valid() bool {

	var m int
	for i := range indx { // The various groups to test
		m = 0 // mask
		for j := 0; j < 9; j++ {
			v := t.Get(indx[i][j])
			if v > 0 { // only test for non empty positions
				b := 1 << v
				if b&m != 0 {
					//fmt.Printf("Invalid grid for %d data, rank %d, value %d\n", i, j, v) // debug
					return false
				} else {
					m = m | b
				}
			}
		}

	}

	return true
}

// To string
// Output can be scanned back using Scan.
func (t *Table) String() string {
	var sb strings.Builder
	h := "---------------------------------"
	for a := 0; a < 9*9; a++ {
		if a%27 == 0 {
			fmt.Fprintln(&sb, h)
		}
		fmt.Fprintf(&sb, "%3d", t.Get(a))
		switch a % 9 {
		case 2, 5:
			fmt.Fprint(&sb, " | ")
		case 8:
			fmt.Fprintln(&sb)
		default:
		}
	}
	fmt.Fprintln(&sb, h)
	return sb.String()
}

// Same as String, but 0 are printed a '.'
// Output cannot be scanned back using Scan.
func (t *Table) StringDot() string {
	var sb strings.Builder
	h := "---------------------------------"
	for a := 0; a < 9*9; a++ {
		if a%27 == 0 {
			fmt.Fprintln(&sb, h)
		}
		if t.Get(a) == 0 {
			fmt.Fprint(&sb, "  .")
		} else {
			fmt.Fprintf(&sb, "%3d", t.Get(a))
		}
		switch a % 9 {
		case 2, 5:
			fmt.Fprint(&sb, " | ")
		case 8:
			fmt.Fprintln(&sb)
		default:
		}
	}
	fmt.Fprintln(&sb, h)
	return sb.String()
}

// Print in a condensed format.
// Cannot be scanned back by Scan.
func (t *Table) PrintCondensed() {
	fmt.Println()
	for a := 0; a < 9*9; a++ {
		if t.Get(a) == 0 {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", t.Get(a))
		}
		if a%9 == 8 {
			fmt.Println()
		}
	}
}

// Print in standard format.
// Output can be scanned back using Scan.
func (t *Table) Print() {
	fmt.Print("\n", t.String())
}

// Same as Print, but 0 are printed as '.'.
// Output cannot be scanned back using Scan.
func (t *Table) PrintDot() {
	fmt.Print("\n", t.StringDot())
}

// Same as Print, but with additionnal output information.
func (t *Table) Dump(message ...string) {
	fmt.Print("\n\t", message)
	t.Print()
	fmt.Printf("There are %d non-zero values and %d zero values\n", t.n, t.Difficulty())
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
	for i := 0; i < 9*9; i++ {

		// remove non digits
		for len(buf) != 0 && (buf[0] < '0' || buf[0] > '9') {
			buf = buf[1:]
		}
		// stop if no more digits
		if len(buf) == 0 {
			return
		}
		// read and use one digit
		t.Set(i, int(buf[0]-'0'))
		buf = buf[1:]
	}
}

// Difficulty returns the number of blank (zero) values.
func (t *Table) Difficulty() int {
	return 9*9 - t.n
}
