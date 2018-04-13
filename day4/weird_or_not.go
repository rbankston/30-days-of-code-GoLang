package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var numbers []string

	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	n64, _ := strconv.ParseUint(numbers[0], 10, 64)
	n := int(n64)

	if InBetween(n, 2, 5) && n%2 == 0 {
		fmt.Println("Not Weird")
	} else if InBetween(n, 6, 20) && n%2 == 0 {
		fmt.Println("Weird")
	} else if n > 20 && n%2 == 0 {
		fmt.Println("Not Weird")
	} else if n%2 == 0 {
		fmt.Println("Not Weird")
	} else {
		fmt.Println("Weird")
	}

}

func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
