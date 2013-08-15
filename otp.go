package main

func OTP(a, b []byte) (c []byte) {
	if len(a) != len(b) {
		return
	}

	c = make([]byte, len(a))
	for i := range a {
		c[i] = a[i]
	}
	return
}
