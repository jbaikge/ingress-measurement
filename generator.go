package main

import (
	"bytes"
	"log"
)

// Generates the various spacing permutations to fill a string to the desired
// width
type Generator struct {
	Len    int
	Fields [][]byte
	Spaces [][]byte
	cache  [][][]byte
	i      int // cache index
}

func NewGenerator(s []byte, length int) (g *Generator) {
	g = &Generator{
		Len:    length,
		Fields: bytes.Fields(s),
	}
	padding := length - len(s)
	log.Printf("Getting SpaceCache[%d][%d]", len(g.Fields), padding)
	if padding < 0 || padding > MaxSpaceCount {
		return
	}
	g.cache = SpaceCache[len(g.Fields)][padding]
	return
}

func (g *Generator) Bytes() (b []byte) {
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
	if len(g.cache) == 0 {
		return nil
	}
	g.i = 0
	g.Spaces = g.cache[g.i]
	return g.Bytes()
}

// Return next state or nil
func (g *Generator) Next() []byte {
	if g.i++; g.i == len(g.cache) {
		return nil
	}

	g.Spaces = g.cache[g.i]
	return g.Bytes()
}
