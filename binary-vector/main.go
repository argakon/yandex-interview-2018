package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	index := 0

	var lineCountStr string
	fmt.Scanln(&lineCountStr)
	lineCount, _ := strconv.ParseInt(lineCountStr, 10, 32)
	if lineCount < 1 {
		fmt.Println(0)
	} else {
		buf := make([]int, lineCount)

		var n int64
		for n = 0; n < lineCount; n++ {
			var numStr string
			fmt.Scanln(&numStr)
			num, _ := strconv.ParseInt(numStr, 10, 32)

			if num == 1 {
				buf[index]++
			} else {
				index++
			}
		}

		sort.Ints(buf)

		fmt.Println(buf[len(buf)-1])
	}
}
