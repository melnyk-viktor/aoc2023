package day14

import (
	_ "embed"
	"strings"
)

var RUN_CYCLES = 1000000000

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}

	// Solution 1
	rows := strings.Split(input, "\n")
	for i := 0; i < len(rows); i++ {
		for j := range rows[i] {
			// If stone that can move, move it north until possible
			if rows[i][j] == 'O' {
				pi := i
				for pi >= 1 && rows[pi-1][j] == '.' {
					rows[pi] = rows[pi][:j] + "." + rows[pi][j+1:]
					pi--
					rows[pi] = rows[pi][:j] + "O" + rows[pi][j+1:]
				}
			}
		}
	}
	for i, r := range rows {
		solution1 += strings.Count(r, "O") * (len(rows) - i) // Calculate load per row and add to result
	}

	// Solution 2
	var (
		s_rows = input
		seen   = map[string]int{} // Map of seen configurations to their iteration number
		cycle  = [][2]int{
			{-1, 0}, // North
			{0, -1}, // West
			{1, 0},  // South
			{0, 1},  // East
		}
	)

out: // Loop breaking label
	for i := 1; i < RUN_CYCLES-1; i++ {
		for _, m := range cycle {
			s_rows = run_move(s_rows, m[0], m[1]) // Run move with specified direction in cycle
		}
		if val, ok := seen[s_rows]; ok { // Check if configuration of rocks was seen before
			shifted := val + (RUN_CYCLES-val)%(i-val) // Calculate which configuration in detected cycle we will end up with
			// Build final configutation
			for k, v := range seen {
				if v == shifted {
					rows = strings.Split(k, "\n")
					break out
				}
			}
		}
		seen[s_rows] = i // Update mapping of rows we've seen and their indexes
	}
	for i, r := range rows {
		solution2 += strings.Count(r, "O") * (len(rows) - i) // Calculate load per row and add to result
	}

	return
}

func run_move(s_rows string, xc int, yc int) (res string) {
	rows := strings.Split(s_rows, "\n") // Split into slice for ease of manipulations
	mip := true                         // Check for move in progress
	for mip {
		mip = false
		for i := range rows {
			if i+xc >= len(rows) || i+xc < 0 {
				continue
			}
			for j := range rows[i] {
				if j+yc >= len(rows) || j+yc < 0 {
					continue
				}
				if rows[i][j] == 'O' && rows[i+xc][j+yc] == '.' {
					mip = true // Move in progress

					// Shift rock that can move
					rows[i] = rows[i][:j] + "." + rows[i][j+1:]
					rows[i+xc] = rows[i+xc][:j+yc] + "O" + rows[i+xc][j+yc+1:]
				}
			}
		}
	}
	return strings.Join(rows, "\n") // Join up to get same format as input
}

//go:embed test_data.txt
var TEST_INPUT string
