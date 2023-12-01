package fileutil

import (
	"os"
	"fmt"
	"bufio"
)

func ReadData(path string) []string {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println("failed to open {}", path)
		os.Exit(1)
	}
	defer readFile.Close()

	var result []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}
	return result
}

func Sum(input []int) int {
	result := 0
	for _, value := range input {
		result += value
	}
	return result
}
