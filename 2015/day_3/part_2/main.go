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
	sx, sy, rx, ry := 0, 0, 0, 0
	for i, command := range input {
		if i%2 == 0 {
			switch command {
			case '<':
				sx--
			case '>':
				sx++
			case '^':
				sy++
			case 'v':
				sy--
			}
			stringified := fmt.Sprintf("%d:%d", sx, sy)
			visitedHouses[stringified] = struct{}{}
		} else {
			switch command {
			case '<':
				rx--
			case '>':
				rx++
			case '^':
				ry++
			case 'v':
				ry--
			}
			stringified := fmt.Sprintf("%d:%d", rx, ry)
			visitedHouses[stringified] = struct{}{}
		}
	}

	fmt.Print(len(visitedHouses))
}
