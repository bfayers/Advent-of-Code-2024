package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/life4/genesis/slices"
)

type coordinate struct {
	x int
	y int
}

func process(satellites map[string][]coordinate, map_width int, map_height int, part2 bool) int {
	var all_antinodes []coordinate

	// Iterate over all satellite positions
	for _, value := range satellites {
		// Generate all combos of satellite positions
		combinations := slices.Permutations(value, 2)
		for combo := range combinations {
			pos1, pos2 := combo[0], combo[1]
			// Get distance between the two
			distance_x := pos1.x - pos2.x
			distance_y := pos1.y - pos2.y

			// Place the antinodes
			if !part2 {
				antinode := coordinate{pos1.x + distance_x, pos1.y + distance_y}
				// Check if the antinodes are within the map bounds, and only add if they are
				if antinode.x >= 0 && antinode.x <= map_width && antinode.y >= 0 && antinode.y <= map_height {
					all_antinodes = append(all_antinodes, antinode)
				}
			} else {
				antinode := coordinate{pos1.x, pos1.y}
				// Go over all positions on the 'line' until we hit the map bounds for part2
				for {
					// Check if the antinodes are within the map bounds, and only add if they are
					if antinode.x >= 0 && antinode.x <= map_width && antinode.y >= 0 && antinode.y <= map_height {
						all_antinodes = append(all_antinodes, antinode)
					} else {
						break
					}
					antinode.x += distance_x
					antinode.y += distance_y
				}
			}
		}
	}

	all_antinodes = slices.Uniq(all_antinodes)
	return len(all_antinodes)
}

func main() {
	// Load data
	dat, _ := os.Open("input.txt")

	// Define a map to store satellites
	satellites := make(map[string][]coordinate)
	// Define values to keep track of the map size
	var map_width int
	var map_height int
	// Create a scanner
	scanner := bufio.NewScanner(dat)
	// Loop through input
	for scanner.Scan() {
		// Get line
		line := scanner.Text()
		// Split line into letters
		letters := strings.Split(line, "")
		map_width = len(letters) - 1
		for i, letter := range letters {
			// Check if the letter is an antenna (if it's not a dot it's an antenna)
			if letter != "." {
				// Add the antenna to the satellite map
				satellites[letter] = append(satellites[letter], coordinate{i, map_height})
			}
		}
		map_height++
	}

	// Account for zero indexing (?)
	map_height--

	fmt.Println("Part 1 Output: ", process(satellites, map_width, map_height, false))
	fmt.Println("Part 2 Output: ", process(satellites, map_width, map_height, true))

}
