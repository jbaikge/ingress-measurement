package main

import (
	"github.com/cznic/mathutil"
	"log"
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

	EndCap    = 2
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
	s.State = make([]byte, (words-1)*s.Max+EndCap*2)
	return
}

// Creates an array of slices to insert between fields. Pads between the endcaps
// get an extra character to account for the normal word space.
func (s *Spacer) Bytes() (b [][]byte) {
	b = make([][]byte, s.Words+1)
	var l int
	for i := range b {
		l = 0
		for _, bit := range s.Slice(i) {
			if bit == 1 {
				l++
			}
		}
		b[i] = spaces[:l]
	}
	return
}

// Initalize the state
func (s *Spacer) Iter() [][]byte {
	log.Printf("%v", *s)
	for i := len(s.State) - s.Spaces; i < len(s.State); i++ {
		s.State[i] = 1
	}
	return s.Bytes()
}

func (s *Spacer) Next() [][]byte {
	if !mathutil.PermutationNext(s) {
		return nil
	}
	// Check for leading high-bits and call Next again if any are found
	// This removes duplicates by forcing the 10's to 01's
	var high bool
	for i := 0; i < s.Words; i++ {
		high = false
		for _, bit := range s.Slice(i) {
			if high && bit == 0 {
				goto ReNext
			}
			high = bit == 1
		}
	}

	return s.Bytes()

ReNext:
	return s.Next()
}

func (s *Spacer) Slice(i int) (slice []byte) {
	switch i {
	case 0:
		slice = s.State[:EndCap]
	case s.Words:
		slice = s.State[len(s.State)-EndCap:]
	default:
		log.Printf("[%d] %d [%d : %d]", i, len(s.State), (i-1)*s.Max+EndCap, (i)*s.Max+EndCap)
		slice = s.State[(i-1)*s.Max+EndCap : (i)*s.Max+EndCap]
	}
	return
}

func (s *Spacer) Len() int           { return len(s.State) }
func (s *Spacer) Less(i, j int) bool { return s.State[i] < s.State[j] } // Intentionally backwards
func (s *Spacer) Swap(i, j int)      { s.State[i], s.State[j] = s.State[j], s.State[i] }
