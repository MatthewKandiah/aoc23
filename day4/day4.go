package main

import (
	"aoc23/util"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	lines, err := util.ReadData("day4")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	winnersList, candidatesList, err := parseLines(lines)

	var matchCounts []int
	for idx := range winnersList {
		numMatches := countMatches(winnersList[idx], candidatesList[idx])
		matchCounts = append(matchCounts, numMatches)
	}

	var scores []int
	for _, count := range matchCounts {
		score := int(math.Pow(2, float64(count-1)))
		scores = append(scores, score)
	}

	fmt.Println("Solution 1:", util.Sum(scores))
}

func countMatches(a []int, b []int) int {
	var result int
	for _, valA := range a {
		for _, valB := range b {
			if valA == valB {
				result++
				break
			}
		}
	}
	return result
}

func parseLines(lines []string) ([][]int, [][]int, error) {
	var winnersList [][]int
	var candidatesList [][]int
	for _, line := range lines {
		var winners []int
		var candidates []int
		var s scanner.Scanner
		s.Init(strings.NewReader(line))
		// skip word Card
		s.Scan()
		// skip card number
		s.Scan()
		// skip colon
		s.Scan()
		// scan winning numbers
		for s.Scan() == scanner.Int {
			value, err := strconv.Atoi(s.TokenText())
			if err != nil {
				return winnersList, candidatesList, err
			}
			winners = append(winners, value)
		}
		// skip |
		// scan candidate numbers
		for s.Scan() == scanner.Int {
			value, err := strconv.Atoi(s.TokenText())
			if err != nil {
				return winnersList, candidatesList, err
			}
			candidates = append(candidates, value)
		}
		// end of line
		winnersList = append(winnersList, winners)
		candidatesList = append(candidatesList, candidates)
	}
	return winnersList, candidatesList, nil
}
