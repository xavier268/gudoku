package sdk

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

var real1 string = "100000006 006020700 789450103 000807004 000030000 090004201 312970040 040012078 908000000" // single solution
var real2 string = "100000006 006020700 789450103 000807004 000030000 090004201 312970040 040010078 908000000" // two solutions

func TestSolverEmpty(t *testing.T) {

	tt := NewTable()
	if tt.Solve() {
		tt.Dump()
		return
	}
	t.Fatalf("Could not find a solution even starting with empty table !")
}

func TestSolverReal(t *testing.T) {
	tt := NewTable()
	tt.Scan(strings.NewReader(real1))
	if tt.Solve() {
		tt.Dump()
		return
	}
	t.Fatalf("Could not find a solution to the real sudoku !")
}

func TestSolvenReal(t *testing.T) {
	tt := NewTable()
	tt.Scan(strings.NewReader(real1))
	out := make(chan *Table, 10) // solution tables
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	go tt.Solven(ctx, out)
	for {
		select {
		case ttt := <-out:
			ttt.Dump()
		case <-ctx.Done():
			return
		}
	}
}

func TestSolvenReal2(t *testing.T) {
	tt := NewTable()
	tt.Scan(strings.NewReader(real2))
	out := make(chan *Table, 10) // solution tables
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	go tt.Solven(ctx, out)
	for {
		select {
		case ttt := <-out:
			ttt.Dump()
		case <-ctx.Done():
			return
		}
	}
}

func TestSolvenEmpty(t *testing.T) {

	tt := NewTable()
	out := make(chan *Table, 10) // solution tables
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	go tt.Solven(ctx, out)
	for {
		select {
		case ttt := <-out:
			ttt.Dump()
		case <-ctx.Done():
			return
		}
	}
}

func TestSolveSlice2(t *testing.T) {
	tt := NewTable()
	tt.Scan(strings.NewReader(real2))
	sol := tt.SolveSlice(2 * time.Second)
	for _, tt := range sol {
		tt.Dump()
	}
	fmt.Printf("There are %d solutions\n", len(sol))
	if len(sol) != 2 {
		t.Fatal("Unexpected length")
	}
	if sol[0].Equal(sol[1]) {
		t.Fatal("Unexpected equal solutions")
	}
}
