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
		total += len(input) - CalculateMemoryUsed(input)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(total)
}

func CalculateMemoryUsed(input string) int {
	counter := 0
	for i := 1; i < len(input)-1; {
		if input[i] == '\\' {
			if input[i+1] == 'x' {
				i += 4
			} else {
				i += 2
			}
		} else {
			i++
		}
		counter++
	}

	return counter
}
