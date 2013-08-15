package main

import (
	"testing"
)

func TestSpacerBytes(t *testing.T) {
	// 1 word, 1 space required
	s := NewSpacer(1, 1)
	b := s.Iter() // [[32] []]
	if b == nil {
		t.Fatalf("Iter returned nil")
	}
	if l := len(b); l != 2 {
		t.Fatalf("Invalid length: %d", l)
	}
	if l := len(b[1]); l != 1 {
		t.Fatalf("Invalid length of first set: %d", l)
	}
	if v := b[1][0]; v != spaces[0] {
		t.Fatalf("Invalid value for first char: %d", v)
	}
}

func TestSpacerSingle(t *testing.T) {
	// Two words, need 3 spaces to complete the required width
	s := NewSpacer(2, 3)
	expect := []byte{
		0, 0, 0, 0, 0, 1, 1, 1, 1,
		0, 0, 0, 0, 1, 1, 0, 1, 1,
		0, 0, 0, 1, 1, 1, 0, 0, 1,
		0, 0, 1, 0, 0, 1, 0, 1, 1,
		0, 0, 1, 0, 1, 1, 0, 0, 1,
		0, 0, 1, 1, 1, 1, 0, 0, 0,
		0, 1, 1, 0, 0, 1, 0, 0, 1,
		0, 1, 1, 0, 1, 1, 0, 0, 0,
		1, 1, 1, 0, 0, 1, 0, 0, 0,
	}
	for i, v := 0, s.Iter(); v != nil; i, v = i+9, s.Next() {
		for j, got := range s.State {
			if exp := expect[i+j]; exp != got {
				t.Errorf("[%d][%d] %v Expected %d; Got %d", i, j, s.State, exp, got)
			}
		}
	}
}
