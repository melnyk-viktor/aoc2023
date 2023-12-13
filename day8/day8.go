package day8

import (
	_ "embed"
	"strings"
)

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}

	// Solution 1
	instructions := strings.Split(input, "\n\n")[0] // Get instructions
	mapping_s := strings.Split(input, "\n\n")[1]    // Get mappings

	var mapping = map[string][2]string{}

	// Build map for fast checks
	for _, line := range strings.Split(mapping_s, "\n") {
		key := strings.Split(line, " = ")[0]
		mapping[key] = [2]string{
			strings.Trim(strings.Split(strings.Split(line, " = ")[1], ", ")[0], "("), // Add L move coordinate
			strings.Trim(strings.Split(strings.Split(line, " = ")[1], ", ")[1], ")"), // Add R move coordinate
		}
	}

	var (
		ii  = 0
		cur = "AAA"
	)

	for {
		if rune(instructions[ii]) == 'R' {
			cur = mapping[cur][1] // Move to R cooerdinate
		} else {
			cur = mapping[cur][0] // Move to L cooerdinate
		}
		solution1++

		if cur == "ZZZ" { // Check if on finish
			break
		}

		// Get next instruction index (can be looped over if finished)
		if ii == len(instructions)-1 {
			ii = 0
		} else {
			ii++
		}
	}

	// Solution 2
	instructions = strings.Split(input, "\n\n")[0] // instructions
	mapping_s = strings.Split(input, "\n\n")[1]    // Get mappings

	mapping = map[string][2]string{}

	for _, line := range strings.Split(mapping_s, "\n") {
		key := strings.Split(line, " = ")[0]
		mapping[key] = [2]string{
			strings.Trim(strings.Split(strings.Split(line, " = ")[1], ", ")[0], "("), // Add L move coordinate
			strings.Trim(strings.Split(strings.Split(line, " = ")[1], ", ")[1], ")"), // Add R move coordinate
		}
	}

	var (
		curs []string
		ress []int
	)
	for k := range mapping { // Get all starting coordinates
		if rune(k[2]) == 'A' {
			curs = append(curs, k)
		}
	}

	for _, cur := range curs { // For every starting coordinate
		var (
			count = 0
			ii    = 0
		)

		for {
			if rune(instructions[ii]) == 'R' {
				cur = mapping[cur][1] // Move to R cooerdinate
			} else {
				cur = mapping[cur][0] // Move to L cooerdinate
			}
			count++

			if rune(cur[2]) == 'Z' { // Check if on finish
				break
			}

			// Get next instruction index (can be looped over if finished)
			if ii == len(instructions)-1 {
				ii = 0
			} else {
				ii++
			}
		}
		ress = append(ress, count)
	}

	// Find first sync point when all starts lead to ends using LCM
	solution2 = LCM(ress...)

	return
}

// LCM implementation adjusted from https://go.dev/play/p/SmzvkDjYlb

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	if len(integers) == 0 {
		return 0
	} else if len(integers) == 1 {
		return integers[0]
	}

	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

//go:embed test_data.txt
var TEST_INPUT string
