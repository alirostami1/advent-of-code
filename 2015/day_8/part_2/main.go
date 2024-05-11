package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		total += CalculateEncodedUsed(input) - len(input)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(total)
}

func CalculateEncodedUsed(input string) int {
	counter := 2
	for _, c := range input {
		if c == '"' || c == '\\' {
			counter += 2
		} else {
			counter++
		}
	}

	return counter
}
