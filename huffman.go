package main

import (
	"container/heap"
	"fmt"
	"os"
	"time"
)

type Tree interface {
	Freq() int
}

type Leaf struct {
	freq int
	value rune
}

type Node struct {
	freq        int
	left, right Tree
}

func (leaf Leaf) Freq() int {
	return leaf.freq
}

func (node Node) Freq() int {
	return node.freq
}

type treeHeap []Tree

func (th treeHeap) Len() int { return len(th) }
func (th treeHeap) Less(i, j int) bool {
	return th[i].Freq() < th[j].Freq()
}
func (th *treeHeap) Push(ele interface{}) {
	*th = append(*th, ele.(Tree))
}
func (th *treeHeap) Pop() (popped interface{}) {
	popped = (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1]
	return
}
func (th treeHeap) Swap(i, j int) { th[i], th[j] = th[j], th[i] }

func walk(tree Tree, symCodes map[rune]string, prefix []byte) {
	switch i := tree.(type) {
	case Leaf:
		symCodes[i.value] = string(prefix)
	case Node:
		prefix = append(prefix, '0')
		walk(i.left, symCodes, prefix)
		prefix = prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		walk(i.right, symCodes, prefix)
	}
}

func encode(sourceStr string) (string, map[rune]string) {
	var tree treeHeap
	symCodes := make(map[rune]string)
	symFreqs := make(map[rune]int)
	for _, ch := range sourceStr {
		symFreqs[ch]++
	}

	for ch, freq := range symFreqs {
		tree = append(tree, Leaf{freq, ch})
	}
	heap.Init(&tree)

	for len(tree) > 1 {
		a := heap.Pop(&tree).(Tree)
		b := heap.Pop(&tree).(Tree)
		heap.Push(&tree, Node{a.Freq() + b.Freq(), a, b})
	}
	if len(symFreqs) == 1 {
		for ch := range symFreqs {
			symCodes[ch] = "0"
		}
	} else {
		walk(heap.Pop(&tree).(Tree), symCodes, []byte{})
	}
	var encodedStr string
	for _, ch := range sourceStr {
		encodedStr += symCodes[ch]
	}
	fmt.Println(len(symFreqs), len(encodedStr))

	return encodedStr, symCodes
}

func decode(encodedStr string, symCodes map[rune]string) string {
	var decodedStr string
	codes := make(map[string]string)

	for ch, code := range symCodes {
		codes[code] = string(ch)
	}

	begin, end := 0, 1
	for end <= len(encodedStr) {
		strChar := encodedStr[begin:end]
		if code, ok := codes[strChar]; ok {
			decodedStr += code
			begin = end
			end++
		} else {
			end++
		}
	}
	return decodedStr
}

func main() {
	t0 := time.Now()
	sourceStr := os.Args[1]

	encodedStr, symCodes := encode(sourceStr)

	for ch, code := range symCodes {
		fmt.Println(string(ch) + ":", code)
	}
	fmt.Println(encodedStr)
	fmt.Printf("Elapsed time: %v", time.Since(t0))
}