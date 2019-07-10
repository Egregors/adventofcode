package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type node struct {
	children []node
	meta     []int
}

func (n node) String() string {
	return fmt.Sprintf("[%d %d, %v]", len(n.children), len(n.meta), n.meta)
}

func (n *node) getValue() int {
	// log.Println(n)
	if len(n.children) == 0 {
		var s int
		for _, m := range n.meta {
			s += m
		}
		return s
	}

	var vals []int

	for i := 0; i < len(n.meta); i++ {
		id := n.meta[i]
		if id > 0 && id <= len(n.children) {
			chNodeVal := n.children[id-1].getValue()
			vals = append(vals, chNodeVal)
		} else {
			vals = append(vals, 0)
		}
	}

	var s int
	for _, v := range vals {
		s += v
	}

	return s
}

func (n *node) getMeta() []int {
	var meta []int

	if len(n.children) > 0 {
		for _, ch := range n.children {
			meta = append(meta, ch.getMeta()...)
		}
	}

	meta = append(meta, n.meta...)

	return meta
}

func (n *node) getMetaSum() int {
	var s int

	for _, m := range n.getMeta() {
		s += m
	}
	return s
}

func getNode(data []int) (node, []int) {
	chCount := data[0]
	metaLen := data[1]

	data = data[2:]

	var children []node
	if chCount > 0 {
		for i := 0; i < chCount; i++ {
			var child node
			child, data = getNode(data)
			children = append(children, child)
		}
	}

	meta := data[:metaLen]
	data = data[metaLen:]

	n := node{children: children, meta: meta}
	return n, data
}

func main() {
	b, err := ioutil.ReadFile("d8/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nums := []int{}
	for _, x := range strings.Split(string(b), " ") {
		d, _ := strconv.Atoi(x)
		nums = append(nums, d)
	}

	// t := []int{0, 3, 10, 11, 12}

	n, _ := getNode(nums)
	log.Println(n.getValue())
}
