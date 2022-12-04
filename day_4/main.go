package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
- Day 4: Subinterval checking

- URL: https://adventofcode.com/2022/day/4

- Overview:
	- Given a list of inputs of the form 1-3, 1-6
	  how many inputs have an interval that contains the other
*/

const ()

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Day 4 Error:", err)
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
	// Count to count number of matches
	var count int
	for scanner.Scan() {
		// Split the line at the comma to get the two assignments
		assignments := strings.Split(scanner.Text(), ",")
		// Split the assignments at the dash to get the bounds
		a1 := strings.Split(assignments[0], "-")
		a2 := strings.Split(assignments[1], "-")
		// Convert string arrays to int
		int1, int2, err := convertStringArrayToIntArray(a1, a2)
		if err != nil {
			return err
		}
		// Check if a1 overlaps a2
		if isContained := checkOverlap(int1, int2); isContained {
			fmt.Printf("found super: %d %d\n", int1, int2)
			count++
		}
	}
	fmt.Println("Total Number of contained intervals: ", count)
	return nil
}

// part1
func checkSubInterval(i1, i2 []int) bool {
	// Check that the i2 max and min is inside i1
	if i1[0] <= i2[0] && i1[1] >= i2[1] {
		return true
	}
	if i2[0] <= i1[0] && i2[1] >= i1[1] {
		return true
	}
	return false
}

// part2
func checkOverlap(i1, i2 []int) bool {
	// Check if the lower bound of the first interval is
	// contained within the bounds of the second
	if i2[0] <= i1[0] && i1[0] <= i2[1] {
		return true
	}
	// Check if the upper bound of the first interval is
	// contained within the bounds of the second
	if i2[0] <= i1[1] && i1[1] <= i2[1] {
		return true
	}
	// Check if the lower bound of the second interval is
	// contained within the bounds of the first
	if i1[0] <= i2[0] && i2[0] <= i1[1] {
		return true
	}
	// Check if the upper bound of the second interval is
	// contained within the bounds of the first
	if i1[0] <= i2[1] && i2[1] <= i1[1] {
		return true
	}
	return false
}

func convertStringArrayToIntArray(i1, i2 []string) ([]int, []int, error) {
	// Convert the strings to int
	i1Lower, err := strconv.Atoi(i1[0])
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't convert input to int: %q", err)
	}
	i1Upper, err := strconv.Atoi(i1[1])
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't convert input to int: %q", err)
	}
	i2Lower, err := strconv.Atoi(i2[0])
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't convert input to int: %q", err)
	}
	i2Upper, err := strconv.Atoi(i2[1])
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't convert input to int: %q", err)
	}
	return []int{i1Lower, i1Upper}, []int{i2Lower, i2Upper}, nil
}
