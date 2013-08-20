package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"sync"
	"time"
)

var Config = struct {
	Measurement int
	Start       time.Time
	Encrypted   string
	MaxSpaces   int
	MD5         string
	Low, High   int
}{}

var (
	start   string
	formats []Format
)

func init() {
	flag.IntVar(&Config.Measurement, "m", 1, "Measurement")
	flag.IntVar(&Config.MaxSpaces, "max", 3, "Max Spaces")
	flag.StringVar(&start, "s", "8:00", "Start time")
	flag.StringVar(&Config.Encrypted, "e", "", "Encrypted string (from image)")
	flag.StringVar(&Config.MD5, "md5", "", "MD5 (from Ingress)")
	flag.IntVar(&Config.Low, "l", 0, "Low format")
	flag.IntVar(&Config.High, "h", len(Formats), "High format")

	formats = make([]Format, 0, len(Formats))
	for i := range Formats {
		formats = append(formats, Formats[i].Format)
	}
}

func main() {
	var err error
	flag.Parse()

	if Config.Start, err = time.Parse("15:00", start); err != nil {
		log.Fatal(err)
	}

	var wait sync.WaitGroup
	for i := range formats[Config.Low:Config.High] {
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
				log.Printf("Time string: %s", bytes.Replace(p.TimeString, spaces[0:1], []byte{' '}, -1))
				os.Exit(0)
			}
			log.Printf("Completed Analyzing %s", f)
		}(formats[i])
	}
	wait.Wait()
}
