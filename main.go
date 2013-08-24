package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	_ "net/http/pprof"
)

var Config = struct {
	Measurement int
	Start       time.Time
	Encrypted   string
	EncBytes    []byte
	MaxSpaces   int
	MD5Str      string
	MD5         []byte
	Low, High   int
	SaveAs      string
}{}

var start string

func init() {
	flag.IntVar(&Config.Measurement, "m", 1, "Measurement")
	flag.IntVar(&Config.MaxSpaces, "max", 15, "Max Spaces")
	flag.StringVar(&start, "s", "8:00", "Start time")
	flag.StringVar(&Config.Encrypted, "e", "", "Encrypted string (from image)")
	flag.StringVar(&Config.MD5Str, "md5", "", "MD5 (from Ingress)")
	flag.StringVar(&Config.SaveAs, "save", "", "File to save pre-generated MD5s")
	flag.IntVar(&Config.Low, "l", 0, "Low format")
	flag.IntVar(&Config.High, "h", len(Formats), "High format")
}

func main() {
	var err error
	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	if Config.Start, err = time.Parse("15:04", start); err != nil {
		log.Fatal(err)
	}

	if Config.MD5, err = hex.DecodeString(Config.MD5Str); err != nil {
		log.Fatal(err)
	}

	Config.EncBytes = []byte(Config.Encrypted)

	log.Printf("Config: \n%+v", Config)

	/*
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
	*/

	var (
		encoded   [36][16][][]byte
		encodings = 0
		start     = time.Now()
		// unique    = 0
	)

	for t := Config.Start.Add(5 * time.Minute); t.After(Config.Start); t = t.Add(-time.Second) {
		for _, f := range Formats {
			formatted := [][]byte{
				//f.Format.Encode(Config.Measurement, t),
				f.Format.TheHour(Config.Measurement, t),
			}
			for _, e := range formatted {
				spaceReq := len(Config.Encrypted) - len(e)
				if spaceReq < 0 {
					continue
				}

				fields := len(bytes.Fields(e))
				//fmt.Printf("%2d %2d %s\n", fields, spaceReq, e)
				encoded[spaceReq][fields] = append(encoded[spaceReq][fields], e)
				encodings++
			}
		}
	}

	// log.Printf("Pregen took: %s\n", time.Since(start))

	fmt.Printf("   ")
	for i := range encoded[0] {
		fmt.Printf("%4d", i)
	}
	fmt.Printf(" time\n")
	// for i, row := range encoded {
	// 	fmt.Printf("%3d", i)
	// 	for _, col := range row {
	// 		fmt.Printf("%4d", len(col))
	// 		unique++
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("  ^- # spaces required to make 59 chars")
	// fmt.Printf("Encoded times: %d Unique batches: %d\n", encodings, unique)

	start = time.Now()

	var buf *bufio.Writer
	if Config.SaveAs != "" {
		var file *os.File
		if file, err = os.Create(Config.SaveAs); err != nil {
			log.Fatal(err)
		}
		buf = bufio.NewWriter(file)
		defer file.Close()
	}

	// Line below: Go forward through required space-widths
	// for spacesReq, words := range encoded {
	// Two lines below: Reverse through required space-widths
	for spacesReq := len(encoded) - 1; spacesReq >= 0; spacesReq-- {
		words := encoded[spacesReq]

		rowTime := time.Now()
		fmt.Printf("%3d", spacesReq)
		for wordCount, times := range words {
			batchSize := len(times)
			fmt.Printf("%4d", batchSize)
			if batchSize == 0 {
				continue
			}
			s := NewSpacer(wordCount, spacesReq)
			for sp := s.Iter(); sp > 0; sp = s.Next() {
				//fmt.Printf("\n"+fmt.Sprintf("%%0%dx", wordCount+1)+"\n", sp)
				spacers := make([][]byte, wordCount+1)
				spacers[0] = spaces[:sp&0xF]
				for i := 0; i < wordCount; i++ {
					l := sp & (0xF << ((uint(i) + 1) * 4)) / (0xF << (uint(i) * 4))
					spacers[i+1] = spaces[:l]
				}
				//fmt.Println(wordCount, spacesReq, spacers)
				var wait sync.WaitGroup
				for t := range times {
					wait.Add(1)
					go func(b []byte) {
						defer wait.Done()
						str := make([]byte, 0, len(Config.Encrypted))
						fields := bytes.Fields(b)

						str = str[:0]
						for i := range fields {
							str = append(str, spacers[i]...)
							str = append(str, fields[i]...)
						}
						str = append(str, spacers[wordCount]...)

						otp := OTP(str, Config.EncBytes)
						hasher := md5.New()
						hasher.Write(otp)
						hash := hasher.Sum(nil)

						// Save stuff
						if Config.SaveAs != "" {
							fmt.Fprintf(buf, "%x %s\n", hash, str)
							return
						}

						// Actually searching...
						if bytes.Equal(hash, Config.MD5) {
							fmt.Printf("\n\nFound: %s\nTook:  %s\n",
								bytes.Replace(b, spaces[:1], []byte{' '}, -1),
								time.Since(start),
							)
							os.Exit(0)
						}
					}(times[t])
				}
				wait.Wait()
			}
		}
		fmt.Printf(" %s\n", time.Since(rowTime))
	}
	log.Printf("Testing took: %s\n", time.Since(start))
}
