package util

import (
	"bufio"
	"os"
)

func ReadData(path string) ([]string, error) {
	readFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	var result []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}
	return result, nil
}

func Sum(input []int) int {
	result := 0
	for _, value := range input {
		result += value
	}
	return result
}
