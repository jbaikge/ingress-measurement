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
			" moo",
			"moo ",
		},
	},
	{
		NewGenerator("moo moo mr cow", 19),
		[]string{
			"  moo   moo  mr cow",
			"  moo   moo mr  cow",
			"  moo   moo mr cow ",
			"  moo  moo   mr cow",
			"  moo  moo mr   cow",
			"  moo  moo mr cow  ",
			"  moo moo   mr  cow",
			"  moo moo   mr cow ",
			"  moo moo  mr   cow",
			"  moo moo  mr cow  ",
			"  moo moo mr   cow ",
			"  moo moo mr  cow  ",
			" moo   moo   mr cow",
			" moo   moo mr   cow",
			" moo   moo mr cow  ",
			" moo moo   mr   cow",
			" moo moo   mr cow  ",
			" moo moo mr   cow  ",
			"moo   moo   mr  cow",
			"moo   moo   mr cow ",
			"moo   moo  mr   cow",
			"moo   moo  mr cow  ",
			"moo   moo mr   cow ",
			"moo   moo mr  cow  ",
			"moo  moo   mr   cow",
			"moo  moo   mr cow  ",
			"moo  moo mr   cow  ",
			"moo moo   mr   cow ",
			"moo moo   mr  cow  ",
			"moo moo  mr   cow  ",
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
