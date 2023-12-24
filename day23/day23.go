package day23

import (
	_ "embed"
	"slices"
	"strings"
)

var moves = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}
	parsed := strings.Split(input, "\n")

	var (
		start   [2]int
		goal    [2]int
		running map[[2]int][][3]int // map of moves from point with weights
	)

	// Solution 1
	start = [2]int{0, 1}
	goal = [2]int{len(parsed) - 1, len(parsed[0]) - 2}
	running = map[[2]int][][3]int{}
	for i, r := range parsed {
		for j, c := range r {
			if c == '.' { // Path
				for _, m := range moves {
					// If move leads out of bounds, skip it
					if i+m[0] < 0 || i+m[0] >= len(parsed) || j+m[1] < 0 || j+m[1] >= len(r) {
						continue
					}

					if parsed[i+m[0]][j+m[1]] == '.' { // Check move point
						// Connect points with weight 1
						insertRunning(running, [2]int{i, j}, [3]int{i + m[0], j + m[1], 1})
						insertRunning(running, [2]int{i + m[0], j + m[1]}, [3]int{i, j, 1})
					}
				}
			}

			if c == '>' { // Path
				// Connect points with weight 1
				insertRunning(running, [2]int{i, j}, [3]int{i, j + 1, 1})
				insertRunning(running, [2]int{i, j - 1}, [3]int{i, j, 1})
			}

			if c == 'v' { // Path
				// Connect points with weight 1
				insertRunning(running, [2]int{i, j}, [3]int{i + 1, j, 1})
				insertRunning(running, [2]int{i - 1, j}, [3]int{i, j, 1})
			}
		}
	}

	solution1 = walk(running, start, goal)

	// Solution 2
	running = map[[2]int][][3]int{}
	for i, r := range parsed {
		for j, c := range r {
			if c == '.' || c == '>' || c == 'v' { // Path
				for _, m := range moves {
					// If move leads out of bounds, skip it
					if i+m[0] < 0 || i+m[0] >= len(parsed) || j+m[1] < 0 || j+m[1] >= len(r) {
						continue
					}
					if parsed[i+m[0]][j+m[1]] == '.' || parsed[i+m[0]][j+m[1]] == '>' || parsed[i+m[0]][j+m[1]] == 'v' { // Check move point
						// Connect points with weight 1
						insertRunning(running, [2]int{i, j}, [3]int{i + m[0], j + m[1], 1})
						insertRunning(running, [2]int{i + m[0], j + m[1]}, [3]int{i, j, 1})
					}
				}
			}
		}
	}

	// Truncate graph, only to points with crossroads
	for len(running) > 0 {
		done := false // Flag, no more nodes to truncate
		for k, v := range running {
			if len(v) == 2 { // Node is a simple connecting node
				from := v[0]
				to := v[1]

				// Remove connecting node from "parent" node connections
				i := slices.Index(running[[2]int{from[0], from[1]}], [3]int{k[0], k[1], from[2]})
				running[[2]int{from[0], from[1]}] = append(running[[2]int{from[0], from[1]}][:i], running[[2]int{from[0], from[1]}][i+1:]...)

				// Remove connecting node from "child" node connections
				i = slices.Index(running[[2]int{to[0], to[1]}], [3]int{k[0], k[1], to[2]})
				running[[2]int{to[0], to[1]}] = append(running[[2]int{to[0], to[1]}][:i], running[[2]int{to[0], to[1]}][i+1:]...)

				// Connect "parent" & "child" nodes, increasing weight
				running[[2]int{from[0], from[1]}] = append(running[[2]int{from[0], from[1]}], [3]int{to[0], to[1], to[2] + from[2]})
				running[[2]int{to[0], to[1]}] = append(running[[2]int{to[0], to[1]}], [3]int{from[0], from[1], from[2] + to[2]})

				delete(running, k) // Remove connecting node from graph
				done = false
				break
			}
			done = true // Filtered out all connecting nodes
		}
		if done {
			break
		}
	}

	solution2 = walk(running, start, goal)

	return
}

func walk(running map[[2]int][][3]int, start [2]int, goal [2]int) (res int) {
	q := [][3]int{}
	crossed := map[[2]int]bool{} // Map of crossed nodes

	q = append(q, [3]int{start[0], start[1], 0})
	for len(q) > 0 {
		// Pop
		cur := q[len(q)-1]
		q = q[:len(q)-1]

		if cur[2] == -1 { // Check for going back
			crossed[[2]int{cur[0], cur[1]}] = false
		} else if cur[0] == goal[0] && cur[1] == goal[1] { // Check if at goal
			res = max(res, cur[2])
		} else if v, ok := crossed[[2]int{cur[0], cur[1]}]; !ok || !v { // Check if not visited
			crossed[[2]int{cur[0], cur[1]}] = true // Mark visit
			step := cur[2]                         // Tmp val
			cur[2] = -1                            // Mark move back
			q = append(q, cur)                     // Add mark
			for _, m := range running[[2]int{cur[0], cur[1]}] {
				q = append(q, [3]int{m[0], m[1], step + m[2]}) // Move forwards
			}
		}
	}
	return
}

func insertRunning(running map[[2]int][][3]int, key [2]int, value [3]int) {
	if v, ok := running[key]; ok {
		if !slices.Contains(v, value) {
			running[key] = append(running[key], value)
		}
	} else {
		running[key] = [][3]int{value}
	}
}

//go:embed test_data.txt
var TEST_INPUT string
