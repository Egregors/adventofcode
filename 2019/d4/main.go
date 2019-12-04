package main

import (
	"fmt"
	"strconv"
)

// 136760-595730

func isNotDecrease(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)-1; i++ {
		a1, _ := strconv.Atoi(string(s[i]))
		a2, _ := strconv.Atoi(string(s[i+1]))
		if a1 > a2 {
			return false
		}
	}
	return true
}

func hasPair(n int) bool {
	s := strconv.Itoa(n)
	var gs [][]uint8

	for i := 0; i < len(s); i++ {
		ch := s[i]
		gr := []uint8{ch}

		for j := i + 1; j < len(s); j++ {
			ch2 := s[j]

			if ch == ch2 {
				gr = append(gr, ch)
				if j == len(s)-1 {
					gs = append(gs, gr)
					i = j
				}
			} else {
				gs = append(gs, gr)
				i = j - 1
				break
			}

		}
	}
	for _, g := range gs {
		if len(g) == 2 {
			return true
		}
	}
	return false
}

func main() {
	count := 0
	for i := 100000; i <= 999999; i++ {
		if hasPair(i) && isNotDecrease(i) {
			count++
		}
	}
	fmt.Println(count)
}
