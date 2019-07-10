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
	return fmt.Sprintf("Children: %d, META: %v", len(n.children), n.meta)
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
	// log.Printf("chCount: %d metaLen: %d\n", chCount, metaLen)

	data = data[2:]
	// log.Printf("data tail: %v", data)

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
	// log.Printf("NODE is Ready: %v", n)
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

	n, _ := getNode(nums)
	log.Println(n.getMetaSum())
}
