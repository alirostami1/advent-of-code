package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	grid := [1000][1000]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Bytes()

		InterpretCommand(input, &grid)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			total += grid[i][j]
		}
	}
	fmt.Print(total)
}

func InterpretCommand(input []byte, grid *[1000][1000]int) {
	if bytes.Compare(input[:7], []byte("turn on")) == 0 {
		firstX, firstY, secondX, secondY := extractAreaFromInput(input, len("turn on"))
		minX := min(firstX, secondX)
		maxX := max(firstX, secondX)
		minY := min(firstY, secondY)
		maxY := max(firstY, secondY)
		for i := minY; i <= maxY; i++ {
			for j := minX; j <= maxX; j++ {
				grid[i][j]++
			}
		}
	}
	if bytes.Compare(input[:7], []byte("turn of")) == 0 {
		firstX, firstY, secondX, secondY := extractAreaFromInput(input, len("turn off"))
		minX := min(firstX, secondX)
		maxX := max(firstX, secondX)
		minY := min(firstY, secondY)
		maxY := max(firstY, secondY)
		for i := minY; i <= maxY; i++ {
			for j := minX; j <= maxX; j++ {
				if grid[i][j] > 0 {
					grid[i][j]--
				}
			}
		}
	}
	if bytes.Compare(input[:7], []byte("toggle ")) == 0 {
		firstX, firstY, secondX, secondY := extractAreaFromInput(input, len("toggle"))
		minX := min(firstX, secondX)
		maxX := max(firstX, secondX)
		minY := min(firstY, secondY)
		maxY := max(firstY, secondY)
		for i := minY; i <= maxY; i++ {
			for j := minX; j <= maxX; j++ {
				grid[i][j] += 2
			}
		}
	}
}

func extractAreaFromInput(input []byte, commandLength int) (firstX, firstY, secondX, secondY int64) {
	splitted := bytes.Split(input[commandLength+1:], []byte{' '})

	first := bytes.Split(splitted[0], []byte{','})
	second := bytes.Split(splitted[2], []byte{','})

	firstY, _ = strconv.ParseInt(string(first[1]), 10, 64)
	firstX, _ = strconv.ParseInt(string(first[0]), 10, 64)
	secondY, _ = strconv.ParseInt(string(second[1]), 10, 64)
	secondX, _ = strconv.ParseInt(string(second[0]), 10, 64)

	return firstX, firstY, secondX, secondY
}
