package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(left []float64, right []float64) float64 {
	total_distance := float64(0)
	// Iterate through the arrays simultaneously (use the index)
	for index, element := range left {
		// Get the distance
		distance := math.Abs(element - right[index])
		total_distance = total_distance + distance
	}

	return total_distance
}

func count(array []float64, search float64) float64 {
	total := float64(0)
	for _, element := range array {
		if element == search {
			total++
		}
	}
	return total
}

func part2(left []float64, right []float64) float64 {
	total_similarity := float64(0)
	// Iterate over left array
	for _, element := range left {
		// Count occurences of element in right
		right_count := count(right, element)
		this_similarity := element * right_count
		total_similarity += this_similarity
	}

	return total_similarity
}

func main() {
	dat, _ := os.Open("input.txt")
	// Define left and right arrays (slices?)
	var left []float64
	var right []float64
	// Loop through the input
	// Create bufio scanner
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		// Split into the left and right
		lr := strings.Split(scanner.Text(), "   ")
		li, _ := strconv.ParseFloat(lr[0], 64)
		ri, _ := strconv.ParseFloat(lr[1], 64)
		// Append to the left and right arrays
		left = append(left, li)
		right = append(right, ri)
	}

	// Sort left and right arrays (slices?)
	sort.Float64s(left)
	sort.Float64s(right)

	// Run Part 1
	total_distance := part1(left, right)
	fmt.Printf("Total distance: %f\n", total_distance)

	// Run Part 2
	total_similarity := part2(left, right)
	fmt.Printf("Total similarity: %f\n", total_similarity)

}
