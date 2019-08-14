package main

import (
	"fmt"
)

func main() {
	var jewels string
	var stones string
	fmt.Scanln(&jewels)
	fmt.Scanln(&stones)

	counter := 0
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(jewels); j++ {
			if jewels[j] == stones[i] {
				counter++
				break
			}
		}
	}

	fmt.Println(counter)
}
