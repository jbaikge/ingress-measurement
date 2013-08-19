package main

import (
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		Measurement int
		Time        time.Time
		Format      Format
		Expect      string
	}{
		{
			1, time.Date(2000, 1, 1, 14, 4, 38, 0, time.Local),
			Formats["FMinator"],
			"TWO O CLOCK FOUR MINUTES AND THIRTY EIGHT SECONDS",
		},
		{
			1, time.Date(2000, 1, 1, 15, 1, 55, 0, time.Local),
			Formats["FMinator"],
			"THREE O CLOCK ONE MINUTE AND FIFTY FIVE SECONDS",
		},
		{
			1, time.Date(2000, 1, 1, 16, 3, 20, 0, time.Local),
			Formats["FMinator"],
			"FOUR O CLOCK THREE MINUTES AND TWENTY SECONDS",
		},
		{
			1, time.Date(2000, 1, 1, 17, 2, 27, 0, time.Local),
			Formats["FMinator"],
			"FIVE O CLOCK TWO MINUTES AND TWENTY SEVEN SECONDS",
		},
		{
			1, time.Date(2000, 1, 1, 8, 2, 54, 0, time.Local),
			Formats["FMinator"],
			"EIGHT O CLOCK TWO MINUTES AND FIFTY FOUR SECONDS",
		},
		{
			1, time.Date(2000, 1, 1, 1, 1, 34, 0, time.Local),
			Formats["FCassandraSydney1"],
			"ONE MINUTE THIRTY FOUR SECONDS PAST ONE O CLOCK",
		},
		{
			1, time.Date(2000, 1, 1, 3, 3, 52, 0, time.Local),
			Formats["FCassandraSydney1"],
			"THREE MINUTES FIFTY TWO SECONDS PAST THREE O CLOCK",
		},
		{
			1, time.Date(2000, 1, 1, 14, 4, 55, 0, time.Local),
			Formats["FCassandraSydney2"],
			"TWO O CLOCK AND FOUR MINUTES AND FIFTY FIVE SECONDS",
		},
		{
			1, time.Date(2000, 1, 1, 19, 3, 55, 0, time.Local),
			Formats["FCassandraMilan1"],
			"FIFTY FIVE SECONDS AND THREE MINUTES AFTER SEVEN",
		},
		{
			3, time.Date(2000, 1, 1, 21, 3, 30, 0, time.Local),
			Formats["FCassandraMilan3"],
			"MEASUREMENT THREE IS AT NINE O THREE AND THIRTY SECONDS",
		},
	}

	for i, test := range tests {
		if got := test.Format.Encode(test.Measurement, test.Time); test.Expect != string(got) {
			t.Errorf("[%d] Expected %s; Got %s", i, test.Expect, got)
		}
	}
}
