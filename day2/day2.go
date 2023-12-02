package main

import (
	util "aoc23/util"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

func main() {
	data := "data"
	lines, err := util.ReadData(data)
	if err != nil {
		fmt.Println("failed to open", data)
		os.Exit(1)
	}

	var games []Game
	for _, line := range lines {
		game, err := gameFromLine(line)
		if err != nil {
			fmt.Println("failed to parse", line)
			fmt.Println(err)
			os.Exit(1)
		}
		games = append(games, game)
	}

	const maxReds = 12
	const maxGreens = 13
	const maxBlues = 14

	maxRounds := maxRoundsFromGames(games)
	var validGameIds []int
	var invalidGameIds []int
	for _, maxRound := range maxRounds {
		if maxRound.reds > maxReds || maxRound.greens > maxGreens || maxRound.blues > maxBlues {
			invalidGameIds = append(invalidGameIds, maxRound.id)
		} else {
			validGameIds = append(validGameIds, maxRound.id)
		}
	}

	fmt.Println("Answer 1:", util.Sum(validGameIds))

	var roundPowers []int
	for _, maxRound := range maxRounds {
		roundPowers = append(roundPowers, maxRound.reds * maxRound.greens * maxRound.blues)
	}

	fmt.Println("Answer 2:", util.Sum(roundPowers))
	
}

type Colour string

const (
	Red   Colour = "red"
	Green        = "green"
	Blue         = "blue"
	None
)

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	reds   int
	greens int
	blues  int
}

type MaxRounds struct {
	id     int
	reds   int
	greens int
	blues  int
}

/*
return slice of games
on error, return games successfully parsed and non-nil error
*/
func gameFromLine(line string) (Game, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(line))
	var result Game
	result.rounds = append(result.rounds, Round{})

	// skip word "Game"
	s.Scan()

	if s.Scan() != scanner.Int {
		return result, errors.New("Invalid game number token")
	}
	gameId, err := strconv.Atoi(s.TokenText())
	if err != nil {
		return result, err
	}
	result.id = gameId

	// skip colon
	s.Scan()

	for {
		if s.Scan() != scanner.Int {
			return result, errors.New("Invalid dice number token")
		}
		diceNumber, err := strconv.Atoi(s.TokenText())
		if err != nil {
			return result, err
		}

		if s.Scan() != scanner.Ident {
			return result, errors.New("Invalid dice colour token")
		}
		diceColour, err := colourFromString(s.TokenText())
		if err != nil {
			return result, err
		}

		// assume there is only one count for a given colour in a given round
		switch diceColour {
		case Red:
			result.rounds[len(result.rounds)-1].reds = diceNumber
		case Green:
			result.rounds[len(result.rounds)-1].greens = diceNumber
		case Blue:
			result.rounds[len(result.rounds)-1].blues = diceNumber
		}

		separatorToken := s.Scan()
		switch {
		case separatorToken == scanner.EOF:
			return result, nil
		case s.TokenText() == ";":
			result.rounds = append(result.rounds, Round{})
		}
	}
}

func colourFromString(input string) (Colour, error) {
	switch input {
	case string(Red):
		return Red, nil
	case string(Green):
		return Green, nil
	case string(Blue):
		return Blue, nil
	default:
		return None, errors.New("Invalid colour string")
	}
}

func maxRoundsFromGames(games []Game) []MaxRounds {
	var results []MaxRounds
	for _, game := range games {
		results = append(results, MaxRoundFromGame(game))
	}
	return results
}

func MaxRoundFromGame(game Game) MaxRounds {
	var result MaxRounds
	result.id = game.id
	for _, round := range game.rounds {
		if round.reds > result.reds {
			result.reds = round.reds
		}
		if round.greens > result.greens {
			result.greens = round.greens
		}
		if round.blues > result.blues {
			result.blues = round.blues
		}
	}
	return result
}
