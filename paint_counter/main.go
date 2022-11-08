//Counts how many liters are needed to paint all the walls
//For one square meter, we need 10 literes of paint

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if height <= 0 || width <= 0 {
		err := fmt.Errorf("You gave wrong parameteres!")
		return 0, err
	}

	area := width * height
	return area / 10.0, nil
}

func main() {
	fmt.Print("Enter number of walls for painting: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	numWalls, err := strconv.Atoi(input)

	if numWalls <= 0 {
		err = fmt.Errorf("Wrong number of walls!")
		log.Fatal(err)
	}

	paintLiters := 0.0
	sumOfLiters := 0.0

	for walls := 0; walls < numWalls; walls++ {
		fmt.Print("Enter width: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		width, err := strconv.ParseFloat(input, 64)

		fmt.Print("Enter heigth: ")
		input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		heigth, err := strconv.ParseFloat(input, 64)

		paintLiters, err = paintNeeded(width, heigth)
		if err != nil {
			log.Fatal(err)
		} else {
			sumOfLiters += paintLiters
		}
	}

	fmt.Printf("We need %.2f liters of paint\n", sumOfLiters)
}
