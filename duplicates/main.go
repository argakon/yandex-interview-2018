package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	lineCountStr := scanner.Text()

	lineCount, _ := strconv.Atoi(lineCountStr)
	if lineCount < 1 || lineCount > 1000000 {
		fmt.Println()
	} else {
		var last int
		last = 0
		var n int
		for n = 0; n < lineCount; n++ {
			scanner.Scan()
			numStr := scanner.Text()
			num, _ := strconv.Atoi(numStr)
			if n == 0 && num == 0 {
				fmt.Printf("%d\n", num)
			} else if last != num {
				last = num
				fmt.Printf("%d\n", num)
			}
		}
	}
}
