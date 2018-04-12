package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var mealCost float64
	var tipPercent float64
	var taxPercent float64

	var mealDetails []string

	for scanner.Scan() {
		mealDetails = append(mealDetails, scanner.Text())
	}

	mealCost, _ = strconv.ParseFloat(mealDetails[0], 64)
	tipPercent, _ = strconv.ParseFloat(mealDetails[1], 64)
	taxPercent, _ = strconv.ParseFloat(mealDetails[2], 64)

	tip := mealCost * (tipPercent / 100)
	tax := mealCost * (taxPercent / 100)
	totalCost := mealCost + tip + tax
	newTotalCost := int(math.RoundToEven(totalCost))
	fmt.Printf("The total meal cost is %d dollars.", newTotalCost)

}
