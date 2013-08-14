package main

import (
	"testing"
)

func TestSpacerBytes(t *testing.T) {
	// 1 word, 1 space required
	s := NewSpacer(1, 1)
	s.Iter() // prime
	t.Logf("%v", s.Bytes())
}

func TestSingleSpacer(t *testing.T) {
	s := NewSpacer(1, 1)
	for v := s.Iter(); v != nil; v = s.Next() {
		t.Logf("%v", v)
	}
}
