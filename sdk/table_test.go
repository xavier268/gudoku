package sdk

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestScanVisual(t *testing.T) {
	s := " 123456789 012345678 987654321 050607080 666666666 888888888 999999999 000000000 123123123"
	var r io.Reader = strings.NewReader(s)
	t1 := Scan(r)
	t1.Print()
	s = t1.Clone().String()
	t2 := Scan(strings.NewReader(s))
	// t2.Print()
	if !t2.Equal(t1) {
		t.Fatal("rescanning printed table failed")
	}
}

func TestScanAuto(t *testing.T) {
	data := []string{
		" 123456789 012345678 987654321 050607080 666666666 888888888 999999999 000000000 123123123",
		"123456789012345678987654321050607080666666666888888888999999999000000000123123123",
		"12345678901234567898765432105060sdf7080666666666888888888kjhkjh999999999000000000123123123",
		"12345sdf\n\n6789012345678987654321sdf  0506070806666666668888\tùùé88888999999999000000000123123123",
	}

	tt := make([](Table), len(data))
	for i := range data {
		tt[i] = Scan(strings.NewReader(data[i]))
		t2 := Scan(strings.NewReader(tt[i].String()))
		if !t2.Equal(tt[i]) {
			t.Fatalf("Rescannng failed for i = %d", i)
		}
	}
	for i := range data {
		if !tt[i].Equal(tt[0]) {
			t.Fatalf("Comparison failed for i = %d", i)
		}
	}
}

func TestMireVisual(_ *testing.T) {
	fmt.Println("Printing coordinates system\n-----------------------------------------------------------")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					fmt.Printf(" %1d%1d%1d%1d ", i, j, k, l)
				}
				if k != 2 {
					fmt.Print(" | ")
				}
			}
			fmt.Println()
		}
		fmt.Println("-----------------------------------------------------------")
	}

}

func TestScanAuto2(t *testing.T) {

	data2 := []string{
		"kjhsç_èkjh65q654\n6dfç_è65465\t4 n",
		"654qs   qsldj   ljdfg5443 sfd54 354sfgb35 ",
	}

	tt := make([](Table), len(data2))
	for i := range data2 {
		tt[i] = Scan(strings.NewReader(data2[i]))
		t2 := Scan(strings.NewReader(tt[i].String()))
		if !t2.Equal(tt[i]) {
			t.Fatalf("Rescannng failed for i = %d", i)
		}
	}
}

func TestFree1(t *testing.T) {

	pos := [...]int{1, 2, 0, 2} // position to check

	fmt.Printf("Position marker - %v\n", pos)

	data := []struct {
		p [4]int  // position to test
		s string  // table content
		v [9]bool // map of free values at position, 1 to 9 (indexed as 0 to 8)
	}{
		{
			[...]int{1, 2, 0, 2},
			" ",
			[...]bool{true, true, true, true, true, true, true, true, true},
		},
		{
			[...]int{1, 2, 0, 2},
			"123456789 456789123 ",
			[...]bool{true, true, false, true, true, false, true, true, true},
		},
		{
			[...]int{1, 2, 0, 2},
			"123456789 456789123 000000000 456000000 000000000 000000891",
			[...]bool{false, true, false, false, false, false, true, false, false},
		},
		{
			[...]int{1, 2, 0, 2},
			"123456789 456789123 000000000 456000000 000000000 000000891 002000000",
			[...]bool{false, false, false, false, false, false, true, false, false},
		},
	}

	for _, d := range data {

		tt := Scan(strings.NewReader(d.s))
		for i := 1; i <= 9; i++ {
			got := tt.Free(pos[0], pos[1], pos[2], pos[3], i)
			want := d.v[i-1]
			if got != want {
				tt.Print()
				t.Fatalf("Pos %v, Value %d : got %v but want %v", pos, i, got, want)
			}

		}
	}
}
