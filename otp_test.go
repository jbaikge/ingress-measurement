package main

import (
	"bytes"
	"testing"
)

func TestOTP(t *testing.T) {
	str := []byte("TWOXOXCLOCKXFOURXMINUTESXANDXTHIRTYXEIGHTXSECONDS")
	otp := []byte("ZPCDKCAPANHLJTXFBNZEJOHZELDOJPOLPVGLXNMLPBKPNPBQJ")
	enc := []byte("SLQAYZCAOPRIOHRWYZHRDHLRBLQRGIVTGOEIBVSSIYCTPDOTB")

	if c := OTP(str, enc); !bytes.Equal(c, otp) {
		t.Logf("Expected %s", otp)
		t.Logf("Got      %s", c)
		t.Fatal("Mismatch")
	}
}
