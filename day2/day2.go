package day2

import (
	"bufio"
	"strconv"
	"strings"
)

func Solution(input string) (solution1, solution2 int) {
	// Solution 1
	var count = 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		count += 1
		str := strings.Split(scanner.Text(), ": ")[1]

		game_turns := strings.Split(str, "; ")
		game := map[string]int{"green": 0, "red": 0, "blue": 0}

		for _, s_turn := range game_turns {
			colors := strings.Split(s_turn, ", ")

			for _, s_color_val := range colors {
				c_val := strings.Split(s_color_val, " ")
				color := c_val[1]
				value, _ := strconv.Atoi(c_val[0])

				if game[color] < value {
					game[color] = value
				}
			}
		}

		// NOTE: Long if can be separeted into 3 ifs with reverse logic and continue, with increment after them
		if COLOR_MAX["red"] >= game["red"] && COLOR_MAX["green"] >= game["green"] && COLOR_MAX["blue"] >= game["blue"] {
			solution1 += count
		}
	}

	// Solution 2
	count = 0
	scanner = bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		count += 1
		str := strings.Split(scanner.Text(), ": ")[1]

		game_turns := strings.Split(str, "; ")
		game := map[string]int{"green": 0, "red": 0, "blue": 0}

		for _, s_turn := range game_turns {
			colors := strings.Split(s_turn, ", ")

			for _, s_color_val := range colors { // Runs 3 timex
				c_val := strings.Split(s_color_val, " ")
				color := c_val[1]
				value, _ := strconv.Atoi(c_val[0])

				if game[color] < value {
					game[color] = value
				}
			}
		}
		solution2 += game["red"] * game["green"] * game["blue"]
	}

	return
}

var COLOR_MAX = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}
