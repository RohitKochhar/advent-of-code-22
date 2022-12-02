package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
- Day 2: Rock Paper Scissors

- URL: https://adventofcode.com/2022/day/2

- Overview:
	- Given a list of inputs of the form A X
	  determine who will more points in a
	  rock paper scissors tournament where
	  points are awarded based on outcome and choice
	  by two different critera
*/

const (
	// Scores based on outcome
	WIN  = 6
	LOSE = 0
	DRAW = 3
	// Scores based on selection
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Day 2 Error:", err)
	}
}

func run() error {
	file, err := os.Open("./input.txt")
	if err != nil {
		return fmt.Errorf("couldn't open file due to error: %q", err)
	}
	defer file.Close()
	// Create a scanner to go through line by line
	scanner := bufio.NewScanner(file)
	var totalScore int
	for scanner.Scan() {
		// Split the line at the space
		choices := strings.Split(scanner.Text(), " ")
		res := play_2(choices[0], choices[1])
		if res == -1 {
			return fmt.Errorf("couldn't get result")
		}
		totalScore += res
	}
	fmt.Println(totalScore)
	return nil
}

// play_1 is used for part 1 where:
// A,X = rock, B,Y = paper, C,Z = scissors
func play_1(opponent, player string) int {
	// Switch based on opponent's choice
	switch opponent {
	// Opponent choses rock
	case "A":
		switch player {
		case "X":
			// Player draws with rock
			return DRAW + ROCK
		case "Y":
			// Player won with paper
			return WIN + PAPER
		case "Z":
			// Player lost with scissor
			return LOSE + SCISSORS
		}
	// Opponent choses Paper
	case "B":
		switch player {
		case "X":
			// Player loses with rock
			return LOSE + ROCK
		case "Y":
			// Player draws with paper
			return DRAW + PAPER
		case "Z":
			// Player wins with scissor
			return WIN + SCISSORS
		}
	// Opponent choses scissors
	case "C":
		switch player {
		case "X":
			// Player wins with rock
			return WIN + ROCK
		case "Y":
			// Player losts with paper
			return LOSE + PAPER
		case "Z":
			// Player draws with scissor
			return DRAW + SCISSORS
		}
	}
	return -1
}

// play_2 is used for part two, where
// for opponent, A,B,C = rock,paper,scissor
// for player, X,Y,Z = lose,draw,win
func play_2(opponent, player string) int {
	// Switch based on opponent's choice
	switch opponent {
	// Opponent choses rock
	case "A":
		switch player {
		case "X":
			// Player loses with scissors
			return LOSE + SCISSORS
		case "Y":
			// Player draws with rock
			return DRAW + ROCK
		case "Z":
			// Player wins with paper
			return WIN + PAPER
		}
	// Opponent choses Paper
	case "B":
		switch player {
		case "X":
			// Player loses with rock
			return LOSE + ROCK
		case "Y":
			// Player draws with paper
			return DRAW + PAPER
		case "Z":
			// Player wins with scissor
			return WIN + SCISSORS
		}
	// Opponent choses scissors
	case "C":
		switch player {
		case "X":
			// Player loses with paper
			return LOSE + PAPER
		case "Y":
			// Player draws with scissor
			return DRAW + SCISSORS
		case "Z":
			// Player wins with rock
			return WIN + ROCK
		}
	}
	return -1
}
