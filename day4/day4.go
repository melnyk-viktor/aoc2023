package day4

import (
	"slices"
	"strings"
)

func Solution(input string) (solution1, solution2 int) {
	// Solution 1
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var inter_sum int
		str := strings.Split(line, ": ")[1]

		winning_numbers := strings.Split(strings.Split(str, " | ")[0], " ")
		got_numbers := strings.Split(strings.Split(str, " | ")[1], " ")

		for _, winning_number := range winning_numbers {
			if slices.Contains(got_numbers, winning_number) && winning_number != "" {
				// NOTE: below if requires improvement
				if inter_sum < 1 {
					inter_sum += 1
				} else {
					inter_sum *= 2
				}
			}
		}
		solution1 += inter_sum
	}

	// Solution 2
	lines = strings.Split(input, "\n")
	var line_counts = make([]int, len(lines)) // Repetitions array
	for i, line := range lines {
		str := strings.Split(line, ": ")[1]
		line_counts[i]++

		winning_numbers := strings.Split(strings.Split(str, " | ")[0], " ")
		got_numbers := strings.Split(strings.Split(str, " | ")[1], " ")

		win_runner := i 
		for _, winning_number := range winning_numbers {
			// If winning number, run forward and propagate
			if slices.Contains(got_numbers, winning_number) && winning_number != "" {
				win_runner++
				line_counts[win_runner] += line_counts[i]
			}
		}
		solution2 += line_counts[i]
	}

	return
}
