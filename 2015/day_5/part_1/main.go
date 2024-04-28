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
	vowelCounter := 0
	hasDoubleCharacter := false
	hasDisallowedCharacters := false

	for i, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCounter++
		case 'b':
			if i > 0 && str[i-1] == 'a' {
				hasDisallowedCharacters = true
			}
		case 'd':
			if i > 0 && str[i-1] == 'c' {
				hasDisallowedCharacters = true
			}
		case 'q':
			if i > 0 && str[i-1] == 'p' {
				hasDisallowedCharacters = true
			}
		case 'y':
			if i > 0 && str[i-1] == 'x' {
				hasDisallowedCharacters = true
			}
		}
		if i > 0 && char == []rune(str)[i-1] {
			hasDoubleCharacter = true
		}
	}

	if vowelCounter >= 3 && hasDoubleCharacter && !hasDisallowedCharacters {
		return true
	}

	return false
}
