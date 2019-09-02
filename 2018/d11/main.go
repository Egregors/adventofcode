package main

import (
	"fmt"
	"strconv"
)

func getHundredsDigit(n int) int {
	hs := strconv.Itoa(n / 100)
	hsStr := strconv.Itoa(n / 100)[len(hs)-1:]
	res, err := strconv.Atoi(hsStr)
	if err != nil {
		panic(err)
	}

	return res
}

type Rack struct {
	serial int
	x, y   int

	ID    int
	power int
}

func NewRack(serial int, x int, y int) *Rack {
	r := &Rack{serial: serial, x: x, y: y}
	r.ID = r.genID()

	return r
}

func (r *Rack) getPowerLvl() int {
	r.power = r.ID * r.y
	r.power += r.serial
	r.power *= r.ID

	h := getHundredsDigit(r.power)
	r.power = h
	r.power -= 5

	return r.power
}

func (r *Rack) genID() int {
	return r.x + 10
}

type Grid struct {
	g      [300][300]*Rack
	serial int
}

func (g *Grid) getBestTotalPower() (x, y int) {
	g.getPowerForAllCells()

	// todo: this
	return 0, 0
}

func (g *Grid) getPowerForAllCells() {
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			r := NewRack(g.serial, i, j)
			r.getPowerLvl()
			g.g[i][j] = r
		}
	}
}

func NewGrid(serial int) *Grid {
	return &Grid{serial: serial}
}

func main() {
	g := NewGrid(18)
	fmt.Println(g.getBestTotalPower())
}
