package day13

import (
	_ "embed"
	"strings"
)

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}

	patterns := strings.Split(input, "\n\n")

	// Solution 1
OUTER:
	for _, pt := range patterns {
		grid := strings.Split(pt, "\n")

		for i := 0; i < len(grid)-1; i++ {
			l := grid[:i+1] // Left side
			r := grid[i+1:] // Right side

			check := true                              // Reflection flag
			for j := 0; j < min(len(l), len(r)); j++ { // Move the corner
				if l[len(l)-j-1] != r[j] { // Check reflection
					check = false
					break
				}
			}
			if check { // Add run of reflection until border
				solution1 += 100 * len(l)
				continue OUTER
			}
		}

		grid = transpose(grid)
		for i := 0; i < len(grid)-1; i++ {
			l := grid[:i+1] // Left side (transpositioned up)
			r := grid[i+1:] // Right side (transpositioned down)

			check := true                              // Reflection flag
			for j := 0; j < min(len(l), len(r)); j++ { // Move the corner
				if l[len(l)-j-1] != r[j] { // Check reflection
					check = false
					break
				}
			}
			if check { // Add run of reflection until border
				solution1 += len(l)
				break
			}
		}
	}

	// Solution 2
OUTER_C:
	for _, pt := range patterns {
		grid := strings.Split(pt, "\n")
		for i := 0; i < len(grid)-1; i++ {
			l := grid[:i+1] // Left side
			r := grid[i+1:] // Right side

			check := 0                                 // Reflection flag accounting for smudging allows for 1 smudge
			for j := 0; j < min(len(l), len(r)); j++ { // Move the corner
				if l[len(l)-j-1] != r[j] { // Check reflection
					check = smudger(l[len(l)-j-1], r[j])
					break
				}
			}
			if check <= 1 {
				solution2 += 100 * len(l)
				continue OUTER_C
			}
		}

		grid = transpose(grid)
		for i := 0; i < len(grid)-1; i++ {
			l := grid[:i+1] // Left side (transpositioned up)
			r := grid[i+1:] // Right side (transpositioned down)

			check := 0                                 // Reflection flag accounting for smudging allows for 1 smudge
			for j := 0; j < min(len(l), len(r)); j++ { // Move the corner
				if l[len(l)-j-1] != r[j] { // Check reflection
					check = smudger(l[len(l)-j-1], r[j])
					break
				}
			}
			if check <= 1 {
				solution2 += len(l)
				break
			}
		}
	}

	return
}

// Grid transposition on []string
func transpose(grid []string) (n_grid []string) {
	n_grid = make([]string, len(grid[0]))
	for i := 0; i < len(n_grid); i++ {
		for j := 0; j < len(grid); j++ {
			n_grid[i] += string(grid[j][i])
		}
	}
	return
}

/*
Allows for one smudge in reflection of given lines to be ignored.
(Taken out from main func due to nesting and bad readability)
*/
func smudger(l1, l2 string) int {
	var unsmudged int // Allows for unsmudging
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] { // Found difference, smudge or break
			if unsmudged > 1 {
				break // Unsmudged more tnah once, not really an optimization
			}
			unsmudged++
		}
	}
	return unsmudged
}

//go:embed test_data.txt
var TEST_INPUT string
