package sdk

import "math/rand"

type Shuffler struct {
	ii [10]int    // map values
	pp [9 * 9]int // map positions
	rd *rand.Rand
}

func NewShuffler(rd *rand.Rand) *Shuffler {
	s := new(Shuffler)
	s.rd = rd
	s.Reset()

	return s
}

func (s *Shuffler) Reset() {

	for i := 1; i < 10; i++ {
		s.ii[i] = i
	}
	s.ii[0] = 0
	for i := 0; i < 9*9; i++ {
		s.pp[i] = i
	}

	// shuffle values
	for i := 0; i < 20; i++ {
		a := s.rd.Intn(8) + 1
		b := s.rd.Intn(8) + 1
		s.ii[a], s.ii[b] = s.ii[b], s.ii[a]
	}

	// shuffle lines/cols in the same group
	for i := 0; i < 20; i++ {
		a := s.rd.Intn(18)
		b := 3*(a/3) + s.rd.Intn(3)
		for r := 0; r < 9; r++ {
			s.pp[indx[a][r]], s.pp[indx[b][r]] = s.pp[indx[b][r]], s.pp[indx[a][r]]
		}
	}

}

// Shuffle all tables in a random, but always identical way, until a Reset is performed.
// Use it to simultaneously shuffle puzzle and solution.
func (s *Shuffler) Shuffle(tt ...*Table) {

	for _, t := range tt {
		s.shuffleValue(t)
	}
	for _, t := range tt {
		s.shufflePos(t)
	}

}

// shuffleValue will map each value in place to another.
// Zero are obviously unchanged.
func (s *Shuffler) shuffleValue(t *Table) {
	if t == nil {
		return
	}
	for a := 0; a < 9*9; a++ {
		t.Set(a, s.ii[t.Get(a)])
	}
}

func (s Shuffler) shufflePos(t *Table) {
	if t == nil {
		return
	}
	tt := t.Clone()
	for p := 0; p < 9*9; p++ {
		t.Set(p, tt.Get(s.pp[p]))
	}
}
