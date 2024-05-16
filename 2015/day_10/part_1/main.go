package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	for i := 0; i < 40; i++ {
		input = LookAndSay(input)
	}

	fmt.Printf("answer = %d", len(input))
}

func LookAndSay(input string) string {
	result := strings.Builder{}
	counter := 1
	digit := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] == digit {
			counter++
		} else {
			ls := fmt.Sprintf("%d%c", counter, digit)
			result.WriteString(ls)
			digit = input[i]
			counter = 1
		}
	}
	result.WriteString(fmt.Sprintf("%d%c", counter, digit))
	return result.String()
}
