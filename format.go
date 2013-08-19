package main

import (
	"bytes"
	"time"
)

// Format of time string:
// #  Measurement number
// h  Hour
// h_ Hour followed by HOUR or HOURS
// m  Minute
// m_ Minute followed by MINUTE or MINUTES
// s  Second
// s_ Second followed by SECOND or SECONDS
type Format []byte

var (
	Shour   = []byte(` HOUR`)
	Sminute = []byte(` MINUTE`)
	Ssecond = []byte(` SECOND`)
)

var Formats = map[string]Format{
	"FMinator":          Format(`h O CLOCK m_ AND s_`),
	"FCassandraSydney1": Format(`m_ s_ PAST h O CLOCK`),
	"FCassandraSydney2": Format(`h O CLOCK AND m_ AND s_`),
	"FCassandraMilan1":  Format(`s_ AND m_ AFTER h`),
	"FCassandraMilan3":  Format(`MEASUREMENT # IS AT h O m AND s_`),
	"FDusseldorf2":      Format(`m_ AND s_ PAST h O CLOCK`),
	"FDusseldorf3":      Format(`s_ PAST h O CLOCK SHARP`),
	"FDusseldorf4":      Format(`m_ s_ AFTER h O CLOCK`),
	"Manilla1":          Format(`h O CLOCK m_ s_`),
	"Manilla2":          Format(`m_ s_ PAST h PM`),
	"Manilla3":          Format(`m_ AND s_ PAST h PM`),
	"DC1":               Format(`s_ AND m_ PAST h PM`),
	"DC2":               Format(`h O m AND s_`),
	"DC3":               Format(`h O CLOCK AND s_`),
}

func (f Format) Encode(n int, t time.Time) (b []byte) {
	b = make([]byte, 0, 64)
	for _, f := range bytes.Fields(f) {
		switch f[0] {
		case '#':
			b = append(b, Numbers[n]...)
		case 'h':
			h := t.Hour() % 12
			if h == 0 {
				h = 12
			}
			b = append(b, Numbers[h]...)
			if len(f) == 2 && f[1] == '_' {
				b = append(b, Shour...)
				if h != 1 {
					b = append(b, 'S')
				}
			}
		case 'm':
			m := t.Minute()
			b = append(b, Numbers[m]...)
			if len(f) == 2 && f[1] == '_' {
				b = append(b, Sminute...)
				if m != 1 {
					b = append(b, 'S')
				}
			}
		case 's':
			s := t.Second()
			b = append(b, Numbers[s]...)
			if len(f) == 2 && f[1] == '_' {
				b = append(b, Ssecond...)
				if s != 1 {
					b = append(b, 'S')
				}
			}
		default:
			b = append(b, f...)
		}
		b = append(b, ' ')
	}
	return b[:len(b)-1]
}
