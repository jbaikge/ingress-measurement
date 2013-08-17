package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"sync"
	"time"
)

type Package struct {
	Format      Format       // Possible time string format
	Hash        []byte       // Hash provided by Ingress
	Measurement int          // Measurement number provided by Ingress
	TimeRange   [2]time.Time // Range of times to test, t0 provided by Ingress
	Encrypted   []byte       // Timestring provided by Ingress
	OTP         []byte
	TimeString  []byte
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
		Encrypted: []byte(ts),
		hasher:    md5.New(),
	}
	if p.Hash, err = hex.DecodeString(h); err != nil {
		return
	}
	return
}

func (p *Package) Find() bool {
	found := make(chan bool)
	var wg sync.WaitGroup

	for t := p.TimeRange[0]; t.Before(p.TimeRange[1]); t = t.Add(time.Second) {
		wg.Add(1)
		go func(t time.Time) {
			defer wg.Done()
			hasher := md5.New()

			f := p.Format.Encode(p.Measurement, t)

			g, err := NewGenerator(f, len(p.Encrypted))
			if err != nil {
				return
			}

			for s := g.Iter(); s != nil; s = g.Next() {
				otp := OTP(s, p.Encrypted)
				hasher.Write(otp)
				sum := hasher.Sum(nil)
				if sum[0] == p.Hash[0] && bytes.Equal(sum, p.Hash) {
					p.OTP = otp
					p.TimeString = s
					found <- true
					return
				}
				hasher.Reset()
			}
		}(t)
	}

	go func() {
		wg.Wait()
		found <- false
	}()

	return <-found
}
