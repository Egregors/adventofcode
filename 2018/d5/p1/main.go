package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.Open("d5/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	s.Scan()
	data := s.Text()

	var done bool

	for !done {
		done = true
		for i := 0; i < len(data)-1; i++ {
			p1 := string(data[i])
			p2 := string(data[i+1])

			if strings.ToLower(p1) != strings.ToLower(p2) {
				continue
			}

			if p1 == strings.Title(p1) {
				if p2 == strings.ToLower(p2) {
					data = data[:i] + data[i+2:]
					done = false
				}
			} else {
				if p2 == strings.Title(p2) {
					data = data[:i] + data[i+2:]
					done = false
				}
			}
		}
	}
	fin := time.Since(start)
	log.Printf("%d at %v", len(data), fin)
}
