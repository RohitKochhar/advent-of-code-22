package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
- Day 5: Crate sorting

- URL: https://adventofcode.com/2022/day/5

- Overview:
	- Lots of slice sorting
*/

const ()

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Day 5 Error:", err)
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
	gridFinished := false
	var crateGrid *Grid
	var grid []string
	for scanner.Scan() {
		if !gridFinished {
			if scanner.Text() == "" {
				gridFinished = true
				crateGrid = createGrid(grid)
			} else {
				grid = append(grid, scanner.Text())
			}
		} else {
			crateGrid.parseInstruction(scanner.Text())
		}
	}
	result := ""
	for i := 0; i < len(crateGrid.Stacks); i++ {
		result += crateGrid.Stacks[i][0]
	}
	fmt.Println(result)
	return nil
}

type Grid struct {
	Stacks map[int][]string
}

func createGrid(g []string) *Grid {
	// New grid object to track crate position
	grid := Grid{}
	grid.Stacks = make(map[int][]string)
	// The last row contains the number of grids
	numCols := len(strings.ReplaceAll(g[len(g)-1], " ", ""))
	// Go through the beginning lines
	for _, l := range g[:len(g)-1] {
		crates := strings.Split(l, " ")
		fmt.Println(crates)
		for i := 0; i < numCols; i++ {
			if !(crates[i] == "[/]") {
				grid.Stacks[i] = append(grid.Stacks[i], crates[i])
			}
		}
	}
	return &grid
}

func (g *Grid) parseInstruction(inst string) error {
	fmt.Println(inst)
	// Split the instruction at spaces
	insts := strings.Split(inst, " ")

	numCrates, err := strconv.Atoi(insts[1])

	if err != nil {
		return err
	}
	source, err := strconv.Atoi(insts[3])
	if err != nil {
		return err
	}
	dest, err := strconv.Atoi(insts[5])
	if err != nil {
		return err
	}
	// Account for zero indexing
	source += -1
	dest += -1
	fmt.Println(g.Stacks)
	// Move crates from dest to source
	//for i := 0; i < numCrates; i++ {
	// Part 1: Move like a stack, LIFO
	// Get the crate from the source
	// crate := g.Stacks[source][0]
	// // Delete it from the source
	// g.Stacks[source] = g.Stacks[source][1:]
	// // Add it to the front of the destination
	// g.Stacks[dest] = append([]string{crate}, g.Stacks[dest]...)\
	//}
	// Part 2: Order is maintained
	// Get set of crates from source
	crates := make([]string, numCrates)
	copy(crates, g.Stacks[source][:numCrates])
	fmt.Println("source crates (pre-change):", g.Stacks[source])
	// Delete crates from source
	g.Stacks[source] = g.Stacks[source][numCrates:]
	fmt.Println("source crates (post-change):", g.Stacks[source])
	// Add crates to the front of destination
	g.Stacks[dest] = append(crates[:], g.Stacks[dest][:]...)
	fmt.Println("source crates (p):", g.Stacks[source])
	fmt.Println(crates)
	fmt.Println(g.Stacks)
	return nil
}
