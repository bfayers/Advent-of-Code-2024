package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/life4/genesis/slices"
)

func create_int_slice(str_slice []string) []int {
	// Create a slice of integers from a slice of strings
	int_slice := make([]int, len(str_slice))
	for i, v := range str_slice {
		int_slice[i], _ = strconv.Atoi(v)
	}
	return int_slice
}

func part1(calibrations map[int][][]int, operators []string) int {
	// var total_valid_results int
	// Find the two entries that sum to 2020
	var total_valid_results int
	for key, master_value := range calibrations {
		for _, value := range master_value {
			// Get the total possible iterations of the two operators in the equation
			total_possible_positions := (len(value) - 1)
			permutations := slices.Product(operators, total_possible_positions)
		every_permutation_loop:
			for permutation := range permutations {
				var perm_slot int
				var perm_total int
				perm_total = value[0]
				// Iterate in a way that we can run the numbers and operators in order
				for i := 1; i < len(value); i++ {
					switch permutation[perm_slot] {
					case "*":
						// Multiple
						perm_total *= value[i]
					case "+":
						// Add
						perm_total += value[i]
					case "||":
						// Concatenate as a string
						perm_total, _ = strconv.Atoi(strconv.Itoa(perm_total) + strconv.Itoa(value[i]))
					}
					perm_slot++
					if perm_total > key {
						continue every_permutation_loop
					}
				}
				if perm_total == key {
					total_valid_results += key
					break every_permutation_loop
				}
			}
		}
	}
	return total_valid_results
}

func main() {
	// Load data
	dat, _ := os.Open("input.txt")

	// Define word search array
	calibrations := make(map[int][][]int)
	// Loop through input
	// Create a scanner
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		// Get line
		line := scanner.Text()
		// Split line into target value and equation
		parts := strings.Split(line, ": ")
		// Split equation into parts
		equation := strings.Split(parts[1], " ")
		// Add to letters array
		target_value, _ := strconv.Atoi(parts[0])
		entry, ok := calibrations[target_value]
		if ok {
			calibrations[target_value] = append(entry, create_int_slice(equation))
		} else {
			calibrations[target_value] = [][]int{create_int_slice(equation)}
		}
	}
	fmt.Println("Part 1 output: ", part1(calibrations, []string{"*", "+"}))
	fmt.Println("Part 2 output: ", part1(calibrations, []string{"*", "+", "||"}))
}
