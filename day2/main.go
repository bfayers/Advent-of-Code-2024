package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func p1_check_report(report []int) bool {
	// Go through the report, up to index - 1 for the sake of checking
	var increasing bool
	var decreasing bool
	for index := range len(report) - 1 {
		// Get the difference between index and index+1
		difference := math.Abs(float64(report[index]) - float64(report[index+1]))
		if report[index] < report[index+1] {
			if decreasing {
				// This report has flipped from decreasing to increasing, it is unsafe
				return false
			}
			increasing = true
		} else if report[index] > report[index+1] {
			if increasing {
				// This report has flipped from increasing to decreasing, it is unsafe
				return false
			}
			decreasing = true
		} else if report[index] == report[index+1] {
			// This report is the same as the previous, it is unsafe
			return false
		}
		if !(difference >= 1 && difference <= 3) {
			// This report is unsafe
			return false
		}
	}
	return true
}

func part1(reports [][]int) int {
	// Check every report
	var safe_reports int
	for _, report := range reports {
		report_status := p1_check_report(report)
		if report_status {
			safe_reports++
		}
	}
	return safe_reports
}

// Taken from https://stackoverflow.com/a/57213476
func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func p2_problem_dampener(report []int) bool {
	// Go through every value of the report removing it and checking if it is now safe as a result
	for index := range len(report) {
		new_report := RemoveIndex(report, index)
		// Check this 'new report'
		if p1_check_report(new_report) {
			// This report is now safe
			return true
		} else {
			continue
		}
	}
	return false
}

func part2(reports [][]int) int {
	//Check every report, but implement the "dampener"
	var safe_reports int
	for _, report := range reports {
		report_status := p1_check_report(report)
		if !report_status {
			if p2_problem_dampener(report) {
				safe_reports++
			}
		} else {
			safe_reports++
		}
	}
	return safe_reports
}

func main() {
	dat, _ := os.Open("input.txt")
	// Define reports array
	var reports [][]int
	// Loop through the input
	// Create bufio scanner
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		// Define this report array
		var this_report []int
		// Split into the left and right
		report_strings := strings.Split(scanner.Text(), " ")
		for _, v := range report_strings {
			// Convert to int
			i, _ := strconv.Atoi(v)
			// Append to the left and right arrays
			this_report = append(this_report, i)
		}
		// Append to the left and right arrays
		reports = append(reports, this_report)
	}

	fmt.Printf("Part 1 Safe Reports: %d\n", part1(reports))
	fmt.Printf("Part 2 Safe Reports: %d\n", part2(reports))
}
