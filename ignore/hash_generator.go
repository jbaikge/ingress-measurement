package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	KeyLength = 59
	MinLetter = 'A'
	MaxLetter = 'Z'
)

type Key [KeyLength]byte

func NewKey(init byte) (k *Key) {
	k = new(Key)
	for i := 0; i < len(k); i++ {
		k[i] = init
	}
	return
}

func (k Key) Bytes() []byte {
	return k[:]
}

func (k *Key) Increment() {
	k[0]++

	if k[0] <= MaxLetter {
		return
	}

	for c := 0; c < len(k)-1; c++ {
		if k[c] > MaxLetter {
			k[c] = MinLetter
			k[c+1]++
			continue
		}
		break
	}
}

func (k Key) Max() bool {
	if k[len(k)-1] != MaxLetter {
		return false
	}

	for i := 0; i < len(k)-1; i++ {
		if k[i] < MaxLetter {
			return false
		}
	}

	return true
}

//const Max = 304388002238065446743317502658335238267777393435301483051006458552163515529718398976

func Generate(ch chan Key) {
	k := NewKey(MinLetter)

	for ; !k.Max(); k.Increment() {
		ch <- *k
	}
	ch <- *k
	close(ch)
}

func main() {
	start := time.Now()
	defer fmt.Printf("Took: %s\n", time.Since(start))

	var wait sync.WaitGroup
	ch := make(chan Key)
	w := make(chan Key, 1024)

	wait.Add(1)

	go func(w chan Key) {
		f, err := os.Create("/home/jake/rainbow")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()

		nl := []byte{'\n'}
		buf := bufio.NewWriter(f)
		hasher := md5.New()

		for k := range w {
			hasher.Write(k.Bytes())
			fmt.Fprintf(buf, "%X ", hasher.Sum(nil))
			buf.Write(k.Bytes())
			buf.Write(nl)
		}

		buf.Flush()
		wait.Done()
	}(w)

	s, n := time.Now(), 0
	go Generate(ch)
	for k := range ch {
		w <- k
		if n++; n == 100000 {
			fmt.Printf("%d in %s\n", n, time.Since(s))
			s, n = time.Now(), 0
		}
	}
	close(w)
	wait.Wait()
}
