package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
- Day 6: Message decoding

- URL: https://adventofcode.com/2022/day/6

- Overview:
	- Given a string containing a datastream of chars
		- Part 1: Find the first set of 4 unique chars
		- Part 2: Find the first set of 14 unique chars
*/

const ()

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Day 6 Error:", err)
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
	var res1, res2 int
	for scanner.Scan() {
		res1 = Part1(scanner.Text())
		res2 = Part2(scanner.Text())
	}
	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)
	return nil
}

func Part1(in string) int {
	fmt.Println(in)
	// Create a sliding window that looks at
	// 4 chars at a time
	for idx := 0; idx < len(in)-3; idx++ {
		sw := []rune{rune(in[idx]), rune(in[idx+1]), rune(in[idx+2]), rune(in[idx+3])}
		// Check if the string only contains unique runes
		if isMarker(sw) {
			// Offset since we want when the start of packet marker ends
			return idx + 4
		}
	}

	return -1
}

func Part2(in string) int {
	fmt.Println(in)
	// Create a sliding window that looks at
	// 14 chars at a time
	for idx := 0; idx < len(in)-13; idx++ {
		sw := []rune{}
		for i := 0; i < 14; i++ {
			sw = append(sw, rune(in[idx+i]))
		}
		// Check if the string only contains unique runes
		if isMarker(sw) {
			// Offset since we want when the start of packet marker ends
			return idx + 14
		}
	}

	return -1
}

func isMarker(in []rune) bool {
	for i := 0; i < len(in); i++ {
		for j := i; j < len(in); j++ {
			if in[i] == in[j] && i != j {
				return false
			}
		}
	}
	return true
}
