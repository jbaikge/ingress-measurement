package main

import (
	//"log"
	"sync"
)

// Generates the number of pads to insert between "fields" of a string to fill
// the string's width
type Spacer struct {
	Spaces uint64 // Spaces required to pad rest of string
	Words  uint64
	State  uint64
	ch     chan uint64
	wait   sync.WaitGroup
}

var spaces = []byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

func NewSpacer(words, spaces int) (s *Spacer) {
	s = &Spacer{
		Spaces: uint64(spaces),
		Words:  uint64(words),
	}
	return
}

// Thanks to Derek Mauro for the recursive algorithm
func (s *Spacer) Space(pos, remain, maxWidth uint64) {
	var min, max uint64
	if pos == s.Words+1 {
		s.ch <- s.State
		return
	}

	switch pos {
	case 0:
		min, max = 0, remain
	case s.Words:
		min, max = remain, remain
	default:
		min, max = 1, remain
	}

	if max > maxWidth {
		max = maxWidth
	}

	for i := min; i <= max; i++ {
		s.State &= ^(0xF << uint(pos*4))
		s.State |= i << uint(pos*4)
		s.Space(pos+1, remain-i, 0xF)
	}
	if pos == 0 {
		close(s.ch)
	}
}

// Initalize the state
func (s *Spacer) Iter() uint64 {
	s.ch = make(chan uint64)
	go s.Space(0, s.Spaces+s.Words-1, s.Spaces+1)
	return <-s.ch
}

func (s *Spacer) Next() uint64 {
	return <-s.ch
}
