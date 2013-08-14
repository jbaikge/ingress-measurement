package main

import (
	"bytes"
)

type Generator struct {
	Len    int
	Fields [][]byte
	spacer *Spacer
}

func NewGenerator(s string, length int) (g *Generator) {
	g = &Generator{
		Len:    length,
		Fields: bytes.Fields([]byte(s)),
	}
	g.spacer = NewSpacer(len(g.Fields), length-len(s))
	return
}

// Reset state to beginning state
func (g *Generator) Iter() []byte {
	return nil
}

// Return next state or nil
func (g *Generator) Next() []byte {
	return nil
}
