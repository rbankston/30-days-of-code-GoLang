package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var size int
	fmt.Scan(&size)
	scanner := bufio.NewReader(os.Stdin)
	buffer, _ := scanner.ReadString('\n')
	buffer = strings.Trim(buffer, "\r\n")
	bufferSlice := strings.Split(buffer, " ")
	for i := size - 1; i >= 0; i-- {
		fmt.Printf("%s ", bufferSlice[i])
	}
}
