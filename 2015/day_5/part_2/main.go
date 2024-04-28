package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counter := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if IsCompliant(input) {
			counter++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(counter)
}

func IsCompliant(str string) bool {
	hasTwoPairOfTwos := false
	hasRepeatingChar := false

	runedStr := []rune(str)
	for i := 2; i < len(str); i++ {
		if runedStr[i] == runedStr[i-2] {
			hasRepeatingChar = true
		}

		for j := i - 2; j > 0; j-- {
			if runedStr[i] == runedStr[j] && runedStr[i-1] == runedStr[j-1] {
				hasTwoPairOfTwos = true
			}
		}
		if hasRepeatingChar && hasTwoPairOfTwos {
			return true
		}
	}

	return false
}
