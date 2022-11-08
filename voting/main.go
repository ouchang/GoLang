package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var nameMap map[string]int
	nameMap = make(map[string]int)

	file, err := os.Open("votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		name := scanner.Text()
		nameMap[name]++
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	for key, value := range nameMap {
		fmt.Println(key, ":", value)
	}
}
