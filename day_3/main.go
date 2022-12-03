package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
- Day 3: Compartment sorting

- URL: https://adventofcode.com/2022/day/3

- Overview:
	- This one was tough, its saturday, ill write this later
*/

const ()

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Day 2 Error:", err)
	}
}

func getBadge(lines []string) int {
	// Sort the lines from largest to smallest
	sort.Slice(lines, func(i, j int) bool {
		return len(lines[i]) > len(lines[j])
	})
	chars := map[rune]int{}
	// Go through the longest line and check if the char
	// appears in the next two
	for _, char := range lines[0] {
		if strings.Contains(lines[1], string(char)) && strings.Contains(lines[2], string(char)) {
			chars[char] += 1
		}
	}
	// Get the priority for each char
	for char := range chars {
		if char >= 97 && char <= 122 {
			return int(char) - 96
		}
		if char >= 65 && char <= 90 {
			return int(char) - 38
		}
		fmt.Println(char)
	}
	return -1
}

func run() error {
	file, err := os.Open("./input.txt")
	if err != nil {
		return fmt.Errorf("couldn't open file due to error: %q", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// var prioritySum int
	// Create a scanner to go through line by line
	var lineCount int
	var lines []string
	var prioritySum int
	for scanner.Scan() {
		// Go through line by line, and group lines in groups of three
		if lineCount%3 == 0 && lineCount != 0 {
			prioritySum += getBadge(lines)
			lines = []string{}
		}
		lines = append(lines, scanner.Text())
		// Increment the line count
		lineCount++
	}
	fmt.Println(prioritySum)
	return nil
}
