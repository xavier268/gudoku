package sdk

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

// To string
func (t table) String() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "-----------------------------------")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {

					fmt.Fprintf(&sb, "%3d", t.Get(i, j, k, l))
				}
				if k != 2 {
					fmt.Fprint(&sb, "  |")
				}

			}
			fmt.Fprintln(&sb)
		}
		fmt.Fprintln(&sb, "-----------------------------------")
	}
	return sb.String()
}

func (t *table) Print() {
	fmt.Println(t)
}

// Scan from io.Reader
func Scan(r io.Reader) Table {
	t := NewTable()
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
						return t
					}
					// read and use one digit
					t.Set(i, j, k, l, int(buf[0]-'0'))
					buf = buf[1:]
				}
			}
		}
	}
	return t
}

// RandValue provides a random value between 1 and 9 included.
func RandValue() int {
	return rand.Intn(8) + 1
}
