package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	firstAscii := make([]int, 26)
	secondAscii := make([]int, 26)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fLen := 0
	for k, v := range scanner.Bytes() {
		firstAscii[v-97]++
		fLen = k
	}

	scanner.Scan()
	sLen := 0
	for k, v := range scanner.Bytes() {
		secondAscii[v-97]++
		sLen = k
	}

	if sLen != fLen {
		fmt.Println(0)
	} else {
		equal := true
		for k, v := range firstAscii {
			if v != secondAscii[k] {
				equal = false
				break
			}
		}

		if equal {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
