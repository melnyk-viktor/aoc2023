package day8

import (
	"strings"
)

func Solution(input string) (solution1, solution2 int) {
	// Solution 1
	instructions := strings.Split(input, "\n\n")[0]
	mapping_s := strings.Split(input, "\n\n")[1]

	var mapping = map[string][2]string{}

	for _, line := range strings.Split(mapping_s, "\n") {
		key := strings.Split(line, " = ")[0]
		mapping[key] = [2]string{
			strings.Trim(strings.Split(strings.Split(line, " = ")[1], ", ")[0], "("),
			strings.Trim(strings.Split(strings.Split(line, " = ")[1], ", ")[1], ")"),
		}
	}

	var (
		ii  = 0
		cur = "AAA"
	)

	for {
		if rune(instructions[ii]) == 'R' {
			cur = mapping[cur][1]
		} else {
			cur = mapping[cur][0]
		}
		solution1++

		if cur == "ZZZ" {
			break
		}

		if ii == len(instructions)-1 {
			ii = 0
		} else {
			ii++
		}
	}

	// Solution 2
	instructions = strings.Split(input, "\n\n")[0]
	mapping_s = strings.Split(input, "\n\n")[1]

	mapping = map[string][2]string{}

	for _, line := range strings.Split(mapping_s, "\n") {
		key := strings.Split(line, " = ")[0]
		mapping[key] = [2]string{
			strings.Trim(strings.Split(strings.Split(line, " = ")[1], ", ")[0], "("),
			strings.Trim(strings.Split(strings.Split(line, " = ")[1], ", ")[1], ")"),
		}
	}

	var (
		curs []string
		ress []int
	)
	for k := range mapping {
		if rune(k[2]) == 'A' {
			curs = append(curs, k)
		}
	}

	for _, cur := range curs {
		var (
			count = 0
			ii    = 0
		)

		for {
			if rune(instructions[ii]) == 'R' {
				cur = mapping[cur][1]
			} else {
				cur = mapping[cur][0]
			}
			count++

			if rune(cur[2]) == 'Z' {
				break
			}

			if ii == len(instructions)-1 {
				ii = 0
			} else {
				ii++
			}
		}
		ress = append(ress, count)
	}
	if len(ress) > 1 {
		solution2 = LCM(ress[0], ress[1], ress...)
	} else {
		solution1 = ress[0]
	}

	return
}

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
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
