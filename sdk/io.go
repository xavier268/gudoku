package sdk

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

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
	fmt.Println(t)
}

// Scan from io.Reader, replacing current table content.
func (t *Table) Scan(r io.Reader) {
	t = NewTable()
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
