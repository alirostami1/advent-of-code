package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalWrapNeeded := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		dimensions := strings.Split(input, "x")
		l, _ := strconv.ParseInt(dimensions[0], 10, 64)
		w, _ := strconv.ParseInt(dimensions[1], 10, 64)
		h, _ := strconv.ParseInt(dimensions[2], 10, 64)

		minimum := min(l*w, w*h, h*l)

		wrapNeeded := 2*(l*w+w*h+h*l) + minimum
		totalWrapNeeded += wrapNeeded
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(totalWrapNeeded)
}
