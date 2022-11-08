package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func findMax(numbers ...float64) float64 {
	max := math.Inf(-1)

	for _, number := range numbers {
		if max < number {
			max = number
		}
	}

	return max
}

func convertToFloat(arguments []string) ([]float64, error) {
	var numbers []float64
	var number float64
	var err error

	numbers = make([]float64, 0)

	for _, arg := range arguments {
		number, err = strconv.ParseFloat(arg, 64)

		if err != nil {
			return numbers, err
		}

		numbers = append(numbers, number)
	}

	return numbers, err
}

func main() {
	arguments := os.Args[1:]
	numbers, err := convertToFloat(arguments)

	if err != nil {
		log.Fatal(err)
	}

	max := findMax(numbers...)
	fmt.Printf("Maximum is %0.2f\n", max)
}
