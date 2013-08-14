package main

import (
	"testing"
)

var genTests = []struct {
	G   *Generator
	Exp []string
}{
	{
		NewGenerator("moo", 4),
		[]string{
			"moo ",
			" moo",
		},
	},
}

func TestGenerator(t *testing.T) {
	for _, test := range genTests {
		t.Logf("%v", test.G)
		for v := test.G.Iter(); v != nil; v = test.G.Next() {
			t.Logf("%v", v)
		}
	}
}
