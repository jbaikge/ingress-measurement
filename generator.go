package main

import (
	"bytes"
	"fmt"
)

// Generates the various spacing permutations to fill a string to the desired
// width
type Generator struct {
	Len    int
	Fields [][]byte
	Spaces uint64
	spacer *Spacer
}

func NewGenerator(s []byte, length int) (g *Generator, err error) {
	if l := len(s); l > length {
		err = fmt.Errorf("String too long to fit in container: %d > %d", l, length)
		return
	}

	g = &Generator{
		Len:    length,
		Fields: bytes.Fields(s),
	}
	g.spacer = NewSpacer(len(g.Fields), length-len(s))
	return
}

func (g *Generator) Bytes() (b []byte) {
	if g.Spaces == 0 {
		return nil
	}
	b = make([]byte, 0, g.Len)
	b = append(b, spaces[:g.Spaces&0xF]...)
	for i := uint(0); i < uint(len(g.Fields)); i++ {
		b = append(b, g.Fields[i]...)
		l := g.Spaces & (0xF << ((i + 1) * 4)) / (0xF << (i * 4))
		b = append(b, spaces[:l]...)
	}
	return b
}

// Reset state to beginning state
func (g *Generator) Iter() []byte {
	g.Spaces = g.spacer.Iter()
	return g.Bytes()
}

// Return next state or nil
func (g *Generator) Next() []byte {
	if g.Spaces = g.spacer.Next(); g.Spaces != 0 {
		return g.Bytes()
	}
	
	return nil
}
