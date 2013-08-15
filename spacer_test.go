package main

import (
	"testing"
)

func TestSpacerBytes(t *testing.T) {
	// 1 word, 1 space required
	s := NewSpacer(1, 1)
	b := s.Iter() // [[32] []]
	t.Logf("%v", b)
	if b == nil {
		t.Fatalf("Iter returned nil")
	}
	if l := len(b); l != 2 {
		t.Fatalf("Invalid length: %d", l)
	}
	if l := len(b[0]); l != 1 {
		t.Fatalf("Invalid length of first set: %d", l)
	}
	if v := b[0][0]; v != spaces[0] {
		t.Fatalf("Invalid value for first char: %d", v)
	}
}

func TestSingleSpacer(t *testing.T) {
	// Two words, need 3 spaces to complete the required width
	s := NewSpacer(2, 3)
	expect := []int{
		2, 1, 0,
		2, 0, 1,
		1, 2, 0,
		1, 0, 2,
		0, 2, 1,
		0, 1, 2,
	}
	for i, v := 0, s.Iter(); v != nil; i, v = i+3, s.Next() {
		for j, got := range s.State {
			if exp := expect[i+j]; exp != got {
				t.Error("[%d][%d] Expected %d; Got %d", i, j, exp, got)
			}
		}
	}
}
