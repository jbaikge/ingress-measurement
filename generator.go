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
	Spaces [][]byte
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
	if len(g.Spaces) == 0 {
		return nil
	}
	b = make([]byte, 0, g.Len)
	b = append(b, g.Spaces[0]...)
	for i := 0; i < len(g.Fields); i++ {
		b = append(b, g.Fields[i]...)
		b = append(b, g.Spaces[i+1]...)
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
	if g.Spaces = g.spacer.Next(); g.Spaces != nil {
		return g.Bytes()
	}
	return nil
}
