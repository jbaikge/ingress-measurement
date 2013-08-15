package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"time"
)

type Package struct {
	Format      Format       // Possible time string format
	Hash        []byte       // Hash provided by Ingress
	Measurement int          // Measurement number provided by Ingress
	TimeRange   [2]time.Time // Range of times to test, t0 provided by Ingress
	TimeString  []byte       // Timestring provided by Ingress
	OTP         []byte
	hasher      hash.Hash
}

func NewPackage(f Format, m int, start time.Time, ts string, h string) (p *Package, err error) {
	p = &Package{
		Format:      f,
		Measurement: m,
		TimeRange: [2]time.Time{
			start,
			start.Add(5 * time.Minute),
		},
		TimeString: []byte(ts),
		hasher:     md5.New(),
	}
	if p.Hash, err = hex.DecodeString(h); err != nil {
		return
	}
	return
}

func (p *Package) Find() bool {
	for t := p.TimeRange[0]; t.Before(p.TimeRange[1]); t = t.Add(time.Second) {
		f := p.Format.Encode(p.Measurement, t)
		g := NewGenerator(f, len(p.TimeString))
		for s := g.Iter(); s != nil; s = g.Next() {
			otp := OTP(s, p.TimeString)
			p.hasher.Write(otp)
			sum := p.hasher.Sum(nil)
			if sum[0] == p.Hash[0] && bytes.Equal(sum, p.Hash) {
				p.OTP = otp
				return true
			}
			p.hasher.Reset()
		}
	}
	return false
}
