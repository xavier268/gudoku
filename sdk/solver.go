package sdk

// Solve attempts to solve provided table, returning true on success, or false if failed.
// The table contains a solution, if found.
func (t *Table) Solve() bool {

	// fmt.Println("Solving for ", t.n)
	// t.Dump()

	if t.n == 9*9 {
		return t.Valid() // done !
	}

	for a := 0; a < 9*9; a++ {
		v := t.Get(a)
		if v == 0 {
			for i := 1; i <= 9; i++ {
				t.Set(a, i)
				if t.Valid() && t.Solve() {
					return true
				} // else iterate values
			}
			t.Set(a, 0) // restore state before returning
			return false
		}

	}

	return false
}

// SolveBack is the same as Solve, but working from the other end of the solution space.
// Used to ensure unicity of a solution.
func (t *Table) SolveBack() bool {

	if t.n == 9*9 {
		return t.Valid()
	}

	for a := 9*9 - 1; a >= 0; a-- {
		v := t.Get(a)
		if v == 0 {
			for i := 9; i > 0; i-- {
				t.Set(a, i)
				if t.Valid() && t.SolveBack() {
					return true
				} // else iterate values
			}
			t.Set(a, 0) // restore state before returning
			return false
		}
	}
	return false
}
