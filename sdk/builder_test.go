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
		p, s := BuildRandom(rand)

		fmt.Println("PUZZLE :")
		p.Dump()
		fmt.Println("SOLUTION :")
		s.Dump()
	}
}

func TestBuilderStats(t *testing.T) {
	rand := rand.New(rand.NewSource(42)) // deterministic

	sum := 0
	count := 0

	for i := 0; i < 50; i++ {
		ti := time.Now().UnixMicro()
		p, _ := BuildRandom(rand)
		fmt.Printf("%d\tBuild a puzzle with %d values and %d zeros\t%9d ms\n", i, p.n, 9*9-p.n, time.Now().UnixMicro()-ti)
		sum += p.n
		count++
	}
	n := float64(sum) / float64(count)
	fmt.Printf("\n\tAverage : %.1f non zero and %.1f zero values\n", n, 9*9-n)
}

func TestBuildFromSolution(_ *testing.T) {

	solution := NewTable()
	solution.Scan(strings.NewReader("123789456 456123789 789456123 231897564 564231897 897564231 312978645 645312978 978645312"))

	ti := time.Now().UnixMicro()
	p := BuildFromSolution(solution)
	p.Dump()
	fmt.Printf("\nBuildFrom achieved %d values and %d zeros\t%9d ms\n", p.n, 9*9-p.n, time.Now().UnixMicro()-ti)

}

func TestShuffleVisual(t *testing.T) {

	s := NewShuffler(rand.New(rand.NewSource(42)))

	for a := 0; a < 9*9; a++ {
		fmt.Printf("%3d", s.pp[a])
		switch a % 9 {
		case 2, 5:
			fmt.Print(" | ")
		case 8:
			fmt.Println()
		default:
		}
	}

	t0 := NewTable()
	t0.Scan(strings.NewReader(real1))
	t0.Dump()

	t1 := t0.Clone()
	t2 := t1.Clone()
	t1.Dump("t1 before shuffle")
	s.Shuffle(t1, t2)
	t1.Dump("t1 after shuffle")
	if t0.Equal(t1) {
		t.Fatalf("shuffle should not leave a table unchanged")
	}
	if !t1.Equal(t2) {
		t.Fatalf("shuffle was not applied in identical ways")
	}
	/*s.Shuffle(t1)
	if !t1.Equal(t0) {
		t0.Dump()
		t1.Dump()
		t2.Dump()
		t.Fatal("unshuffle did not succeeded")
	}*/

}
