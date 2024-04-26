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

	visitedHouses := map[string]struct{}{}
	visitedHouses["0:0"] = struct{}{}
	x := 0
	y := 0
	for _, command := range input {
		switch command {
		case '<':
			x--
		case '>':
			x++
		case '^':
			y++
		case 'v':
			y--
		}

		stringified := fmt.Sprintf("%d:%d", x, y)
		visitedHouses[stringified] = struct{}{}
	}

	fmt.Print(len(visitedHouses))
}
