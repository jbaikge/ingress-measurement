package main

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	tests := []struct {
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
				"mooXmooXmrXXXcowXXX",
				"mooXmooXXmrXXcowXXX",
				"mooXmooXXmrXXXcowXX",
				"mooXmooXXXmrXcowXXX",
				"mooXmooXXXmrXXcowXX",
				"mooXmooXXXmrXXXcowX",
				"mooXXmooXmrXXcowXXX",
				"mooXXmooXmrXXXcowXX",
				"mooXXmooXXmrXcowXXX",
				"mooXXmooXXmrXXcowXX",
				"mooXXmooXXmrXXXcowX",
				"mooXXmooXXXmrXcowXX",
				"mooXXmooXXXmrXXcowX",
				"mooXXmooXXXmrXXXcow",
				"mooXXXmooXmrXcowXXX",
				"mooXXXmooXmrXXcowXX",
				"mooXXXmooXmrXXXcowX",
				"mooXXXmooXXmrXcowXX",
				"mooXXXmooXXmrXXcowX",
				"mooXXXmooXXmrXXXcow",
				"mooXXXmooXXXmrXcowX",
				"mooXXXmooXXXmrXXcow",
				"XmooXmooXmrXXcowXXX",
				"XmooXmooXmrXXXcowXX",
				"XmooXmooXXmrXcowXXX",
				"XmooXmooXXmrXXcowXX",
				"XmooXmooXXmrXXXcowX",
				"XmooXmooXXXmrXcowXX",
				"XmooXmooXXXmrXXcowX",
				"XmooXmooXXXmrXXXcow",
				"XmooXXmooXmrXcowXXX",
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
				"XXmooXmooXmrXcowXXX",
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
				"XXXmooXmooXmrXcowXX",
				"XXXmooXmooXmrXXcowX",
				"XXXmooXmooXmrXXXcow",
				"XXXmooXmooXXmrXcowX",
				"XXXmooXmooXXmrXXcow",
				"XXXmooXmooXXXmrXcow",
				"XXXmooXXmooXmrXcowX",
				"XXXmooXXmooXmrXXcow",
				"XXXmooXXmooXXmrXcow",
				"XXXmooXXXmooXmrXcow",
			},
		},
	}

	for _, test := range tests {
		for i, v := 0, test.G.Iter(); v != nil; i, v = i+1, test.G.Next() {
			if exp := test.Exp[i]; exp != string(v) {
				t.Errorf("Expected '%s'; Got '%s'", exp, v)
			}
		}
	}
}
