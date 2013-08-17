package main

import (
	"testing"
)

func TestSpacerBytes(t *testing.T) {
	// 1 word, 1 space required
	s := NewSpacer(1, 1)
	b := s.Iter() // [[] [88]]
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
	expect := []int{
		0, 1, 3,
		0, 2, 2,
		0, 3, 1,
		0, 4, 0,
		1, 1, 2,
		1, 2, 1,
		1, 3, 0,
		2, 1, 1,
		2, 2, 0,
		3, 1, 0,
	}
	for i, v := 0, s.Iter(); v != nil; i, v = i+3, s.Next() {
		for j, slice := range v {
			got := len(slice)
			if exp := expect[i+j]; exp != got {
				t.Errorf("[%d][%d] %v Expected %d; Got %d", i, j, s.State, exp, got)
			}
		}
	}
}
