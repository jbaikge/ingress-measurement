package main

import (
	"github.com/cznic/mathutil"
	"sort"
)

// Generates the number of pads to insert between "fields" of a string to fill
// the string's width
type Spacer struct {
	Min, Max int // Min - Max space count for each grouping
	Spaces   int // Spaces required to pad rest of string
	State    []int
}

var (
	_ sort.Interface = new(Spacer) // Ensure Spacer follows sort.Interface

	MaxSpaces = len(spaces)
	spaces    = []byte("   ")
)

func NewSpacer(words, spaces int) (s *Spacer) {
	s = &Spacer{
		Min:    0,
		Max:    MaxSpaces - 1,
		Spaces: spaces,
		State:  make([]int, words+1),
	}
	return
}

// Creates an array of slices to insert between fields. Pads between the endcaps
// get an extra character to account for the normal word space.
func (s *Spacer) Bytes() (b [][]byte) {
	b = make([][]byte, len(s.State))
	for i, l := range s.State {
		b[i] = spaces[:l+1]
	}
	b[0] = b[0][1:]
	b[len(b)-1] = b[len(b)-1][1:]
	return
}

func (s *Spacer) Iter() [][]byte {
	i := 0
	for ; i < s.Spaces/s.Max && i < len(s.State); i++ {
		s.State[i] = s.Max
	}
	s.State[i] = s.Spaces % s.Max
	return s.Bytes()
}

func (s *Spacer) Next() [][]byte {
	if mathutil.PermutationNext(s) {
		return s.Bytes()
	}
	return nil
}

func (s *Spacer) Len() int           { return len(s.State) }
func (s *Spacer) Less(i, j int) bool { return s.State[i] > s.State[j] } // Intentionally backwards
func (s *Spacer) Swap(i, j int)      { s.State[i], s.State[j] = s.State[j], s.State[i] }
