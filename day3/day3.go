package main

import (
	"aoc23/util"
	"fmt"
	"os"
)

type number struct {
	value  int
	startX int
	endX   int
	y      int
}

type symbol struct {
	value rune
	x     int
	y     int
}

func main() {
	lines, err := util.ReadData("day3")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	numbers, symbols := parseLines(lines)

	var partNumberValues []int
	for _, number := range numbers {
		for _, symbol := range symbols {
			if (symbol.x >= number.startX-1) && (symbol.x <= number.endX+1) && (symbol.y >= number.y-1) && (symbol.y <= number.y+1) {
				partNumberValues = append(partNumberValues, number.value)
			}
		}
	}

	fmt.Println("First solution:", util.Sum(partNumberValues))
}

func parseLines(lines []string) ([]number, []symbol) {
	var numbers []number
	var symbols []symbol
	for y, line := range lines {
		lineNumbers, lineSymbols := parseLine(line, y)
		numbers = append(numbers, lineNumbers...)
		symbols = append(symbols, lineSymbols...)
	}
	return numbers, symbols
}

func parseLine(line string, y int) ([]number, []symbol) {
	var numbers []number
	var symbols []symbol
	var currentNumber *number
	for x, char := range line {
		switch char {
		case '.':
			if currentNumber != nil {
				numbers = append(numbers, *currentNumber)
				currentNumber = nil
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if currentNumber != nil {
				currentNumber.endX += 1
				currentNumber.value = currentNumber.value*10 + int(char-'0')
			} else {
				currentNumber = new(number)
				currentNumber.startX = x
				currentNumber.endX = x
				currentNumber.y = y
				currentNumber.value = int(char - '0')
			}
		default:
			if currentNumber != nil {
				numbers = append(numbers, *currentNumber)
				currentNumber = nil
			}
			symbols = append(symbols, symbol{value: char, x: x, y: y})
		}
	}
	if currentNumber != nil {
		numbers = append(numbers, *currentNumber)
	}
	return numbers, symbols
}
