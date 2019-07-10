package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Egregors/adventofcode2018/utils"
)

const size = 1000

type claim struct {
	ID   int
	X, Y int
	W, H int
}

type arr [size][size]int

func (a *arr) apply(claims []claim) error {
	for _, c := range claims {
		for i := c.Y; i < c.Y+c.H; i++ {
			for j := c.X; j < c.X+c.W; j++ {
				if a[i][j] == 0 {
					a[i][j] = c.ID
				} else {
					a[i][j] = -1
				}
			}
		}
	}

	var res int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if a[i][j] == -1 {
				res++
			}
		}
	}
	return nil
}

func (a *arr) s1() int {
	var res int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if a[i][j] == -1 {
				res++
			}
		}
	}
	return res
}

func (a *arr) s2(claims []claim) (int, error) {
LOOP:
	for _, c := range claims {
		for i := c.Y; i < c.Y+c.H; i++ {
			for j := c.X; j < c.X+c.W; j++ {
				if a[i][j] != c.ID {
					continue LOOP
				}
			}
		}
		return c.ID, nil
	}
	return -1, errors.New("Nothing")
}

func main() {
	log.Println("D3")
	log.Printf("https://adventofcode.com/2018/day/3\n\n")
	data, _ := utils.ReadFileLines("d3/input.txt")

	var claims []claim
	for _, l := range data {
		line := strings.Split(l, " ")
		id, _ := strconv.Atoi(strings.Replace(line[0], "#", "", 1))
		xyStr := strings.Split(line[2], ",")
		x, _ := strconv.Atoi(xyStr[0])
		y, _ := strconv.Atoi(strings.Replace(xyStr[1], ":", "", 1))
		whStr := strings.Split(line[3], "x")
		w, _ := strconv.Atoi(whStr[0])
		h, _ := strconv.Atoi(whStr[1])
		claims = append(claims, claim{ID: id, X: x, Y: y, W: w, H: h})
	}

	start := time.Now()
	var a arr
	a.apply(claims)
	res := a.s1()
	end := time.Since(start)
	log.Printf("res 1: %d at [%v]", res, end)

	start = time.Now()
	var b arr
	b.apply(claims)
	res, err := b.s2(claims)
	if err != nil {
		panic(err)
	}
	end = time.Since(start)
	log.Printf("res 2: %d at [%v]", res, end)
}
