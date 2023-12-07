package day1

import (
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

func Solution(input string) (solution1, solution2 int) {
	// Solution 1
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		str := scanner.Text()

		var start, end rune
		var s_num string
		var num int

		// check for start digit and break
		for _, n := range str {
			if unicode.IsNumber(n) {
				start = n
				break
			}
		}

		// check for end digit and break
		for i := len(str) - 1; i >= 0; i-- {
			if unicode.IsNumber(rune(str[i])) {
				end = rune(str[i])
				break
			}
		}

		s_num = string(start) + string(end)
		num, _ = strconv.Atoi(s_num)
		solution1 += num
	}

	// Solution 2
	// Commented-out parts are less efficient solution*
	s_digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	scanner = bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		str := scanner.Text()

		start_index, end_index := len(str), -1
		var start, end, s_num string
		var num int

		// find start and end indices of string-denoted digits (comment out for 2nd approach)
		for d_i, s_digit := range s_digits {
			interm_index_s := strings.Index(str, s_digit)
			if interm_index_s != -1 && interm_index_s < start_index {
				start_index = interm_index_s
				start = strconv.Itoa(d_i + 1)
			}

			interm_index_e := strings.LastIndex(str, s_digit)
			if interm_index_e != -1 && interm_index_e > end_index {
				end_index = interm_index_e
				end = strconv.Itoa(d_i + 1)
			}
		}

		// check for start digit in normal form and break
		for _, n := range str[:start_index] { // change index to s_i for 2nd approach and use whole string
			// Uncomment below for 2nd approach
			// for d_i, s_digit := range s_digits {
			// 	if strings.HasPrefix(str[s_i:], s_digit) {
			// 		start = strconv.Itoa(d_i + 1)
			// 		goto start_break
			// 	}
			// }
			if unicode.IsNumber(n) {
				start = string(n)
				break
			}
		}
		// start_break:  // uncomment below for 2nd approach

		// check for end digit in normal form and bresk
		for i := len(str) - 1; i > end_index; i-- { // use whole string in 2nd approach
			// Uncomment below for 2nd approach
			// for d_i, s_digit := range s_digits {
			// 	if strings.HasPrefix(str[i:], s_digit) {
			// 		end = strconv.Itoa(d_i + 1)
			// 		goto end_break
			// 	}
			// }
			if unicode.IsNumber(rune(str[i])) {
				end = string(str[i])
				break
			}
		}
		// end_break:  // uncomment below for 2nd approach

		s_num = start + end
		num, _ = strconv.Atoi(s_num)
		solution2 += num
	}

	return
}
