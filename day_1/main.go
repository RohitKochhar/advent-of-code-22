package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
- Day 1: Calorie Counting

- URL: https://adventofcode.com/2022/day/1

- Overview:
	- Given a list of values (calories) seperated into
	  sublists by blank lines, find the sublist with
	  the highest total
*/

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Day 1 Error:", err)
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
	var (
		globalFirst  int // Keep track of global maximum
		globalSecond int // Keep track of the second max
		globalThird  int // Keep track of the third max
		localMax     int // Keep track of local maximum
	)
	for scanner.Scan() {
		if scanner.Text() == "" {
			// End of a collection, compare
			if localMax > globalFirst {
				globalThird = globalSecond
				globalSecond = globalFirst
				globalFirst = localMax
			} else if localMax > globalSecond {
				globalThird = globalSecond
				globalSecond = localMax
			} else if localMax > globalThird {
				globalThird = localMax
			}
			// Reset local max
			localMax = 0
		} else {
			amt, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return fmt.Errorf("error while converting input to int: %q", err)
			}
			localMax += amt
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error while scanning: %q", err)
	}
	fmt.Println(globalFirst, globalSecond, globalThird, globalFirst+globalSecond+globalThird)
	return nil
}
