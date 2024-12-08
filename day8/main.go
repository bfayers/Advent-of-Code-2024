package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

type coordinate struct {
	x int
	y int
}

var antenna_re = regexp.MustCompile(`([A-Z]|[a-z]|[0-9])`)

func remove_dupes(input []coordinate) []coordinate {
	var result []coordinate
	for _, value := range input {
		if !slices.Contains(result, value) {
			result = append(result, value)
		}
	}
	return result
}

func part1(source_map map[coordinate]string, map_width int, map_height int) int {
	var all_antinodes []coordinate

	for key, value := range source_map {
		// Use regex to decide if there is an antenna here
		if antenna_re.MatchString(value) {
			// Use this value to hunt for another one
			for key2, value2 := range source_map {
				if value2 == value && key != key2 {
					// Get distance between the two
					distance1_x := key.x - key2.x
					distance2_x := key2.x - key.x
					distance1_y := key.y - key2.y
					distance2_y := key2.y - key.y
					// Place the antinodes
					// Check if the antinodes are within the map bounds, and only add if they are
					antinode1 := coordinate{key.x + distance1_x, key.y + distance1_y}
					if antinode1.x >= 0 && antinode1.x <= map_width && antinode1.y >= 0 && antinode1.y <= map_height {
						all_antinodes = append(all_antinodes, antinode1)
					}
					antinode2 := coordinate{key2.x + distance2_x, key2.y + distance2_y}
					if antinode2.x >= 0 && antinode2.x <= map_width && antinode2.y >= 0 && antinode2.y <= map_height {
						all_antinodes = append(all_antinodes, antinode2)
					}
				}
			}
		}
	}
	all_antinodes = remove_dupes(all_antinodes)
	fmt.Println("Antinodes: ", all_antinodes)
	fmt.Println("Antinode count: ", len(all_antinodes))
	return len(all_antinodes)
}

func main() {
	// Load data
	dat, _ := os.Open("input.txt")

	// Define word search array
	source_map := make(map[coordinate]string)
	// Define values to keep track of the map size
	var map_width int
	var map_height int
	// Loop through input
	// Create a scanner
	scanner := bufio.NewScanner(dat)
	var line_number int
	for scanner.Scan() {
		// Get line
		line := scanner.Text()
		// Split line into letters
		letters := strings.Split(line, "")
		for i, letter := range letters {
			source_map[coordinate{i, line_number}] = letter
			if i > map_width {
				map_width = i
			}
		}
		line_number++
	}

	map_height = line_number

	fmt.Println("Map Width: ", map_width)
	fmt.Println("Map Height: ", map_height)
	map_height--

	// fmt.Println(source_map)

	fmt.Println("Part 1 Output: ", part1(source_map, map_width, map_height))

}
