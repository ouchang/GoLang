//challenge for players to guess the random generated number

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	success := false

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1 //Range between 1 and 100

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, " guesses left")

		fmt.Print("Enter your guess: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if guess == target {
			fmt.Println("Bravo! You guessed the target!")
			success = true
			break
		} else if guess > target {
			fmt.Println("Oops. Your guess is too HIGH")
		} else {
			fmt.Println("Oops. Your guess is too LOW")
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my taget. It was: ", target)
	}
}
