package main

type Spacer struct {
	Len      int // Number of spaces to define, should be words + 2
	Min, Max int // Min - Max space count for each grouping
	Spaces   int // Spaces required to pad rest of string
	State    []int
}

const MaxSpaces = 3

var spaces [MaxSpaces]byte

func init() {
	for i := range spaces {
		spaces[i] = ' '
	}
}

func NewSpacer(words, spaces int) (s *Spacer) {
	s = &Spacer{
		Len:    words + 1,
		Min:    0,
		Max:    MaxSpaces - 1,
		Spaces: spaces,
		State:  make([]int, words+1),
	}
	return
}

func (s *Spacer) Bytes() (b [][]byte) {
	b = make([][]byte, s.Len)
	for i, l := range s.State {
		b[i] = spaces[:l]
	}
	return
}

// Increments next spacing configuration
func (s *Spacer) Increment() {
	s.State[0]++

	if s.State[0] <= s.Max {
		return
	}

	for c := 0; c < s.Len-1; c++ {
		if s.State[c] > s.Max {
			s.State[c] = s.Min
			s.State[c+1]++
			continue
		}
		break
	}
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
	return nil
}
