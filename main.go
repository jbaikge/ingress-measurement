package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

var Config = struct {
	Measurement int
	Start       time.Time
	Encrypted   string
	MD5         string
}{}

var start string

func init() {
	flag.IntVar(&Config.Measurement, "m", 1, "Measurement")
	flag.StringVar(&start, "s", "8:00", "Start time")
	flag.StringVar(&Config.Encrypted, "e", "", "Encrypted string (from image)")
	flag.StringVar(&Config.MD5, "md5", "", "MD5 (from Ingress)")
}

func main() {
	var err error
	flag.Parse()

	if Config.Start, err = time.Parse("15:00", start); err != nil {
		log.Fatal(err)
	}

	var wait sync.WaitGroup
	for i := range Formats {
		wait.Add(1)
		go func(f Format) {
			defer wait.Done()
			p, err := NewPackage(
				f,
				Config.Measurement,
				Config.Start,
				Config.Encrypted,
				Config.MD5,
			)
			if err != nil {
				log.Fatal(err)
			}
			if p.Find() {
				log.Printf("Found OTP: %s", p.OTP)
			}
		}(Formats[i])
	}
	wait.Wait()
}
