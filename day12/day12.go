package day12

import (
	"strconv"
	"strings"
)

func Solution(input string) (solution1, solution2 int) {
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
		solution1 += recur2(rs, runs, nil)

		new_rs = rs
		new_runs = runs
		for i := 0; i < 4; i++ {
			new_rs = strings.Join([]string{new_rs, rs}, "?")
			new_runs = append(new_runs, runs...)
		}

		cache = map[float64]int{}
		solution2 += recur2(new_rs, new_runs, cache)
		// fmt.Println(cache)
	}

	return
}

// Cached Using map, to map integers, Cantor pairing function is used
func recur2(line string, runs []int, cache map[float64]int) (res int) {
	uid := 0.5*float64(len(line)+len(runs))*float64(len(line)+len(runs)+1) + float64(len(runs)) // Calculate cache key

	if len(line) == 0 {
		if len(runs) == 0 {
			res++ // Runs & line done
		}
		return // Line ended
	}

	if cache != nil {
		if v, ok := cache[uid]; ok {
			return v // Return cache if found
		}
	}

	if line[0] == '.' {
		res = recur2(line[1:], runs, cache) // Move further, trimming "."
	} else {
		if line[0] == '?' {
			res += recur2(line[1:], runs, cache) // Move further, trimming "."
		}
		if len(runs) > 0 {
			count := 0 // Acceptable rune counter
			for _, char := range line {
				if count > runs[0] || char == '.' || count == runs[0] && char == '?' {
					break // No point checking further
				}
				count += 1
			}

			if count == runs[0] {
				if count < len(line) && line[count] != '#' {
					res += recur2(line[count+1:], runs[1:], cache) // Jump all counted and done run, and move firther
				} else {
					res += recur2(line[count:], runs[1:], cache) // Jump all counted, last one is "." anyway, and done run, and move firther
				}
			}
		}
	}
	if cache != nil {
		cache[uid] = res // Update cache
	}
	return
}
