package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func moveGuard(lines [][]string, guard_y int, guard_x int) int {
	previous_x := guard_x
	previous_y := guard_y
	for {
		previous_x = guard_x
		previous_y = guard_y
		// Calculate next position for guard
		switch lines[guard_y][guard_x] {
		case "^":
			guard_y--
		case "V":
			guard_y++
		case "<":
			guard_x--
		case ">":
			guard_x++
		}
		// Check if the new position is beyond the bounds of the array
		if guard_x < 0 || guard_x >= len(lines[0]) || guard_y < 0 || guard_y >= len(lines) {
			break
		}
		// Check if the new position is a wall
		if lines[guard_y][guard_x] == "#" {
			// Rotate the guard 90 degrees to the right and go back to the previous position
			switch lines[previous_y][previous_x] {
			case "^":
				lines[previous_y][previous_x] = ">"
			case "V":
				lines[previous_y][previous_x] = "<"
			case "<":
				lines[previous_y][previous_x] = "^"
			case ">":
				lines[previous_y][previous_x] = "V"
			}
			guard_x = previous_x
			guard_y = previous_y
		} else {
			// Mark the previous position as visited
			lines[guard_y][guard_x] = lines[previous_y][previous_x]
			lines[previous_y][previous_x] = "X"
			// Keep moving
		}
	}
	// Count the number of visited positions
	visited := 1
	for _, line := range lines {
		for _, letter := range line {
			if letter == "X" {
				visited++
			}
		}
	}
	return visited
}

func part1(lines [][]string) int {
	// Find starting position of the guard
	var guard_x int
	var guard_y int
outer:
	for i, line := range lines {
		for j, letter := range line {
			if strings.ContainsAny(letter, "^<>V") {
				guard_y = i
				guard_x = j
				break outer
			}
		}
	}
	fmt.Println("Guard at: ", guard_y, guard_x)
	return moveGuard(lines, guard_y, guard_x)
}

func main() {
	// Load data
	dat, _ := os.Open("input.txt")

	// Define word search array
	var lines [][]string
	// Loop through input
	// Create a scanner
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		// Get line
		line := scanner.Text()
		// Split line into letters
		letter := strings.Split(line, "")
		// Add to letters array
		lines = append(lines, letter)
	}
	fmt.Println("Part 1 output: ", part1(lines))
}
