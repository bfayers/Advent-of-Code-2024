package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func xmas_hunt(lines [][]string, index1 int, index2 int) int {
	var hits int
	// Check if it's possible to go far enough right
	if len(lines[index1])-1 >= index2+3 {
		// Build the slice for going right
		var right string
		right += lines[index1][index2+1] + lines[index1][index2+2] + lines[index1][index2+3]
		if right == "MAS" {
			hits++
		}
	}
	// Check if it's possible to go far enough left
	if index2-3 >= 0 {
		// Build the slice for going left
		var left string
		left += lines[index1][index2-1] + lines[index1][index2-2] + lines[index1][index2-3]
		if left == "MAS" {
			hits++
		}
	}
	// Check if it's possible to go far enough down
	if len(lines)-1 >= index1+3 {
		// Build the slice for going down
		var down string
		down += lines[index1+1][index2] + lines[index1+2][index2] + lines[index1+3][index2]
		if down == "MAS" {
			hits++
		}
	}
	// Check if it's possible to go far enough up
	if index1-3 >= 0 {
		// Build the slice for going up
		var up string
		up = lines[index1-1][index2] + lines[index1-2][index2] + lines[index1-3][index2]
		if up == "MAS" {
			hits++
		}
	}

	// Diagonal versions
	// Check if it's possible to go far enough up and left
	if index1-3 >= 0 && index2-3 >= 0 {
		// Build the slice for going up and left
		var up_left string
		up_left += lines[index1-1][index2-1] + lines[index1-2][index2-2] + lines[index1-3][index2-3]
		if up_left == "MAS" {
			hits++
		}
	}
	// Check if it's possible to go far enough down and left
	if len(lines)-1 >= index1+3 && index2-3 >= 0 {
		// Build the slice for going down and left
		var down_left string
		down_left += lines[index1+1][index2-1] + lines[index1+2][index2-2] + lines[index1+3][index2-3]
		if down_left == "MAS" {
			hits++
		}
	}
	// Check if it's possible to go far enough up and right
	if index1-3 >= 0 && len(lines[index1])-1 >= index2+3 {
		// Build the slice for going up and right
		var up_right string
		up_right += lines[index1-1][index2+1] + lines[index1-2][index2+2] + lines[index1-3][index2+3]
		if up_right == "MAS" {
			hits++
		}
	}
	// Check if it's possible to go far enough down and right
	if len(lines)-1 >= index1+3 && len(lines[index1])-1 >= index2+3 {
		// Build the slice for going down and right
		var down_right string
		down_right += lines[index1+1][index2+1] + lines[index1+2][index2+2] + lines[index1+3][index2+3]
		if down_right == "MAS" {
			hits++
		}
	}

	return hits
}

func part1(lines [][]string) int {
	var count int = 0
	// Loop through the word search entirely
	for index1, line := range lines {
		for index2, char := range line {
			// If this is an "X" run the xmas_hunt function which should find if this has "XMAS"
			if char == "X" {
				count += xmas_hunt(lines, index1, index2)
			}
		}
	}
	return count
}

func x_mas_hunt(lines [][]string, index1 int, index2 int) int {
	// Check if it's possible to go far enough for the cross around this character
	// Checks if it can go far enough up
	if !(index1-1 >= 0) {
		return 0
	}
	// Checks if it can go far enough down
	if !(len(lines)-1 >= index1+1) {
		return 0
	}
	// Checks if it can go far enough left
	if !(index2-1 >= 0) {
		return 0
	}
	// Checks if it can go far enough right
	if !(len(lines[index1])-1 >= index2+1) {
		return 0
	}

	// If we've reached here we can go far enough in every direction to check if this is an "X-MAS"
	// Up left + down right
	left_diagonal := lines[index1-1][index2-1] + lines[index1+1][index2+1]
	// Up right + down left
	right_diagonal := lines[index1-1][index2+1] + lines[index1+1][index2-1]

	if (left_diagonal == "MS" || left_diagonal == "SM") && (right_diagonal == "MS" || right_diagonal == "SM") {
		return 1
	}
	return 0
}

func part2(lines [][]string) int {
	var count int = 0
	for index1, line := range lines {
		for index2, char := range line {
			// If this is an "A" run the x_mas_hunt function
			if char == "A" {
				count += x_mas_hunt(lines, index1, index2)
			}
		}
	}
	return count
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
	fmt.Printf("Part 1 Output: %d\n", part1(lines))
	fmt.Printf("Part 2 Output: %d\n", part2(lines))
}
