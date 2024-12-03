package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) int {
	// Compile a regex to match the "mul" instruction
	instruction_re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	// Compile a regex to match the numbers in the "mul" instruction
	numbers_re := regexp.MustCompile(`\d+`)
	// Find all the instructions
	hits := instruction_re.FindAllString(input, -1)
	// Value to hold the result
	var result int
	for _, hit := range hits {
		numbers := numbers_re.FindAllString(hit, -1)
		// Convert to integers
		var number1 int
		var number2 int
		number1, _ = strconv.Atoi(numbers[0])
		number2, _ = strconv.Atoi(numbers[1])
		// Run the multiply function on the numbers
		result += (number1 * number2)
	}
	return result
}

func part2(input string) int {
	// Compile a regex to match the mul instruction, or do, or dont
	instruction_re := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	// Compile a regex to match the numbers in the "mul" instruction
	numbers_re := regexp.MustCompile(`\d+`)
	// Find all the instructions
	hits := instruction_re.FindAllString(input, -1)
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
		// We are enabled, so we can run the instruction
		numbers := numbers_re.FindAllString(hit, -1)
		// Convert to integers
		var number1 int
		var number2 int
		number1, _ = strconv.Atoi(numbers[0])
		number2, _ = strconv.Atoi(numbers[1])
		// Run the multiply function on the numbers
		result += (number1 * number2)

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
	// fmt.Println(sb.String())

	fmt.Printf("Part 1 Output: %d\n", part1(sb.String()))
	fmt.Printf("Part 2 Output: %d\n", part2(sb.String()))

}
