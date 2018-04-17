package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var number []string

	for scanner.Scan() {
		number = append(number, scanner.Text())
	}

	n64, _ := strconv.ParseUint(number[0], 10, 64)
	n := int(n64)
	sum := 1

	for sum < 11 {
		multi := n * sum
		fmt.Printf("%d x %d = %d\n", n, sum, multi)
		sum = sum + 1
	}
}
