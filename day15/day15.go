package day15

import (
	_ "embed"
	"strconv"
	"strings"
)

type lens struct {
	label string
	fl    int
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}
	// Solution 1
	cmds := strings.Split(strings.ReplaceAll(input, "\n", ""), ",") // Split input data into cmds
	for _, cmd := range cmds {
		solution1 += get_hash(cmd) // Get hash for cmd in cmds and append to result
	}

	// Solution 2
	boxwall := [256][]lens{}
	for _, cmd := range cmds {
		// Get label and fl if present from cmd
		c_parts := strings.FieldsFunc(cmd, func(r rune) bool {
			return r == '=' || r == '-'
		})
		bn := get_hash(c_parts[0]) // Get hash for label

		if cmd[len(cmd)-1] == '-' { // Remove cmd
			// Find and remove lens from box
			for i, l := range boxwall[bn] {
				if l.label == c_parts[0] {
					boxwall[bn] = append(boxwall[bn][:i], boxwall[bn][i+1:]...)
				}
			}
		} else if strings.Contains(cmd, "=") { // Add/replace cmd
			inserted := false
			// Find and replace with new lens
			for i, l := range boxwall[bn] {
				if l.label == c_parts[0] {
					fl, _ := strconv.Atoi(c_parts[1])
					boxwall[bn][i] = lens{label: c_parts[0], fl: fl}
					inserted = true
					break
				}
			}
			// Insert new lens
			if !inserted {
				fl, _ := strconv.Atoi(c_parts[1])
				boxwall[bn] = append(boxwall[bn], lens{label: c_parts[0], fl: fl})
			}
		}
	}
	for i, b := range boxwall {
		for j, l := range b {
			solution2 += (i + 1) * (j + 1) * l.fl // Calculate focusing power and add to result
		}
	}

	return
}

// Generate hash from string
func get_hash(cmd string) (res int) {
	for _, c := range cmd {
		res += int(rune(c))
		res *= 17
		res %= 256
	}
	return
}

//go:embed test_data.txt
var TEST_INPUT string
