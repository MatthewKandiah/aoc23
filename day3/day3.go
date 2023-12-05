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

type gear struct {
	left  int
	right int
}

func (g gear) product() int {
	return g.left * g.right
}

func main() {
	lines, err := util.ReadData("day3")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	numbers, symbols := parseLines(lines)

	var partNumbers []number
	for _, number := range numbers {
		for _, symbol := range symbols {
			if (symbol.x >= number.startX-1) && (symbol.x <= number.endX+1) && (symbol.y >= number.y-1) && (symbol.y <= number.y+1) {
				partNumbers = append(partNumbers, number)
			}
		}
	}

	var partNumberValues []int
	for _, part := range partNumbers {
		partNumberValues = append(partNumberValues, part.value)
	}

	var gears []gear
	for _, symbol := range symbols {
		if symbol.value != '*' {
			continue
		}
		// since '-' is a symbol, we can't get negative numbers, so this is a safe sentinel value
		// found the problem, I've assumed they need to be in a line, the example makes it clear this is wrong, it's a gear if it's adjacent to exactly two *s in any direction!
		left, right := -1, -1
		for _, number := range partNumbers {
			if number.y == symbol.y && number.endX == symbol.x-1 {
				left = number.value
			}
			if number.y == symbol.y && number.startX == symbol.x+1 {
				right = number.value
			}
			if left != -1 && right != -1 {
				gears = append(gears, gear{left: left, right: right})
				break
			}
		}
	}

	for idx, gear := range gears {
		fmt.Println(idx, gear)
	}

	var gearProducts []int
	for _, gear := range gears {
		gearProducts = append(gearProducts, gear.product())
	}

	fmt.Println("First solution:", util.Sum(partNumberValues))
	fmt.Println("Second solution:", util.Sum(gearProducts))
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
