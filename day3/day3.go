package day3

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func Solution(input string) (solution1, solution2 int) {
	// Solution 1
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for j := 0; j < len(line); j++ { // With jumps
			var res string
			var part_num bool

			// Check if we found a number
			if unicode.IsNumber(rune(line[j])) {
				y := int(math.Max(0, float64(i-1)))
				y_max := int(math.Min(float64(i+1), float64(len(lines)-1)))
				x := int(math.Max(0, float64(j-1)))

				// Check if number is a part number and handle full number
				for x_run := x; x_run < len(line); x_run++ { // Number length iterations + 2
					for y_run := y; y_run <= y_max; y_run++ { // At most 3 iterations
						if lines[y_run][x_run] != '.' && !unicode.IsNumber(rune(lines[y_run][x_run])) {
							part_num = true
						}
					}

					// Grow resulting number or if fully grown, jump
					if !unicode.IsNumber(rune(line[x_run])) && x != x_run {
						j = x_run // Jump
						break
					} else if unicode.IsNumber(rune(line[x_run])) {
						res += string(line[x_run])
					}
				}

				if part_num {
					val, _ := strconv.Atoi(res)
					solution1 += val
				}
			}
		}
	}

	// Solution 2
	// TODO: optimize by looking for gears and grow numbers into both/one direction depending on whe way they neighbor gear
	lines = strings.Split(input, "\n")
	var gears = map[string][]int{}
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			var res string
			var gear_index string

			// Check if we found a number
			if unicode.IsNumber(rune(line[j])) {
				y := int(math.Max(0, float64(i-1)))
				y_max := int(math.Min(float64(i+1), float64(len(lines)-1)))
				x := int(math.Max(0, float64(j-1)))

				// Check if number is a part number and handle full number
				for x_run := x; x_run < len(line); x_run++ { // Number length iterations + 2
					for y_run := y; y_run <= y_max; y_run++ { // At most 3 iterations
						if lines[y_run][x_run] == '*' {
							gear_index = strconv.Itoa(y_run) + "," + strconv.Itoa(x_run) // Map all * (gears)
						}
					}

					// Grow resulting number or if fully grown, jump
					if !unicode.IsNumber(rune(line[x_run])) && x != x_run {
						j = x_run // Jump
						break
					} else if unicode.IsNumber(rune(line[x_run])) {
						res += string(line[x_run])
					}
				}

				// Add num to gear mapping
				if gear_index != "" {
					val, _ := strconv.Atoi(res)
					gears[gear_index] = append(gears[gear_index], val)
				}
			}
		}
	}

	// separate gears and sum gear ratios
	for _, v := range gears {
		if len(v) > 1 {
			solution2 += v[0] * v[1]
		}
	}

	return
}
