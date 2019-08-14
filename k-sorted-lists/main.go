package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	lineCount, _ := strconv.Atoi(scanner.Text())
	size := 100
	result := make([]int, size)

	for n := 0; n < lineCount; n++ {
		scanner.Scan()
		arrayStr := strings.Fields(scanner.Text())
		aLen, _ := strconv.Atoi(arrayStr[0])

		for s := 1; s <= aLen; s++ {
			num, _ := strconv.Atoi(arrayStr[s])
			result[num]++
		}
	}

	wsCounter := 0
	for num, count := range result {

		for i := 0; i < count; i++ {
			if wsCounter != 0 {
				fmt.Print(" ")
			}

			fmt.Printf("%d", num)
			wsCounter++
		}

	}
}
