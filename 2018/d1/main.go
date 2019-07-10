package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.Println("D1")
	log.Printf("https://adventofcode.com/2018/day/1\n\n")

	content, err := ioutil.ReadFile("d1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(content), "\n")

	start := time.Now()
	result_1, err := solve(data)
	if err != nil {
		log.Fatal(err)
	}
	elapsed_1 := time.Since(start)

	start = time.Now()
	result_2, err := solve_2(data)
	if err != nil {
		log.Fatal(err)
	}
	elapsed_2 := time.Since(start)

	log.Printf("Result part 1: %d at %s", result_1, elapsed_1)
	log.Printf("Result part 2: %d at %s", result_2, elapsed_2)
}

const (
	plus  = 43
	minus = 45
)

func solve(data []string) (int, error) {

	var (
		res  int
		line string
	)

	for i := 0; i < len(data); i++ {
		line = data[i]
		if data[i][0] == plus {
			val, _ := strconv.Atoi(line[1:])
			res += val

		} else if data[i][0] == minus {
			val, _ := strconv.Atoi(line[1:])
			res -= val
		} else {
			return 1, errors.New("Bad sign")
		}
	}

	return res, nil
}

func solve_2(data []string) (int, error) {
	var (
		res  int
		line string
	)

	i := 0
	fs := map[int]bool{0: true}

	for {
		line = data[i]
		if data[i][0] == plus {
			val, _ := strconv.Atoi(line[1:])
			res += val
		} else if data[i][0] == minus {
			val, _ := strconv.Atoi(line[1:])
			res -= val
		} else {
			return 1, errors.New("Bad sign")
		}

		if fs[res] {
			return res, nil
		} else {
			fs[res] = true
		}
		i++
		if i == len(data) {
			i = 0
		}
	}
}
