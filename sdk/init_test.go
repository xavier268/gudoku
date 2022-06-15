package sdk

import (
	"fmt"
	"testing"
)

func TestInitVisual(_ *testing.T) {

	for i := 0; i < 3*9; i++ {
		fmt.Printf("\n%3d\t", i)
		for j := 0; j < 9; j++ {
			fmt.Printf("%5d", indx[i][j])
		}
	}
	fmt.Println()

}
