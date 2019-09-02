// https://adventofcode.com/2018/day/10
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := strings.Split(string(c), "\n")
	return data, nil
}

func NewScreen(points []string) *Screen {
	s := new(Screen)
	for _, l := range points {
		s.points = append(s.points, NewPoint(l))
	}
	return s
}

type Screen struct {
	points []*Point
}

func (s *Screen) show() (string, bool) {
	scr := ""
	xTo, xFrom := s.getMaxAndMinXs()
	yTo, yFrom := s.getMaxAndMinYs()
	// screen size
	w := xTo - xFrom
	h := yTo - yFrom
	if w > 300 || h > 300 {
		return fmt.Sprintf("too big: %d x %d", w, h), false
	}

	for y := yFrom; y <= yTo; y++ {
		for x := xFrom; x <= xTo; x++ {
			if s.isPointExist(x, y) {
				scr += "#"
			} else {
				scr += "."
			}
		}
		scr += "\n"
	}
	return scr, true
}

func (s *Screen) getMaxAndMinXs() (int, int) {
	sort.Slice(s.points, func(i, j int) bool {
		return s.points[i].position.x > s.points[j].position.x
	})
	return s.points[0].position.x, s.points[len(s.points)-1].position.x
}

func (s *Screen) getMaxAndMinYs() (int, int) {
	sort.Slice(s.points, func(i, j int) bool {
		return s.points[i].position.y > s.points[j].position.y
	})
	return s.points[0].position.y, s.points[len(s.points)-1].position.y
}

func (s *Screen) isPointExist(x int, y int) bool {
	for _, p := range s.points {
		if p.position.x == x && p.position.y == y {
			return true
		}
	}

	return false
}

func (s *Screen) makeStep(sec int) {
	if sec == 0 {
		sec = 1
	}
	for _, p := range s.points {
		p.move(sec)
	}

}

type Position struct {
	x, y int
}

type Velocity struct {
	horizontal, vertical int
}

func NewPoint(raw string) *Point {
	p := new(Point)
	startPosition, err := getPositionFromRaw(raw)
	if err != nil {
		panic(err)
	}
	startVelocity, err := getVelocityFromRaw(raw)
	if err != nil {
		panic(err)
	}

	p.startPosition = *startPosition
	p.startVelocity = *startVelocity

	p.position = p.startPosition
	p.velocity = p.startVelocity

	return p
}

type Point struct {
	startPosition, position Position
	startVelocity, velocity Velocity
}

func (p *Point) String() string {
	return fmt.Sprintf(
		"p: %d, %d v: %d, %d",
		p.position.x,
		p.position.y,
		p.velocity.horizontal,
		p.velocity.vertical,
	)
}

func (p *Point) move(sec int) {
	for i := 0; i < sec; i++ {
		p.position.y = p.position.y + p.velocity.vertical
		p.position.x = p.position.x + p.velocity.horizontal
	}
}

func wait() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	_, _ = buf.ReadBytes('\n')
}

func main() {
	points, _ := readLines("input.txt")
	s := NewScreen(points)
	startSec := 1

	for {
		s.makeStep(1)
		if scr, ok := s.show(); !ok {
			fmt.Printf("\rOn %d sec", startSec)
		} else {
			fmt.Println(scr)
			wait()
		}
		startSec++
	}
}

func getPositionFromRaw(s string) (*Position, error) {
	openIdx, err := findSubstrIdx(s, "<")
	if err != nil {
		return nil, err
	}
	closeIdx, err := findSubstrIdx(s, ">")
	if err != nil {
		return nil, err
	}

	xAndY := strings.Split(s[openIdx+1:closeIdx], ",")

	x, errX := strconv.Atoi(strings.TrimSpace(xAndY[0]))
	y, errY := strconv.Atoi(strings.TrimSpace(xAndY[1]))

	if errX != nil || errY != nil {
		return nil, errors.New("can't got x or y")
	}
	return &Position{x: x, y: y}, nil
}

func getVelocityFromRaw(s string) (*Velocity, error) {

	vIdx, err := findSubstrIdx(s, "v")
	if err != nil {
		return nil, err
	}

	s = s[vIdx:]

	openIdx, err := findSubstrIdx(s, "<")
	if err != nil {
		return nil, err
	}
	closeIdx, err := findSubstrIdx(s, ">")
	if err != nil {
		return nil, err
	}

	xAndY := strings.Split(s[openIdx+1:closeIdx], ",")

	x, errX := strconv.Atoi(strings.TrimSpace(xAndY[0]))
	y, errY := strconv.Atoi(strings.TrimSpace(xAndY[1]))

	if errX != nil || errY != nil {
		return nil, errors.New("can't got x or y")
	}
	return &Velocity{horizontal: x, vertical: y}, nil
}

func findSubstrIdx(s, subs string) (int, error) {
	for idx, ch := range s {
		if string(ch) == subs {
			return idx, nil
		}
	}
	return -1, errors.New(fmt.Sprintf("can't find %s in %s", subs, s))
}
