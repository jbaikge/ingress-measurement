package main

import (
	"testing"
)

var genTests = []struct {
	G   *Generator
	Exp []string
}{
	{
		NewGenerator([]byte("moo"), 4),
		[]string{
			"mooX",
			"Xmoo",
		},
	},
	{
		NewGenerator([]byte("moo moo mr cow"), 19),
		[]string{
			"mooXmooXXmrXXXcowXX",
			"mooXmooXXXmrXXcowXX",
			"mooXmooXXXmrXXXcowX",
			"mooXXmooXmrXXXcowXX",
			"mooXXmooXXmrXXcowXX",
			"mooXXmooXXmrXXXcowX",
			"mooXXmooXXXmrXcowXX",
			"mooXXmooXXXmrXXcowX",
			"mooXXmooXXXmrXXXcow",
			"mooXXXmooXmrXXcowXX",
			"mooXXXmooXmrXXXcowX",
			"mooXXXmooXXmrXcowXX",
			"mooXXXmooXXmrXXcowX",
			"mooXXXmooXXmrXXXcow",
			"mooXXXmooXXXmrXcowX",
			"mooXXXmooXXXmrXXcow",
			"XmooXmooXmrXXXcowXX",
			"XmooXmooXXmrXXcowXX",
			"XmooXmooXXmrXXXcowX",
			"XmooXmooXXXmrXcowXX",
			"XmooXmooXXXmrXXcowX",
			"XmooXmooXXXmrXXXcow",
			"XmooXXmooXmrXXcowXX",
			"XmooXXmooXmrXXXcowX",
			"XmooXXmooXXmrXcowXX",
			"XmooXXmooXXmrXXcowX",
			"XmooXXmooXXmrXXXcow",
			"XmooXXmooXXXmrXcowX",
			"XmooXXmooXXXmrXXcow",
			"XmooXXXmooXmrXcowXX",
			"XmooXXXmooXmrXXcowX",
			"XmooXXXmooXmrXXXcow",
			"XmooXXXmooXXmrXcowX",
			"XmooXXXmooXXmrXXcow",
			"XmooXXXmooXXXmrXcow",
			"XXmooXmooXmrXXcowXX",
			"XXmooXmooXmrXXXcowX",
			"XXmooXmooXXmrXcowXX",
			"XXmooXmooXXmrXXcowX",
			"XXmooXmooXXmrXXXcow",
			"XXmooXmooXXXmrXcowX",
			"XXmooXmooXXXmrXXcow",
			"XXmooXXmooXmrXcowXX",
			"XXmooXXmooXmrXXcowX",
			"XXmooXXmooXmrXXXcow",
			"XXmooXXmooXXmrXcowX",
			"XXmooXXmooXXmrXXcow",
			"XXmooXXmooXXXmrXcow",
			"XXmooXXXmooXmrXcowX",
			"XXmooXXXmooXmrXXcow",
			"XXmooXXXmooXXmrXcow",
		},
	},
}

func TestGenerator(t *testing.T) {
	for _, test := range genTests {
		for i, v := 0, test.G.Iter(); v != nil; i, v = i+1, test.G.Next() {
			if exp := test.Exp[i]; exp != string(v) {
				t.Errorf("Expected '%s'; Got '%s'", exp, v)
			}
		}
	}
}
