package main

import (
	"fmt"
	"os"
	"time"
)

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
