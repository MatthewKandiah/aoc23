package main

import (
	utils "aoc23/util"
	"fmt"
	"os"
)

type Number struct {
	value  int
	xStart int
	xEnd   int
	y      int
}

type Symbol struct {
	value rune
	x     int
	y     int
}

type PartNumber struct {
	value int
	numStartX int
	numEndX int
	numY int
	symX int
	symY int
}

func main() {
	lines, err := utils.ReadData("day3.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	charMatrix := linesToCharMatrix(lines)

	numbers, symbols := parseCharMatrix(charMatrix)

	var partNumbers []PartNumber
	for _, number := range numbers {
		for _, symbol := range symbols {
			isInNeighbouringColumn := (symbol.x >= number.xStart-1) && (symbol.x <= number.xEnd+1)
			isInNeighbouringRow := (symbol.y >= number.y-1) && (symbol.y <= number.y+1)
			if isInNeighbouringColumn && isInNeighbouringRow {
				partNumbers = append(partNumbers, PartNumber{value: number.value, numStartX: number.xStart, numEndX: number.xEnd, numY: number.y, symX: symbol.x, symY: symbol.y})
				break
			}
		}
	}

	// debugging
	// for _, n := range partNumbers {
	// 	fmt.Println(n)
	// }

	var partNumberValues []int
	for _, partNumber := range partNumbers {
		partNumberValues = append(partNumberValues, partNumber.value)
	}
	solution1 := utils.Sum(partNumberValues)

	fmt.Println("Solution 1", solution1)
}

func parseCharMatrix(data [][]rune) ([]Number, []Symbol) {
	var numbers []Number
	var symbols []Symbol
	var currentNumber *Number
	for y, chars := range data {
		for x, char := range chars {
			switch char {
			case '.':
				if currentNumber != nil {
					currentNumber.xEnd = x - 1
					numbers = append(numbers, *currentNumber)
					currentNumber = nil
				}
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if currentNumber != nil {
					currentNumber.value = currentNumber.value*10 + int(char - '0')
				} else {
					currentNumber = new(Number)
					*currentNumber = Number{value: int(char - '0'), xStart: x, y: y}
				}
			default:
				if currentNumber != nil {
					currentNumber.xEnd = x - 1
					numbers = append(numbers, *currentNumber)
					currentNumber = nil
				}
				symbols = append(symbols, Symbol{value: char, x: x, y: y})
			}
		}
	}
	if currentNumber != nil {
		currentNumber.xEnd = len(data[0]) - 1
		numbers = append(numbers, *currentNumber)
		currentNumber = nil
	}
	return numbers, symbols
}

func linesToCharMatrix(lines []string) [][]rune {
	var result [][]rune
	for _, line := range lines {
		result = append(result, lineToCharRow(line))
	}
	return result
}

func lineToCharRow(str string) []rune {
	var result []rune
	for _, ch := range str {
		result = append(result, ch)
	}
	return result
}
