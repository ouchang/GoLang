package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func inRange(start float64, end float64, numbers ...float64) []float64 {
	var result []float64
	result = make([]float64, 0)

	for _, number := range numbers {
		if number >= start && number <= end {
			result = append(result, number)
		}
	}

	return result
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

	if len(numbers) >= 2 {
		start := numbers[0]
		end := numbers[1]
		numbers = numbers[2:]

		result := inRange(start, end, numbers...)
		fmt.Println("Numbers with the range are: ", result)
	}
}
