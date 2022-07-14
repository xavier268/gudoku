package sdk

import (
	"fmt"
	"strings"
	"testing"
)

var real1 string = "100000006 006020700 789450103 000807004 000030000 090004201 312970040 040012078 908000000" // single solution
//var real2 string = "100000006 006020700 789450103 000807004 000030000 090004201 312970040 040010078 908000000" // two solutions

func TestSolverEmpty(t *testing.T) {
	tt := NewTable()
	if tt.Solve() {
		tt.Dump()
	} else {
		t.Fatalf("Could not find a solution even starting with empty table !")
	}
	ttt := NewTable()
	if !ttt.SolveBack() {
		t.Fatalf("Could not find a backward solution even starting with empty table !")
	}
	if tt.Equal(ttt) {
		t.Fatalf("Unexpected unicity of the soution to the empty table !")
	}
}

func TestSolverReal(t *testing.T) {
	// Forward
	tt := NewTable()
	tt.Scan(strings.NewReader(real1))
	if !tt.Solve() {
		t.Fatalf("Could not solve real1 backwards")
	}
	// Backward
	ttt := NewTable()
	ttt.Scan(strings.NewReader(real1))
	if !ttt.Solve() {
		t.Fatalf("Could not solve real1 backwards")
	}
	if !tt.Equal(ttt) {
		t.Fatal("Expected a unique solution !")
	}
	fmt.Println("Solution is UNIQUE !")
	tt.Dump()
}
