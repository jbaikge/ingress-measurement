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
	State    []byte
	Words    int
}

var (
	_ sort.Interface = new(Spacer) // Ensure Spacer follows sort.Interface

	MaxSpaces = len(spaces)
	spaces    = []byte("XXX")
)

func NewSpacer(words, spaces int) (s *Spacer) {
	s = &Spacer{
		Min:    0,
		Max:    MaxSpaces,
		Spaces: spaces,
		Words:  words,
	}
	s.State = make([]byte, (s.Words+1)*s.Max)
	return
}

// Creates an array of slices to insert between fields. Pads between the endcaps
// get an extra character to account for the normal word space.
func (s *Spacer) Bytes() (b [][]byte) {
	b = make([][]byte, len(s.State)/s.Max)
	var l int
	for i := range b {
		l = 0
		for _, bit := range s.State[i*s.Max : (i+1)*s.Max] {
			if bit == 1 {
				l++
			}
		}
		b[i] = spaces[:l]
	}
	//b[0] = b[0][1:]
	//b[len(b)-1] = b[len(b)-1][1:]
	return
}

// Initalize the state
func (s *Spacer) Iter() [][]byte {
	for i := len(s.State) - s.Spaces - s.Words + 1; i < len(s.State); i++ {
		s.State[i] = 1
	}
	// Two words don't need the single-space redistribution that more than
	// two words require
	if s.Words > 2 {
		return s.Next()
	}
	return s.Bytes()
}

func (s *Spacer) Next() [][]byte {
Retry:
	if !mathutil.PermutationNext(s) {
		return nil
	}
	// Check for leading high-bits and call Next again if any are found
	// This removes duplicates by forcing the 10's to 01's
	var high bool
	for i := 0; i < len(s.State); i += s.Max {
		high = false
		for j := i; j < i+s.Max; j++ {
			if high && s.State[j] == 0 {
				goto Retry
			}
			high = s.State[j] == 1
		}
		if !high && i > 0 && i < len(s.State)-s.Max {
			goto Retry
		}
	}

	return s.Bytes()
}

func (s *Spacer) Len() int           { return len(s.State) }
func (s *Spacer) Less(i, j int) bool { return s.State[i] < s.State[j] } // Intentionally backwards
func (s *Spacer) Swap(i, j int)      { s.State[i], s.State[j] = s.State[j], s.State[i] }
