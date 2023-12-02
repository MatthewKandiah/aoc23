package main

import (
	util "aoc23/util"
	"fmt"
	"os"
	"unicode"
)

func main() {
	dataPath := "data"
	data, err := util.ReadData(dataPath)
	if err != nil {
		fmt.Println("Failed to open", data)
		os.Exit(1)
	}

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

	firstAnswer := util.Sum(calibrationValues)
	secondAnswer := util.Sum(calibrationValues2)

	fmt.Println("1.", firstAnswer)
	fmt.Println("2.", secondAnswer)
}

func getDigits(input string, includeStrings bool) []int {
	var result []int
	for idx, char := range input {
		if unicode.IsDigit(char) {
			result = append(result, int(char-'0'))
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
		switch buffer {
		case "one":
			return 1, true
		case "two":
			return 2, true
		case "three":
			return 3, true
		case "four":
			return 4, true
		case "five":
			return 5, true
		case "six":
			return 6, true
		case "seven":
			return 7, true
		case "eight":
			return 8, true
		case "nine":
			return 9, true
		case "zero":
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
