package main

import (
	"sort"
)

// Generates the number of pads to insert between "fields" of a string to fill
// the string's width
type Spacer struct {
	Spaces   uint64 // Spaces required to pad rest of string
	Words    uint64
	State    uint64
	Required uint64
	ch       chan uint64
	notify   chan uint64
	past     Seen
}

var spaces = []byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

func NewSpacer(words, spaces int) (s *Spacer) {
	s = &Spacer{
		Spaces: uint64(spaces),
		Words:  uint64(words),
	}
	s.Required = s.Spaces + s.Words - 1
	return
}

// Thanks to Derek Mauro for the recursive algorithm
func (s *Spacer) Space(pos, remain, maxWidth uint64) {
	var min, max uint64
	if pos == s.Words+1 {
		s.notify <- s.State
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
		s.State &= ^(0xF << (pos * 4))
		s.State |= i << (pos * 4)
		s.Space(pos+1, remain-i, maxWidth)
	}
	if pos == 0 {
		close(s.notify)
	}
}

// Initalize the state
func (s *Spacer) Iter() uint64 {
	s.ch = make(chan uint64)
	go func() {
		// 3-space check
		// log.Println("3-space check")
		s.notify = make(chan uint64)
		go func() {
			for v := range s.notify {
				if !s.past.Seen(v) {
					s.ch <- v
				}
			}
		}()
		s.Space(0, s.Required, uint64(3))

		// log.Println("DeltaDown")
		s.DeltaDown()
		// If we return here, it means the entire program has not exited..

		s.notify = make(chan uint64)
		go func() {
			for v := range s.notify {
				if !s.past.Seen(v) {
					s.ch <- v
				}
			}
			close(s.ch)
		}()
		// log.Println("Longhaul")
		s.Space(0, s.Required, uint64(Config.MaxSpaces))
	}()
	return <-s.ch
}

func (s *Spacer) Next() uint64 {
	return <-s.ch
}

func (s *Spacer) DeltaDown() {
	for r := s.Required / 2; r < s.Required-3; r++ {
		Δ := uint64(s.Required - r)

		s.notify = make(chan uint64)
		go s.Space(0, r, uint64(Config.MaxSpaces))

		for state := range s.notify {
			for i := uint64(0); i <= s.Words; i++ {
				modified := state + (Δ << (i * 4))
				if !s.past.Seen(modified) {
					s.ch <- modified
				}
			}
		}
	}
}

type Seen []uint64

func (s Seen) Seen(n uint64) bool {
	i := sort.Search(len(s), func(i int) bool { return s[i] >= n })
	if i < len(s) && s[i] == n {
		return true
	} else {
		s = append(s, uint64(0))
		copy(s[i+1:], s[i:])
		s[i] = n
	}
	return false
}
