package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFloats(fileName string) ([]float64, error) {
	var numbers []float64 //slice
	numbers = make([]float64, 0)
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	idx := 0

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)

		numbers = append(numbers, number)

		if err != nil {
			return nil, err
		}

		idx++
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}

func main() {
	beef, err := getFloats("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, sample := range beef {
		sum += sample
	}

	average := sum / (float64(len(beef)))
	fmt.Printf("Average beef weight is %0.2f\n", average)
}
