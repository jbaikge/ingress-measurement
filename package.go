package main

import (
	"time"
)

type Package struct {
	Format     Format       // Time string layout
	OTPKey     string       // OTP Key (gets hashed)
	TargetHash string       // Hash provided by Ingress
	Time       time.Time    // Time inserted into Format
	TimeRange  [2]time.Time // Range of times to test
}

func (p Package) String() string {
	return ""
}

func (p Package) Hash() {}
