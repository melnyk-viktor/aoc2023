package day12

import (
	_ "embed"
	"strconv"
	"strings"
)

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}

	// Solution 1 & 2
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		var (
			rs       string
			new_rs   string
			runs     []int
			new_runs []int
			cache    map[float64]int
		)
		row_split := strings.Split(row, " ")

		rs = row_split[0]
		runs = []int{}
		for _, rc := range strings.Split(row_split[1], ",") {
			rci, _ := strconv.Atoi(rc)
			runs = append(runs, rci)
		}
		solution1 += recur(rs, runs, nil)

		new_rs = rs
		new_runs = runs
		for i := 0; i < 4; i++ {
			new_rs = strings.Join([]string{new_rs, rs}, "?")
			new_runs = append(new_runs, runs...)
		}

		cache = map[float64]int{}
		solution2 += recur(new_rs, new_runs, cache)
	}

	return
}

// Cached Using map, to map integers, Cantor pairing function is used
func recur(line string, runs []int, cache map[float64]int) (res int) {
	uid := 0.5*float64(len(line)+len(runs))*float64(len(line)+len(runs)+1) + float64(len(runs)) // Calculate cache key

	defer func() { // Update cache after all cchecks
		if cache != nil {
			cache[uid] = res
		}
	}()

	if len(line) == 0 {
		if len(runs) == 0 {
			res = 1 // Runs & line done
		}
		return // Line ended
	}

	if cache != nil {
		if v, ok := cache[uid]; ok {
			return v // Return cache if found
		}
	}

	switch line[0] {
	case '?':
		res = recur(line[1:], runs, cache) + recur("#"+line[1:], runs, cache) // Check both "." and "#" cases going further down with current run
		return
	case '#':
		if len(runs) == 0 {
			return // No more runs to check
		} else if len(line) < runs[0] {
			return // Run does not fit
		} else if strings.Count(line[:runs[0]], ".") > 0 {
			return // Run does not fit due toi a "."
		} else if len(runs) > 1 {
			if len(line) > runs[0] && line[runs[0]] != '#' {
				res = recur(line[runs[0]+1:], runs[1:], cache) // Jump the run in line, go further down with next run
			}
			return
		} else {
			res = recur(line[runs[0]:], runs[1:], cache) // Jump the run in line, go further down with next run
			return
		}
	case '.':
		res = recur(line[1:], runs, cache) // Go further with current run
		return
	}
	panic("Error in branching")
}

//go:embed test_data.txt
var TEST_INPUT string
