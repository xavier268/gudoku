package sdk

var indx [3 * 9][9]int // array of the 9 lines + 9 col + 9 groups of 9 indexes that represents the validity constraints.

// Initialize indx
func init() {

	// init lines
	for i := 0; i < 9*9; i++ {
		indx[i/9][i%9] = i
	}

	// init cols
	for i := 0; i < 9*9; i++ {
		indx[9+(i%9)][i/9] = i
	}

	// init group
	i := 18
	indx[i] = [9]int{0, 1, 2, 9, 10, 11, 18, 19, 20}
	i++
	indx[i] = [9]int{3, 4, 5, 12, 13, 14, 21, 22, 23}
	i++
	indx[i] = [9]int{6, 7, 8, 15, 16, 17, 24, 25, 26}
	i++
	indx[i] = [9]int{27, 28, 29, 36, 37, 38, 45, 46, 47}
	i++
	indx[i] = [9]int{30, 31, 32, 39, 40, 41, 48, 49, 50}
	i++
	indx[i] = [9]int{33, 34, 35, 42, 43, 44, 51, 52, 53}
	i++
	indx[i] = [9]int{54, 55, 56, 63, 64, 65, 72, 73, 74}
	i++
	indx[i] = [9]int{57, 58, 59, 66, 67, 68, 75, 76, 77}
	i++
	indx[i] = [9]int{60, 61, 62, 69, 70, 71, 78, 79, 80}
}
