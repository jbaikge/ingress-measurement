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
	// "TWO O CLOCK FOUR MINUTES AND THIRTY EIGHT SECONDS",
	// "THREE O CLOCK ONE MINUTE AND FIFTY FIVE SECONDS",
	// "FOUR O CLOCK THREE MINUTES AND TWENTY SECONDS",
	// "FIVE O CLOCK TWO MINUTES AND TWENTY SEVEN SECONDS",
	// "EIGHT O CLOCK TWO MINUTES AND FIFTY FOUR SECONDS",
	FMinator = Format(`h O CLOCK m_ AND s_`)

	// "ONE MINUTE THIRTY FOUR SECONDS PAST ONE O CLOCK",
	// "THREE MINUTES FIFTY TWO SECONDS PAST THREE O CLOCK",
	FCassandraSydney1 = Format(`m_ s_ PAST h O CLOCK`)

	// "TWO O CLOCK AND FOUR MINUTES AND FIFTY FIVE SECONDS",
	FCassandraSydney2 = Format(`h O CLOCK AND m_ AND s_`)

	// "FIFTY FIVE SECONDS AND THREE MINUTES AFTER SEVEN",
	FCassandraMilan1 = Format(`s_ AND m_ AFTER h`)

	// "MEASUREMENT THREE IS AT NINE O THREE AND THIRTY SECONDS",
	FCassandraMilan3 = Format(`MEASUREMENT # IS AT h O m AND s_`)
)

var (
	Shour   = []byte(` HOUR`)
	Sminute = []byte(` MINUTE`)
	Ssecond = []byte(` SECOND`)
)

var Formats = []Format{
	FMinator,
	FCassandraSydney1,
	FCassandraSydney2,
	FCassandraMilan1,
	FCassandraMilan3,
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
