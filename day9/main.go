package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type block struct {
	id     int
	length int
	file   bool
}

func filesystem_to_string(filesystem []block) string {
	var sb strings.Builder
	for _, block := range filesystem {
		if block.file {
			// Get a string of the ID to repeat
			this_id := strconv.Itoa(block.id)
			// Repeat for length
			repeated_id := strings.Repeat(this_id, block.length)
			// fmt.Print(repeated_id)
			sb.WriteString(repeated_id)
		} else {
			// Repeat dot for length
			repeated_dot := strings.Repeat(".", block.length)
			// fmt.Print(repeated_dot)
			sb.WriteString(repeated_dot)
		}
	}
	// fmt.Println("")
	return sb.String()
}

func parse_filesystem(line string) []block {
	var filesystem []block
	// Split the string
	entries := strings.Split(line, "")
	var file bool = true
	var id int
	for _, entry := range entries {
		entry_length, _ := strconv.Atoi(entry)
		if file {
			filesystem = append(filesystem, block{id: id, length: entry_length, file: file})
			id++
			file = false
		} else {
			filesystem = append(filesystem, block{id: -1, length: entry_length, file: file})
			file = true
		}
	}
	return filesystem
}

func checksum(filesystem []string) int {
	var total int
	for index, element := range filesystem {
		if element == "." {
			continue
		} else {
			el_num, _ := strconv.Atoi(element)
			total += index * el_num
		}
	}
	return total
}

func part1(filesystem string) int {
	// Move blocks around on the string
	// Kind of a 'defrag' operation?
	fs_slice := strings.Split(filesystem, "")
	// Go bakcwards in slice until we
	// find a number, then swap that for the first dot we find in the string
	for i := len(fs_slice) - 1; i >= 0; i-- {
		if fs_slice[i] != "." {
			// Find the first dot in the slice
			for j := 0; j < len(fs_slice); j++ {
				if fs_slice[j] == "." {
					// Swap the dot for the last number
					// fmt.Println("Swapping ", i, " into ", j)
					fs_slice[j] = fs_slice[i]
					fs_slice[i] = "."
					break
				}
			}
		}
	}
	// Above leaves a dot at the front that should be right at the end, shift the entire slice over by -1
	fs_slice = append(fs_slice[1:], fs_slice[0])
	// fmt.Println(strings.Join(fs_slice, ""))
	return checksum(fs_slice)
}

func main() {
	// Load daata
	dat, _ := os.Open("./input.txt")

	var filesystem []block

	// Create a scanner
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		line := scanner.Text()
		// Parse the 'line' (there's only one) as the filesystem
		// fmt.Println(line)
		// Split the string
		filesystem = parse_filesystem(line)
	}
	fs_string := filesystem_to_string(filesystem)
	// fmt.Println(fs_string)
	fmt.Println("Part 1 Output: ", part1(fs_string))
}
