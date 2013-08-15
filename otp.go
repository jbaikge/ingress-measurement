package main

func OTP(a, b []byte) (c []byte) {
	if len(a) != len(b) {
		return
	}

	c = make([]byte, len(a))
	var Δ byte
	for i := range a {
		if b[i] < a[i] {
			Δ = 26 - (a[i] - b[i])
		} else {
			Δ = b[i] - a[i]
		}
		c[i] = 'A' + Δ
	}
	return
}
