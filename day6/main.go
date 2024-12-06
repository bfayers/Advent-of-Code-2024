package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type coordinate struct {
	x int
	y int
}
type movement struct {
	directions []string
}

func add_to_coordinate_movements(coordinate_movements map[coordinate]movement, x int, y int, direction string) (map[coordinate]movement, bool) {
	entry, ok := coordinate_movements[coordinate{x: x, y: y}]
	if ok {
		if slices.Contains(entry.directions, direction) {
			// We have visited this slot with this direction before, so we are in a loop
			return coordinate_movements, true
		}
		entry.directions = append(entry.directions, direction)
		coordinate_movements[coordinate{x: x, y: y}] = entry
	} else {
		coordinate_movements[coordinate{x: x, y: y}] = movement{directions: []string{direction}}
	}
	return coordinate_movements, false
}

func moveGuard(lines [][]string, guard_y int, guard_x int) (visited int, coordinate_movements map[coordinate]movement, loop bool) {
	previous_x := guard_x
	previous_y := guard_y
	coordinate_movements = make(map[coordinate]movement)
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
			// Record the last position as visited
			coordinate_movements, _ = add_to_coordinate_movements(coordinate_movements, previous_x, previous_y, lines[previous_y][previous_x])
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
			coordinate_movements, loop = add_to_coordinate_movements(coordinate_movements, previous_x, previous_y, lines[previous_y][previous_x])
			if loop {
				return 0, coordinate_movements, loop
			}
			lines[previous_y][previous_x] = "X"
			// Keep moving
		}
	}
	// Count the number of visited positions
	visited = 1
	for _, line := range lines {
		for _, letter := range line {
			if letter == "X" {
				visited++
			}
		}
	}
	return visited, coordinate_movements, loop
}

func find_guard_position(lines [][]string) (int, int) {
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
	return guard_y, guard_x
}

func duplicate_lines(lines [][]string) [][]string {
	new_lines := make([][]string, len(lines))
	for i := range lines {
		new_lines[i] = make([]string, len(lines[i]))
		copy(new_lines[i], lines[i])
	}
	return new_lines
}

func part1(lines [][]string) (int, map[coordinate]movement) {
	// Find starting position of the guard
	guard_y, guard_x := find_guard_position(lines)
	fmt.Println("Guard at: ", guard_y, guard_x)
	// Move the guard
	visited, coordinate_movements, _ := moveGuard(lines, guard_y, guard_x)
	return visited, coordinate_movements
}

func part2(lines [][]string, coordinate_movements map[coordinate]movement) int {
	var loops int
	// Find starting position of the guard
	guard_y, guard_x := find_guard_position(lines)
	fmt.Println("Guard at: ", guard_y, guard_x)
	// Move the guard
	// Inject a new obstacle at one of the visisted positions
	for key, _ := range coordinate_movements {
		if key.x == guard_x && key.y == guard_y {
			// Skip this one because it will never start
			continue
		}
		// use a copy of the lines array
		new_lines := duplicate_lines(lines)
		new_lines[key.y][key.x] = "#"
		_, _, loop := moveGuard(new_lines, guard_y, guard_x)
		if loop {
			loops++
		}
	}
	return loops
}

func main() {
	// Load data
	dat, _ := os.Open("input.txt")

	// Define word search array
	var source_lines [][]string
	// Loop through input
	// Create a scanner
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		// Get line
		line := scanner.Text()
		// Split line into letters
		letter := strings.Split(line, "")
		// Add to letters array
		source_lines = append(source_lines, letter)
	}
	// Part 1
	// use a copy of the lines array
	lines_copy := duplicate_lines(source_lines)
	visited, coordinate_movements := part1(lines_copy)
	fmt.Println("Part 1 output: ", visited)
	// Part 2
	fmt.Println("Part 2 output: ", part2(source_lines, coordinate_movements))
}
