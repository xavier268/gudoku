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
	t1 := NewTable()
	t1.Scan(r)
	t1.Dump()
	s = t1.Clone().String()
	t2 := NewTable()
	t2.Scan(strings.NewReader(s))
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

	tt := make([](*Table), len(data))
	t2 := NewTable()
	for i := range data {
		tt[i] = NewTable()
		tt[i].Scan(strings.NewReader(data[i]))
		t2.Scan(strings.NewReader(tt[i].String()))
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
	h := "---------------------------------"
	for a := 0; a < 9*9; a++ {
		if a%27 == 0 {
			fmt.Println(h)
		}
		fmt.Printf("%3d", a)
		switch a % 9 {
		case 2, 5:
			fmt.Print(" | ")
		case 8:
			fmt.Println()
		default:
		}
	}
	fmt.Println(h)

}

func TestScanAuto2(t *testing.T) {

	data2 := []string{
		"kjhsç_èkjh65q654\n6dfç_è65465\t4 n",
		"654qs   qsldj   ljdfg5443 sfd54 354sfgb35 ",
	}

	tt := make([](*Table), len(data2))
	t2 := NewTable()
	for i := range data2 {
		tt[i] = NewTable()
		tt[i].Scan(strings.NewReader(data2[i]))
		t2.Scan(strings.NewReader(tt[i].String()))
		if !t2.Equal(tt[i]) {
			t.Fatalf("Rescannng failed for i = %d", i)
		}
	}
}

func TestValid(t *testing.T) {

	data := []struct {
		s string // table content
		v bool   // valid or not
	}{
		{"", true},
		{"123789456 456123789 789456123 231897564 564231897 897564231 312978645 645312978 978645312", true},
		{"123789156 456123789 789456123 231897564 564231897 897564231 312978645 645312978 978645312", false},
		{"123789456 456123789 789456123 231897564 564231897 897564231 312879645 645312978 978645312", false},
		{"123789456 456123789 789456123 431897562 564231897 897564231 312978645 645312978 978645312", false},
	}

	tt := NewTable()
	for _, d := range data {
		tt.Scan(strings.NewReader(d.s))
		tt.PrintCondensed()
		got := tt.Valid()
		want := d.v
		if got != want {
			tt.Dump()
			t.Fatalf("Expected validity : %v, but actual validity : %v\n%s\n", want, got, d.s)
		}

	}
}
