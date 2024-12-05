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

func part1(rule_map map[int]number_rules, updates [][]int) [][]int {
	var valid_updates [][]int
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
		}
	}
	return valid_updates
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
	// fmt.Println(rule_map)
	return rule_map
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

	// fmt.Println(rules)
	// fmt.Println(updates)

	// Process the rules
	rule_map := process_rules(rules)

	p1_valid_updates := part1(rule_map, updates)
	var middle_totals int
	for _, update := range p1_valid_updates {
		// Get middle element
		mid := int(math.Ceil(float64(len(update))/2)) - 1
		middle_totals += update[mid]
	}
	fmt.Println("Part 1 Output:", middle_totals)

}
