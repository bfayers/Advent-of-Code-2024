package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Define all regex used in the program
var part1_instruction_re = regexp.MustCompile(`mul\(\d+,\d+\)`)
var part2_instruction_re = regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
var numbers_re = regexp.MustCompile(`\d+`)

func get_instructions(input string, part2 bool) []string {
	if part2 {
		return part2_instruction_re.FindAllString(input, -1)
	}
	return part1_instruction_re.FindAllString(input, -1)
}

// Run instruction
func run_instruction(instruction string) int {
	numbers := numbers_re.FindAllString(instruction, -1)
	// Convert to integers
	var number1 int
	var number2 int
	number1, _ = strconv.Atoi(numbers[0])
	number2, _ = strconv.Atoi(numbers[1])
	// Run the multiply function on the numbers
	return number1 * number2
}

func part1(input string) int {
	// Find all the instructions
	hits := get_instructions(input, false)
	// Value to hold the result
	var result int
	for _, hit := range hits {
		// Run the instruction
		result += run_instruction(hit)
	}
	return result
}

func part2(input string) int {
	// Find all the instructions
	hits := get_instructions(input, true)
	// Value to hold the result
	var result int
	// Value to hold if we're enabled or not
	var enabled bool = true

	// Loop through the instructions
	for _, hit := range hits {

		// Enable processing if we hit a "do()" instruction, disable if we hit a "don't()" instruction
		switch hit {
		case "do()":
			enabled = true
			continue
		case "don't()":
			enabled = false
			continue
		}

		if !enabled {
			// We're not enabled right now, skip this instruction
			continue
		}

		// Run the instruction
		result += run_instruction(hit)
	}

	return result

}

func main() {
	dat, _ := os.Open("input.txt")
	// Make a stringsbuilder to concat all the lines
	var sb strings.Builder
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}

	// Define the known correct outputs
	var known_outputs []int = []int{178886550, 87163705}
	// Run part1 and part2
	var part1_output int = part1(sb.String())
	var part2_output int = part2(sb.String())

	// Print the results and check if they are correct
	fmt.Printf("Part 1 Output: %d - ", part1_output)
	if part1_output == known_outputs[0] {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("Part 2 Output: %d - ", part2_output)
	if part2_output == known_outputs[1] {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}

}
