package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("d2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var words []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	for i := range words {
		for j := range words[1:] {
			common, ok := comp(words[i], words[j])
			if ok {
				log.Println(common)
				return
			}
		}
	}
}

func comp(a, b string) (string, bool) {
	idx := -1

	if len(a) != len(b) {
		return "", false
	}

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		if idx >= 0 {
			return "", false
		}
		idx = i
	}

	if idx < 0 {
		return "", false
	}

	return a[:idx] + a[idx+1:], true
}
