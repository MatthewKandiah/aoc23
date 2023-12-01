package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	dataPath := "data/day1"
	data := readData(dataPath)

	var digitsList [][]int
	for _, str := range data {
		digitsList = append(digitsList, getDigits(str, false))
	}

	var digitsList2 [][]int
	for _, str := range data {
		digitsList2 = append(digitsList2, getDigits(str, true))
	}

	var calibrationValues []int
	for _, digits := range digitsList {
		calibrationValues = append(calibrationValues, getCalibrationValue(digits))
	}

	var calibrationValues2 []int
	for _, digits := range digitsList2 {
		calibrationValues2 = append(calibrationValues2, getCalibrationValue(digits))
	}

	firstAnswer := sum(calibrationValues)
	secondAnswer := sum(calibrationValues2)

	fmt.Println("1.", firstAnswer)
	fmt.Println("2.", secondAnswer)
}

func readData(path string) []string {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println("failed to open {}", path)
		os.Exit(1)
	}

	var result []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}
	readFile.Close()
	return result
}

func getDigits(input string, includeStrings bool) []int {
	var result []int
	for idx, char := range input {
		if unicode.IsDigit(char) {
			result = append(result, int(char - '0'))
		} else if includeStrings {
			digit, found := getDigitFromString(input[idx:])
			if found {
				result = append(result, digit)
			}
		}
	}
	return result
}

func getDigitFromString(input string) (int, bool) {
	var buffer string
	for _, char := range input {
		buffer += string(char)
		if buffer == "one" {
			return 1, true
		} else if buffer == "two" {
			return 2, true
		} else if buffer == "three" {
			return 3, true
		} else if buffer == "four" {
			return 4, true
		} else if buffer == "five" {
			return 5, true
		} else if buffer == "six" {
			return 6, true
		} else if buffer == "seven" {
			return 7, true
		} else if buffer == "eight" {
			return 8, true
		} else if buffer == "nine" {
			return 9, true
		} else if buffer == "zero" {
			return 0, true
		} 
	}
	return -1, false
}

func getCalibrationValue(input []int) int {
	if len(input) == 0 {
		// ignores empty lines
		return 0
	}
	firstDigit := input[0]
	secondDigit := input[len(input)-1]
	return firstDigit*10 + secondDigit
}

func sum(input []int) int {
	result := 0
	for _, value := range input {
		result += value
	}
	return result
}
