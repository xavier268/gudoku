package sdk

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestBuilderVisual(t *testing.T) {
	rand := rand.New(rand.NewSource(42)) // deterministic

	for i := 0; i < 10; i++ {
		p, s := Build(rand)

		fmt.Println("PUZZLE :")
		p.Dump()
		fmt.Println("SOLUTION :")
		s.Dump()
	}
}

func TestBuilderStats(t *testing.T) {
	rand := rand.New(rand.NewSource(42)) // deterministic

	for i := 0; i < 50; i++ {
		ti := time.Now().UnixMicro()
		p, _ := Build(rand)
		fmt.Printf("%d\tBuild a puzzle with %d values and %d zeros\t%9d ms\n", i, p.n, 9*9-p.n, time.Now().UnixMicro()-ti)
	}
}

func TestBuildFromSolution1(t *testing.T) {

	solution := NewTable()
	solution.Scan(strings.NewReader("123789456 456123789 789456123 231897564 564231897 897564231 312978645 645312978 978645312"))

	ti := time.Now().UnixMicro()
	p := BuildFromSolution(solution)
	p.Dump()
	fmt.Printf("\nBuildFrom achieved %d values and %d zeros\t%9d ms\n", p.n, 9*9-p.n, time.Now().UnixMicro()-ti)

}
