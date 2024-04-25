package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	totalRibbonNeeded := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		dimensions := strings.Split(input, "x")
		l, _ := strconv.ParseInt(dimensions[0], 10, 64)
		w, _ := strconv.ParseInt(dimensions[1], 10, 64)
		h, _ := strconv.ParseInt(dimensions[2], 10, 64)

		arr := []int64{l, w, h}
		slices.Sort(arr)

		ribbon := arr[0]*2 + arr[1]*2 + l*w*h

		totalRibbonNeeded += ribbon
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(totalRibbonNeeded)
}
