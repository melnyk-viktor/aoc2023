package day21

import (
	_ "embed"
	"slices"
	"strings"
)

var (
	maxSteps1 = 64        // Max steps part 1
	maxSteps2 = 26501365  // Max steps p2
	moves     = [][2]int{ // Possible step directions
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
)

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}
	m := strings.Split(input, "\n")

	// Get start point
	var start [4]int
	for x := range m {
		for y := range m[x] {
			if m[x][y] == 'S' {
				start = [4]int{x, y, 0, 0}
			}
		}
	}

	// Solution 1
	// Use modified BFS that oscilates and runs for only 64 steps
	res1 := unbfs(m, start, maxSteps1)
	for _, v := range res1 {
		if v%2 == 0 { // Oscilation left this coordinates reachable
			solution1++
		}
	}

	// Solution 2
	// Use modified BFS that oscilates and runs for rollover plains from -3 up to 3
	res2 := mbfs(m, start)

	ngs := []int{-3, -2, -1, 0, 1, 2, 3} // Additional plains (both x and y). Limit is result of guess work
	counted := map[[2]int]int{}          // Account of counted points

	// Check all x, y coordinates of original plane
	for x := range m {
		for y := range m[x] {
			p := [4]int{x, y, 0, 0}
			if _, ok := res2[p]; ok { // Check if coordinate is a reachable point in original plane

				// Check all reflections of points on original plane on reflected planes
				for _, sx := range ngs {
					for _, sy := range ngs {
						p = [4]int{x, y, sx, sy} // Point on reflected plane
						nc := res2[p]            // Parity

						if nc%2 == maxSteps2%2 && nc <= maxSteps2 { // account for initial parity
							solution2 += 1
						}

						if (sx == slices.Min(ngs) || sx == slices.Max(ngs)) && (sy == slices.Min(ngs) || sy == slices.Max(ngs)) { // If rollover on both axes
							pack := [2]int{nc, 2}
							c := 0 // Running count

							// Check if visited same cases
							if v, ok := counted[pack]; ok {
								solution2 += v
								continue
							}

							for cnt := 1; cnt < (maxSteps2-nc)/len(m[x])+1; cnt++ { // Account for all occurrences in maxStep2 number of steps
								if nc+len(m[x])*cnt < maxSteps2+1 && (nc+len(m[x])*cnt)%2 == maxSteps2%2 { // Count the reachable ones witch fit parity
									c += cnt + 1
								}
							}
							counted[pack] = c
							solution2 += c

						} else if (sx == slices.Min(ngs) || sx == slices.Max(ngs)) || (sy == slices.Min(ngs) || sy == slices.Max(ngs)) {
							pack := [2]int{nc, 1}
							c := 0 // Running count

							// Check if visited same cases
							if v, ok := counted[pack]; ok {
								solution2 += v
								continue
							}

							for cnt := 1; cnt < (maxSteps2-nc)/len(m[x])+1; cnt++ { // Account for all occurrences in maxStep2 number of steps
								if nc+len(m[x])*cnt < maxSteps2+1 && (nc+len(m[x])*cnt)%2 == maxSteps2%2 { // Count the reachable ones witch fit parity
									c++
								}
							}
							counted[pack] = c
							solution2 += c
						}
					}
				}
			}
		}
	}
	return
}

// BFS modified with oscilating nodes and runs with limite number of steps
func unbfs(m []string, s [4]int, goal int) map[[4]int]int {
	visited := map[[4]int]int{} // Visited points to remember the parity
	q := [][4]int{}             // Queue for BFS

	// Set up start point
	q = append(q, s)
	visited[s] = 0

	for i := 1; i < goal+1; i++ { // Run for limited number of steps
		nq := [][4]int{}        // New queue
		for _, cur := range q { // All in queue have same parity
			steps := neighbours(cur, m, visited) // Get neighbours to visit

			// Visit neighbours and record parity
			for _, s := range steps {
				visited[s] = i
				nq = append(nq, s)
			}
		}
		q = nq // New queue
	}
	return visited
}

// BFS modified with oscilating nodes, to run with replicated planes
func mbfs(m []string, s [4]int) map[[4]int]int {
	visited := map[[4]int]int{} // Visited points to remember the parity
	q := [][4]int{}             // Queue for BFS

	// Set up start point
	q = append(q, s)
	visited[s] = 0

	for len(q) > 0 { // Run standard BFS queue loop
		// Pop queue
		cur := q[0]
		q = q[1:]

		steps := neighbours(cur, m, visited) // Get neighbours to visit
		// Visit neighbours and record parity
		for _, s := range steps {
			visited[s] = visited[cur] + 1
			q = append(q, s)
		}

	}
	return visited
}

// Get Neighbours
func neighbours(coords [4]int, m []string, visited map[[4]int]int) (res [][4]int) {
	for _, step := range moves { // Check 4 directions
		nc := [4]int{coords[0] + step[0], coords[1] + step[1], coords[2], coords[3]} // Next possible move

		// Check rollovers on x axis and move to next plane
		if nc[0] < 0 {
			nc[0] = len(m) + nc[0]
			nc[2]--
		} else if nc[0] >= len(m) {
			nc[0] = nc[0] - len(m)
			nc[2]++
		}

		// Check rollovers on y axis and move to next plane
		if nc[1] < 0 {
			nc[1] = len(m[nc[0]]) + nc[1]
			nc[3]--
		} else if nc[1] >= len(m[nc[0]]) {
			nc[1] = nc[1] - len(m[nc[0]])
			nc[3]++
		}

		// Limit number of planes to visit. Range value is a result of guess work
		if (nc[2] > 3 || nc[2] < -3) || (nc[3] > 3 || nc[3] < -3) {
			continue
		}

		// If neighbor can be next step, record it
		if _, ok := visited[nc]; !ok && m[nc[0]][nc[1]] != '#' {
			res = append(res, nc)
		}
	}
	return
}

//go:embed test_data.txt
var TEST_INPUT string
