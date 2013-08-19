package main

import (
	"testing"
	"time"
)

// func TestSpacerBytes(t *testing.T) {
// 	// 1 word, 1 space required
// 	s := NewSpacer(1, 1)
// 	b := s.Iter() // [[] [88]]
// 	if b == nil {
// 		t.Fatalf("Iter returned nil")
// 	}
// 	if l := len(b); l != 2 {
// 		t.Fatalf("Invalid length: %d", l)
// 	}
// 	if l := len(b[1]); l != 1 {
// 		t.Fatalf("Invalid length of first set: %d", l)
// 	}
// 	if v := b[1][0]; v != spaces[0] {
// 		t.Fatalf("Invalid value for first char: %d", v)
// 	}
// }

// func TestSpacerSingle(t *testing.T) {
// 	// Two words, need 3 spaces to complete the required width
// 	s := NewSpacer(2, 3)
// 	expect := []int{
// 		0, 1, 3,
// 		0, 2, 2,
// 		0, 3, 1,
// 		0, 4, 0,
// 		1, 1, 2,
// 		1, 2, 1,
// 		1, 3, 0,
// 		2, 1, 1,
// 		2, 2, 0,
// 		3, 1, 0,
// 	}
// 	for i, v := 0, s.Iter(); v != nil; i, v = i+3, s.Next() {
// 		for j, got := range v {
// 			if exp := expect[i+j]; exp != got {
// 				t.Errorf("[%d][%d] %v Expected %d; Got %d", i, j, s.State, exp, got)
// 			}
// 		}
// 	}
// }

func TestSpacerLog(t *testing.T) {
	s := NewSpacer(2, 3)
	start := time.Now()
	for v := s.Iter(); v != 0; v = s.Next() {
		t.Logf("%v %03x", v, v)
		for pos := uint64(0); pos < s.Words; pos++ {
			l := v & (0xF << ((pos + 1) * 4)) / (0xF << (pos * 4))
			//l := v & 0xF00 / 0xF0
			t.Logf("pos[%d] %d %x", pos, l, 0xF<<((pos+1)*4))
		}
	}
	t.Logf("Took %s", time.Since(start))
}
