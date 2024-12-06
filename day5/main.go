package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type number_rules struct {
	before []int
	after  []int
}

func part1(rule_map map[int]number_rules, updates [][]int) ([][]int, [][]int) {
	var valid_updates [][]int
	var invalid_updates [][]int
	for _, update := range updates {
		// Check if the update is valid
		var valid bool = true
		for index, ele := range update {
			// Check if the element is in the map
			entry, ok := rule_map[ele]
			if ok {
				// Check if every element after this element is in the after slice
				if !(index == len(update)-1) {
					// Don't run after check for the last element
					items_after := update[index+1:]
					for _, item := range items_after {
						if !(slices.Contains(entry.after, item)) {
							valid = false
						}
					}
				}
				// Check if every element before this element is in the before slice
				if !(index == 0) {
					// Don't run before check for the first element
					items_before := update[:index]
					for _, item := range items_before {
						if !(slices.Contains(entry.before, item)) {
							valid = false
						}
					}
				}
			}
		}
		if valid {
			// fmt.Println("Valid update:", update)
			valid_updates = append(valid_updates, update)
		} else {
			// fmt.Println("Invalid update:", update)
			invalid_updates = append(invalid_updates, update)
		}
	}
	return valid_updates, invalid_updates
}

func process_rules(rules [][]int) map[int]number_rules {
	rule_map := make(map[int]number_rules)
	for _, rule := range rules {
		// Check if the rule for afters already exists in the map
		entry, ok := rule_map[rule[0]]
		if ok {
			// Rule already exists, append to the after & before slices
			entry.after = append(entry.after, rule[1])
			rule_map[rule[0]] = entry

		} else {
			// Rule does not exist, create a new entry
			var new_entry number_rules
			new_entry.after = append(new_entry.after, rule[1])
			rule_map[rule[0]] = new_entry
		}
		// Check if the rule for befores already exists in the map
		entry, ok = rule_map[rule[1]]
		if ok {
			// Rule already exists, append to the before slice
			entry.before = append(entry.before, rule[0])
			rule_map[rule[1]] = entry
		} else {
			// Rule does not exist, create a new entry
			var new_entry number_rules
			new_entry.before = append(new_entry.before, rule[0])
			rule_map[rule[1]] = new_entry
		}
	}
	return rule_map
}

func part2(rule_map map[int]number_rules, updates [][]int) int {
	var middle_totals int
	// We can build a 'fixed' update by making a slice
	// And then based on the rules we can do some logic with concat to put in right order
	// Hopefully....
	for _, update := range updates {
		// Fix this update
		var fixed_update []int
	outer:
		for _, ele := range update {
			if len(fixed_update) == 0 {
				fixed_update = append(fixed_update, ele)
				continue
			}
			for _, f_ele := range fixed_update {
				if slices.Contains(rule_map[f_ele].before, ele) {
					// This should be before the current item
					fixed_update = append([]int{ele}, fixed_update...)
					continue outer
				} else if slices.Contains(rule_map[f_ele].after, ele) {
					// If this should be after the current item, it could need to be after a future one too
					// Loop backwards over fixed_update to find the latest AFTER point
					for i := len(fixed_update) - 1; i >= 0; i-- {
						if slices.Contains(rule_map[fixed_update[i]].after, ele) {
							// This should be after THIS item
							fixed_update = append(fixed_update[:i+1], append([]int{ele}, fixed_update[i+1:]...)...)
							continue outer
						}
					}
					continue outer
				}
			}
		}
		// Update is now fixed, find the middle page number and add up again
		// Get middle element
		middle_totals += sum_middle_elements([][]int{fixed_update})
	}
	return middle_totals
}

func sum_middle_elements(updates [][]int) int {
	var total int
	for _, update := range updates {
		// Get middle element
		mid := int(math.Ceil(float64(len(update))/2)) - 1
		total += update[mid]
	}
	return total
}

func main() {
	// Load data
	dat, _ := os.Open("input.txt")

	//Define array of rules and updates
	var rules [][]int
	var updates [][]int

	scanner := bufio.NewScanner(dat)
	var processing_rules bool = true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			processing_rules = false
			continue
		}
		if processing_rules {
			// Split the rule on |
			split_line := strings.Split(line, "|")
			var rule []int
			for _, num := range split_line {
				this_num, _ := strconv.Atoi(num)
				rule = append(rule, this_num)
			}
			rules = append(rules, rule)
		} else {
			// Split the update on ,
			split_line := strings.Split(line, ",")
			var update []int
			for _, num := range split_line {
				this_num, _ := strconv.Atoi(num)
				update = append(update, this_num)
			}
			updates = append(updates, update)
		}
	}

	// Process the rules
	rule_map := process_rules(rules)

	p1_valid_updates, invalid_updates := part1(rule_map, updates)
	fmt.Println("Part 1 Output:", sum_middle_elements(p1_valid_updates))

	// Part 2
	fmt.Println("Part 2 Output:", part2(rule_map, invalid_updates))
}
