package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("d5/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	s.Scan()
	data := s.Text()

	min := check(data)
	abc := "abcdefghijklmnopqrstuvwxyz"

	for _, ch := range abc {
		d := strings.Replace(data, strings.ToLower(string(ch)), "", -1)
		d = strings.Replace(d, strings.Title(string(ch)), "", -1)
		m := check(d)

		if m < min {
			min = m
		}
	}
	log.Println(min)
}

func check(data string) int {
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
	return len(data)
}
