package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello, World.")

	inputString := bufio.NewScanner(os.Stdin)
	for inputString.Scan() {
		fmt.Println(inputString.Text()) // Println will add back the final '\n'
	}
	if err := inputString.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
