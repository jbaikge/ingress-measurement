package main

import (
	"testing"
)

func TestSpacerBytes(t *testing.T) {
	// 1 word, 1 space required
	s := NewSpacer(1, 1)
	b := s.Iter() // [[32] []]
	if b == nil {
		t.Fatal("Iter returned nil")
	}
	if l:=len(b);l != 2 {
		t.Fatal("Invalid length: %d", l)
	}
	if l := len(b[0]); l != 1 {
		t.Fatal("Invalid length of first set: %d", l)
	}
	if v := b[0][0]; v != ' ' {
		t.Fatal("Invalid value for first char: %d", v)
	}
}

func TestSingleSpacer(t *testing.T) {
	s := NewSpacer(1, 1)
	for v := s.Iter(); v != nil; v = s.Next() {
		t.Logf("%v", v)
	}
}
