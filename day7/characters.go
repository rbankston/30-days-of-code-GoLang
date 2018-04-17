package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	n64, _ := strconv.ParseUint(input[0], 10, 64)
	limit := int(n64)
	times := 0

	for times < limit {
		word := input[times+1]
		var b = []byte(word)
		var firstportion string
		var secondportion string
		for i := range b {
			if (i%2) == 0 || i == 0 {
				firstportion = firstportion + string(b[i])
			} else {
				secondportion = secondportion + string(b[i])
			}
		}
		fmt.Printf("%s %s\n", firstportion, secondportion)
		times = times + 1
	}
}
