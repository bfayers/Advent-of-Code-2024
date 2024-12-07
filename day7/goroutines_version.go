package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/mowshon/iterium"
)

func create_int_slice(str_slice []string) []int {
	// Create a slice of integers from a slice of strings
	int_slice := make([]int, len(str_slice))
	for i, v := range str_slice {
		int_slice[i], _ = strconv.Atoi(v)
	}
	return int_slice
}

func part1(key int, calibration [][]int, operators []string, wg *sync.WaitGroup, result_channel chan<- int) {
	defer wg.Done()
	// var total_valid_results int
	// Find the two entries that sum to 2020
	var total_valid_results int
	for _, value := range calibration {
		// Get the total possible iterations of the two operators in the equation
		total_possible_positions := (len(value) - 1)
		permutations := iterium.Product(operators, total_possible_positions)
		permutations_slice, _ := permutations.Slice()
	every_permutation_loop:
		for _, permutation := range permutations_slice {
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
	// return total_valid_results
	result_channel <- total_valid_results
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

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	part1_result_channel := make(chan int)
	part2_result_channel := make(chan int)
	for key, value := range calibrations {
		wg1.Add(1)
		wg2.Add(1)
		go part1(key, value, []string{"*", "+"}, &wg1, part1_result_channel)
		go part1(key, value, []string{"*", "+", "||"}, &wg2, part2_result_channel)
	}

	go func() {
		wg1.Wait()
		close(part1_result_channel)
		wg2.Wait()
		close(part2_result_channel)
	}()

	var part1_output int
	for result := range part1_result_channel {
		part1_output += result
	}
	var part2_output int
	for result := range part2_result_channel {
		part2_output += result
	}

	fmt.Println("Part 1 output: ", part1_output)
	fmt.Println("Part 2 output: ", part2_output)
}

// This is a goroutines based solve of the same problem
// It is ~3x faster than the original solve
// Code to work out how to use gorountines was taken from: https://stackoverflow.com/a/61825712
