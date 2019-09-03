package main

import (
	"fmt"
	"sort"
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

type PowerGroup struct {
	x, y  int
	power int
}

func NewPowerGroup(x int, y int, power int) *PowerGroup {
	return &PowerGroup{x: x, y: y, power: power}
}

type Grid struct {
	g      [300][300]*Rack
	serial int

	powerGroups []*PowerGroup
}

func (g *Grid) getBestTotalPower() (x, y int) {
	g.getPowerForAllCells()
	g.makePowerGroups()
	sort.Slice(g.powerGroups, func(i, j int) bool {
		return g.powerGroups[i].power > g.powerGroups[j].power
	})
	bestGroup := g.powerGroups[0]
	return bestGroup.x, bestGroup.y
}

func (g *Grid) getPowerForAllCells() {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			r := NewRack(g.serial, i, j)
			r.getPowerLvl()
			g.g[i][j] = r
		}
	}
}

func (g *Grid) makePowerGroups() {
	for i := 0; i < 298; i++ {
		for j := 0; j < 298; j++ {
			pg := NewPowerGroup(i, j, g.getPowerGroupLvl(i, j))
			g.powerGroups = append(g.powerGroups, pg)
		}
	}
}

func (g *Grid) getPowerGroupLvl(x int, y int) int {
	pwr := 0
	for i := x; i <= x+2; i++ {
		for j := y; j <= y+2; j++ {
			pwr += g.g[i][j].power
		}
	}
	return pwr
}

func NewGrid(serial int) *Grid {
	g := &Grid{serial: serial}
	//g.g = [300][300]*Rack{}
	return g
}

func main() {
	g := NewGrid(2866)
	fmt.Println(g.getBestTotalPower())
}
