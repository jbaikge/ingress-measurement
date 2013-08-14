package main

import (
	"fmt"
	"time"
)

type Encoded string

type Format string

func (f Format) Encode(t time.Time) (e Encoded) {
	return Encoded(fmt.Sprintf(""))
}
