package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var gates map[string]string = make(map[string]string)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		ParseInput(input)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	output := RecursiveCalculator("a")
	log.Print(output)
}

func ParseInput(input string) {
	splitted := strings.Split(input, " ")
	gateName := splitted[len(splitted)-1]
	gates[gateName] = input[:len(input)-len(gateName)-4]
}

func RecursiveCalculator(endGate string) uint16 {
	inputFormula := gates[endGate]
	splittedFormula := strings.Split(inputFormula, " ")

	output := uint16(0)
	switch len(splittedFormula) {
	case 1:
		log.Printf("found case 1 '%s' = '%s'", endGate, inputFormula)
		// it is definitely a number or a gate
		parsedInputNumber, err := strconv.ParseUint(inputFormula, 10, 16)
		inputNumber := uint16(parsedInputNumber)
		if err != nil {
			log.Printf("error %s", err)
			output = RecursiveCalculator(inputFormula)
		} else {
			output = inputNumber
		}
	case 2:
		log.Printf("found case 2 %s = %v", endGate, inputFormula)
		parsedInputNumber, err := strconv.ParseUint(splittedFormula[1], 10, 16)
		inputNumber := uint16(parsedInputNumber)
		if err != nil {
			n := RecursiveCalculator(splittedFormula[1])
			output = ^n
		} else {
			output = ^inputNumber
		}
	case 3:
		// it could be AND, OR, LSHIFT, and RSHIFT
		log.Printf("found case 3 %s = %v", endGate, inputFormula)
		parsedLeftNumber, err := strconv.ParseUint(splittedFormula[0], 10, 16)
		leftNumber := uint16(parsedLeftNumber)
		if err != nil {
			leftNumber = RecursiveCalculator(splittedFormula[0])
		}
		parsedRightNumber, err := strconv.ParseUint(splittedFormula[2], 10, 16)
		rightNumber := uint16(parsedRightNumber)
		if err != nil {
			rightNumber = RecursiveCalculator(splittedFormula[2])
		}
		switch splittedFormula[1] {
		case "AND":
			output = leftNumber & rightNumber
		case "OR":
			output = leftNumber | rightNumber
		case "LSHIFT":
			output = leftNumber << rightNumber
		case "RSHIFT":
			output = leftNumber >> rightNumber
		}
	}

	gates[endGate] = strconv.FormatUint(uint64(output), 10)
	return output
}
