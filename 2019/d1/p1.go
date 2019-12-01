// https://adventofcode.com/2019/day/1
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFuel(mass int) int {
	m := mass/3 - 2
	if m <= 0 {
		return 0
	}
	return m + getFuel(m)
}

func main() {
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	var sum int
	for _, l := range strings.Split(string(b), "\n") {
		mass, _ := strconv.Atoi(l)
		sum += getFuel(mass)
	}

	fmt.Println(sum)
}
