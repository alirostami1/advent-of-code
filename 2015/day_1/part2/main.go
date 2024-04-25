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
	for i, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			fmt.Print(i + 1)
			return
		}
	}
}
