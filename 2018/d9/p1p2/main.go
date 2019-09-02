// https://adventofcode.com/2018/day/9
// [-] (0)
// [1]  0 (1)
// [2]  0 (2) 1
// [3]  0  2  1 (3)
// [4]  0 (4) 2  1  3
// [5]  0  4  2 (5) 1  3
// [6]  0  4  2  5  1 (6) 3
// [7]  0  4  2  5  1  6  3 (7)
// [8]  0 (8) 4  2  5  1  6  3  7
// [9]  0  8  4 (9) 2  5  1  6  3  7
// [1]  0  8  4  9  2(10) 5  1  6  3  7
// [2]  0  8  4  9  2 10  5(11) 1  6  3  7
// [3]  0  8  4  9  2 10  5 11  1(12) 6  3  7
// [4]  0  8  4  9  2 10  5 11  1 12  6(13) 3  7
// [5]  0  8  4  9  2 10  5 11  1 12  6 13  3(14) 7
// [6]  0  8  4  9  2 10  5 11  1 12  6 13  3 14  7(15)
// [7]  0(16) 8  4  9  2 10  5 11  1 12  6 13  3 14  7 15
// [8]  0 16  8(17) 4  9  2 10  5 11  1 12  6 13  3 14  7 15
// [9]  0 16  8 17  4(18) 9  2 10  5 11  1 12  6 13  3 14  7 15
// [1]  0 16  8 17  4 18  9(19) 2 10  5 11  1 12  6 13  3 14  7 15
// [2]  0 16  8 17  4 18  9 19  2(20)10  5 11  1 12  6 13  3 14  7 15
// [3]  0 16  8 17  4 18  9 19  2 20 10(21) 5 11  1 12  6 13  3 14  7 15
// [4]  0 16  8 17  4 18  9 19  2 20 10 21  5(22)11  1 12  6 13  3 14  7 15
// [5]  0 16  8 17  4 18(19) 2 20 10 21  5 22 11  1 12  6 13  3 14  7 15
// [6]  0 16  8 17  4 18 19  2(24)20 10 21  5 22 11  1 12  6 13  3 14  7 15
// [7]  0 16  8 17  4 18 19  2 24 20(25)10 21  5 22 11  1 12  6 13  3 14  7 15
package main

import (
	"fmt"
	"sort"
)

type Player struct {
	id, score int
}

func (p *Player) String() string {
	return fmt.Sprintf("[%d]: %d", p.id, p.score)
}

func (p *Player) addPints(points int) {
	p.score += points
}

func makeCircle(playersNumber int) *Circle {
	c := new(Circle)
	c.addZeroMarble()
	c.setPlayers(playersNumber)
	return c
}

type Circle struct {
	currentMarble *Marble
	marbles       map[int]*Marble
	players       []*Player
	currentPlayer *Player
}

func (c *Circle) addMarble(n int, p *Player) {
	marble := new(Marble)
	marble.n = n

	if marble.isGoal() {
		p.addPints(marble.n)
		prev := c.currentMarble.prev().prev().prev().prev().prev().prev().prev()
		p.addPints(prev.n)
		c.currentMarble = prev.next()
		c.removeMarble(prev)
		return
	}

	prev := c.currentMarble.next()
	next := c.currentMarble.next().next()

	marble.l = prev
	marble.r = next

	prev.r = marble
	next.l = marble

	c.marbles[marble.n] = marble
	c.currentMarble = marble
	c.currentPlayer = p
}

func (c *Circle) show() {
	if len(c.marbles) < 2 {
		fmt.Printf("[-] (%d)\n", c.currentMarble.n)
		return
	}

	marble := c.marbles[0]
	line := fmt.Sprintf("[%d~%d] %d", c.currentPlayer.id, c.currentPlayer.score, marble.n)
	marble = marble.next()
	for marble.n != 0 {
		if marble == c.currentMarble {
			line = line + fmt.Sprintf(" (%d) ", marble.n)
		} else {
			line = line + fmt.Sprintf("  %d ", marble.n)
		}
		marble = marble.next()
	}
	fmt.Println(line)
}

func (c *Circle) addZeroMarble() {
	marble := new(Marble)
	marble.r = marble
	marble.l = marble
	marble.n = 0

	c.marbles = make(map[int]*Marble)
	c.marbles[marble.n] = marble
	c.currentMarble = marble
}

func (c *Circle) setPlayers(playersNumber int) {
	for i := 1; i <= playersNumber; i++ {
		c.players = append(c.players, &Player{id: i, score: 0})
	}
}

func (c *Circle) removeMarble(marble *Marble) {
	if m, ok := c.marbles[marble.n]; ok {
		m.prev().r = m.next()
		m.next().l = m.prev()
	}
}

func (c *Circle) nextPlayer() *Player {
	if c.currentPlayer.id != len(c.players) {
		return c.players[c.currentPlayer.id]
	}
	return c.players[0]
}

func (c *Circle) getWinner() *Player {
	sort.Slice(c.players, func(a, b int) bool {
		return c.players[a].score > c.players[b].score
	})
	return c.players[0]
}

type Marble struct {
	n    int
	l, r *Marble
}

func (m *Marble) String() string {
	return fmt.Sprintf("{ %d <- %d -> %d }", m.prev().n, m.n, m.next().n)
}

func (m *Marble) isGoal() bool {
	return m.n%23 == 0
}

func (m *Marble) next() *Marble {
	return m.r
}

func (m *Marble) prev() *Marble {
	return m.l
}

func playGame(playersNumber, lastMarble int) int {
	circle := makeCircle(playersNumber)
	marbleNumber := 1
	player := circle.players[0]

	for marbleNumber != lastMarble {
		circle.addMarble(marbleNumber, player)
		marbleNumber++
		player = circle.nextPlayer()
	}

	return circle.getWinner().score
}

func main() {
	// playground
	score := playGame(413, 7108200)
	printGameScore(7108200, score)

	return
}

func printGameScore(i int, i2 int) {
	fmt.Printf("last marble is worth %d points: high score is %d\n", i, i2)
}
