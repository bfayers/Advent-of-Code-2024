package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/life4/genesis/slices"
)

type coordinate struct {
	x int
	y int
}

var antenna_re = regexp.MustCompile(`([A-Z]|[a-z]|[0-9])`)

func part1(satellites map[string][]coordinate, map_width int, map_height int) int {
	var all_antinodes []coordinate

	// Iterate over all satellite positions
	for _, value := range satellites {
		// Generate all combos of satellite positions
		combinations := slices.Permutations(value, 2)
		for combo := range combinations {
			pos1, pos2 := combo[0], combo[1]
			// Get distance between the two
			distance1_x := pos1.x - pos2.x
			distance1_y := pos1.y - pos2.y

			distance2_x := pos2.x - pos1.x
			distance2_y := pos2.y - pos1.y
			// Place the antinodes
			// Check if the antinodes are within the map bounds, and only add if they are
			antinode1 := coordinate{pos1.x + distance1_x, pos1.y + distance1_y}
			if antinode1.x >= 0 && antinode1.x <= map_width && antinode1.y >= 0 && antinode1.y <= map_height {
				all_antinodes = append(all_antinodes, antinode1)
			}
			antinode2 := coordinate{pos2.x + distance2_x, pos2.y + distance2_y}
			if antinode2.x >= 0 && antinode2.x <= map_width && antinode2.y >= 0 && antinode2.y <= map_height {
				all_antinodes = append(all_antinodes, antinode2)
			}
		}
	}

	all_antinodes = slices.Uniq(all_antinodes)
	return len(all_antinodes)
}

func part2(satellites map[string][]coordinate, map_width int, map_height int) int {
	var all_antinodes []coordinate
	// Iterate over all satellite positions
	for _, value := range satellites {
		// Generate all combos of satellite positions
		combinations := slices.Permutations(value, 2)
		for combo := range combinations {
			pos1, pos2 := combo[0], combo[1]
			// Get distance between the two
			distance1_x := pos1.x - pos2.x
			distance1_y := pos1.y - pos2.y

			distance2_x := pos2.x - pos1.x
			distance2_y := pos2.y - pos1.y
			// Place the antinodes
			// Check if the antinodes are within the map bounds, and only add if they are
			antinode1 := coordinate{pos1.x, pos1.y}
			for {
				if antinode1.x >= 0 && antinode1.x <= map_width && antinode1.y >= 0 && antinode1.y <= map_height {
					all_antinodes = append(all_antinodes, antinode1)
				} else {
					break
				}
				antinode1.x += distance1_x
				antinode1.y += distance1_y
			}
			antinode2 := coordinate{pos2.x, pos2.y}
			for {
				if antinode2.x >= 0 && antinode2.x <= map_width && antinode2.y >= 0 && antinode2.y <= map_height {
					all_antinodes = append(all_antinodes, antinode2)
				} else {
					break
				}
				antinode2.x += distance2_x
				antinode2.y += distance2_y
			}
		}
	}
	all_antinodes = slices.Uniq(all_antinodes)
	return len(all_antinodes)
}

func main() {
	// Load data
	dat, _ := os.Open("input.txt")

	// Define map to store the map
	// source_map := make(map[coordinate]string)
	// Define a map to store satellites
	satellites := make(map[string][]coordinate)
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
			// source_map[coordinate{i, line_number}] = letter
			// Check if the letter is an antenna
			if antenna_re.MatchString(letter) {
				// Add the antenna to the satellite map
				satellites[letter] = append(satellites[letter], coordinate{i, line_number})
			}
			if i > map_width {
				map_width = i
			}
		}
		line_number++
	}

	map_height = line_number - 1

	fmt.Println("Map Width: ", map_width)
	fmt.Println("Map Height: ", map_height)

	fmt.Println("Part 1 Output: ", part1(satellites, map_width, map_height))
	fmt.Println("Part 2 Output: ", part2(satellites, map_width, map_height))

}
