package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	input := scanner.Text()

	floor := 0
	for _, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
	}

	fmt.Print(floor)
}
