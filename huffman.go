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

func encode(sourceString string) (string, map[rune]string) {
	var tree treeHeap
	var encodedString string

	symbolCodes := make(map[rune]string)
	symbolFrequencies := make(map[rune]int)

	for _, char := range sourceString {
		symbolFrequencies[char]++
	}

	for char, freq := range symbolFrequencies {
		tree = append(tree, Leaf{freq, char})
	}
	heap.Init(&tree)

	for len(tree) > 1 {
		left := heap.Pop(&tree).(Tree)
		tight := heap.Pop(&tree).(Tree)
		heap.Push(&tree, Node{left.Freq() + tight.Freq(), left, tight})
	}
	if len(symbolFrequencies) == 1 {
		for char := range symbolFrequencies {
			symbolCodes[char] = "0"
		}
	} else {
		walk(heap.Pop(&tree).(Tree), symbolCodes, []byte{})
	}

	for _, char := range sourceString {
		encodedString += symbolCodes[char]
	}

	return encodedString, symbolCodes
}

func decode(encodedString string, symbolCodes map[rune]string) string {
	var decodedString string
	symbolCodeAsKey := make(map[string]rune)

	for char, code := range symbolCodes {
		symbolCodeAsKey[code] = char
	}

	begin, end := 0, 1
	for end <= len(encodedString) {
		substringAsCode := encodedString[begin:end]
		if code, ok := symbolCodeAsKey[substringAsCode]; ok {
			decodedString += string(code)
			begin = end
			end++
			continue
		}
		end++
	}
	return decodedString
}

func main() {
	t0 := time.Now()
	sourceString := os.Args[1]

	encodedString, symbolCodes := encode(sourceString)

	for char, code := range symbolCodes {
		fmt.Println(string(char) + ":", code)
	}
	fmt.Println(encodedString)
	fmt.Println(decode(encodedString, symbolCodes))
	fmt.Printf("Elapsed time: %v", time.Since(t0))
}