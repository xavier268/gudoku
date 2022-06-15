package sdk

import (
	"fmt"
	"math/rand"
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
	rand := rand.New(rand.NewSource(4242)) // deterministic

	for i := 0; i < 300; i++ {
		ti := time.Now().UnixMicro()
		p, _ := Build(rand)
		fmt.Printf("%d\tBuild a puzzle with %d values and %d zeros\t%9d ms\n", i, p.n, 9*9-p.n, time.Now().UnixMicro()-ti)
	}
}
