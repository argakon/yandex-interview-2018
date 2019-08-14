package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gen(n int, cOpen int, cClose int, s string) {
	if cOpen+cClose == 2*n {
		fmt.Println(s)
		return
	}

	if cOpen < n {
		gen(n, cOpen+1, cClose, s+"(")
	}

	if cOpen > cClose {
		gen(n, cOpen, cClose+1, s+")")
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	cOpen := 0
	cClose := 0
	s := ""

	gen(n, cOpen, cClose, s)
}
