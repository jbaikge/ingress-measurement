package main

import (
	"encoding/hex"
	"time"
)

type Package struct {
	Format     Format       // Possible time string format
	Hash       []byte       // Hash provided by Ingress
	TimeRange  [2]time.Time // Range of times to test, t0 provided by Ingress
	TimeString []byte       // Timestring provided by Ingress
}

func NewPackage(f Format, start time.Time, ts string, h string) (p *Package, err error) {
	p = &Package{
		Format: f,
		TimeRange: [2]time.Time{
			start,
			start.Add(5 * time.Minute),
		},
		TimeString: []byte(ts),
	}
	if p.Hash, err = hex.DecodeString(h); err != nil {
		return
	}
	return
}
