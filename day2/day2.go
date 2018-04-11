package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	var ii uint64
	var dd float64
	var ss string

	var mess []string

	for scanner.Scan() {
		mess = append(mess, scanner.Text())
	}

	ii, _ = strconv.ParseUint(mess[0], 10, 64)
	dd, _ = strconv.ParseFloat(mess[1], 64)
	ss = mess[2]

	fmt.Println(i + ii)
	fmt.Printf("%.1f\n", d+dd)
	fmt.Println(s + ss)

}
