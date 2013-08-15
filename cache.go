package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

const (
	MinWordCount, MaxWordCount   = 1, 9
	MinSpaceCount, MaxSpaceCount = 0, 12
)

var SpaceCache [MaxWordCount + 1][MaxSpaceCount + 1][][][]byte

func init() {
	fn := "SpaceCache.gob"
	f, err := os.Open(fn)
	if err != nil {
		f, err := os.Create(fn)
		if err != nil {
			fmt.Printf("Error opening %s for writing: %s\n", fn, err)
		}
		defer f.Close()
		RebuildCache(f)
		return
	}

	defer f.Close()
	dec := gob.NewDecoder(f)
	dec.Decode(&SpaceCache)
}
func RebuildCache(w io.Writer) error {
	fmt.Println("Rebuilding cache...")

	for w := MinWordCount; w <= MaxWordCount; w++ {
		for s := MinSpaceCount; s <= MaxSpaceCount && s < w*MaxSpaces-1; s++ {
			fmt.Printf("W: %d S: %d\n", w, s)
			spacer := NewSpacer(w, s)
			for v := spacer.Iter(); v != nil; v = spacer.Next() {
				SpaceCache[w][s] = append(SpaceCache[w][s], v)
			}
		}
	}

	fmt.Println("Saving cache...")
	enc := gob.NewEncoder(w)
	return enc.Encode(SpaceCache)
}
