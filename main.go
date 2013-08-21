package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"runtime"
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

var start string

func init() {
	flag.IntVar(&Config.Measurement, "m", 1, "Measurement")
	flag.IntVar(&Config.MaxSpaces, "max", 3, "Max Spaces")
	flag.StringVar(&start, "s", "8:00", "Start time")
	flag.StringVar(&Config.Encrypted, "e", "", "Encrypted string (from image)")
	flag.StringVar(&Config.MD5, "md5", "", "MD5 (from Ingress)")
	flag.IntVar(&Config.Low, "l", 0, "Low format")
	flag.IntVar(&Config.High, "h", len(Formats), "High format")
}

func main() {
	var err error
	flag.Parse()

	if Config.Start, err = time.Parse("15:00", start); err != nil {
		log.Fatal(err)
	}

	log.Printf("Config: \n%+v", Config)

	go func() {
		for {
			log.Printf("Go Routines: %d", runtime.NumGoroutine())
			<-time.After(15 * time.Second)
		}
	}()

	var wait sync.WaitGroup
	for _, f := range Formats[Config.Low:Config.High] {
		log.Printf("Analyzing %-24s %s", f.Name, f.Format)
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
		}(f.Format)
	}
	wait.Wait()
}
