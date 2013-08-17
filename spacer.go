package main

// Generates the number of pads to insert between "fields" of a string to fill
// the string's width
type Spacer struct {
	Spaces int // Spaces required to pad rest of string
	State  []byte
	Words  int

	ch chan [][]byte
}

var spaces = []byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

func NewSpacer(words, spaces int) (s *Spacer) {
	s = &Spacer{
		Spaces: spaces,
		Words:  words,
	}
	s.State = make([]byte, s.Words+1)
	return
}

// Creates an array of slices to insert between fields. Pads between the endcaps
// get an extra character to account for the normal word space.
func (s *Spacer) Bytes() (b [][]byte) {
	b = make([][]byte, len(s.State))
	for i := range b {
		if i > 0 && i < len(b)-1 {
			// Spacing between words
			b[i] = spaces[:s.State[i]+1]
		} else {
			b[i] = spaces[:s.State[i]]
		}
	}
	return
}

// Thanks to Derek Mauro for the recursive algorithm
func (s *Spacer) Space(pos, remain int) {
	var min, max int
	if pos == len(s.State) {
		s.ch <- s.Bytes()
		return
	}

	switch pos {
	case 0:
		min, max = 0, remain
	case s.Words:
		min, max = remain, remain
	default:
		min, max = 0, remain
	}

	for i := min; i <= max; i++ {
		s.State[pos] = byte(i)
		s.Space(pos+1, remain-i)
	}

	// Close channel after first iteration completes
	if pos == 0 {
		close(s.ch)
	}
}

// Initalize the state
func (s *Spacer) Iter() [][]byte {
	s.ch = make(chan [][]byte)
	go s.Space(0, s.Spaces)
	return <-s.ch
}

func (s *Spacer) Next() [][]byte {
	return <-s.ch
}
