package day16

import (
	_ "embed"
	"slices"
	"strings"
)

type move struct {
	coords [2]int // (x, y)
	dir    rune   // <, >, ^, v
}

// Default mapping for continuing movement
var moveMap = map[rune][2]int{
	'^': {-1, 0}, // Go up on x
	'v': {1, 0},  // Go down on x
	'<': {0, -1}, // Go left on y
	'>': {0, 1},  // Go right on y
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}
	grid := strings.Split(input, "\n")

	// Solution 1
	solution1 = calcPath(grid, [2]int{0, 0}, '>')

	// Solution 2
	resList := []int{}
	for i := range grid { // For every row
		resList = append(resList, calcPath(grid, [2]int{i, 0}, '>'))                // Start on the left
		resList = append(resList, calcPath(grid, [2]int{i, len(grid[i]) - 1}, '<')) // Start on the right
	}
	for i := range grid[0] { // For every column
		resList = append(resList, calcPath(grid, [2]int{0, i}, 'v'))             // Start on top
		resList = append(resList, calcPath(grid, [2]int{len(grid) - 1, i}, '^')) // Start at the bottom
	}
	solution2 = slices.Max(resList) // Get longest path

	return
}

func calcPath(grid []string, start [2]int, dir rune) int {
	mem := map[[2]int][]rune{}             // Map of visited points on a grid
	q := []move{{coords: start, dir: dir}} // Slice used as queue for "next move"

	for len(q) > 0 {
		// Pop the top value
		cur := q[len(q)-1]
		q = q[:len(q)-1]

		// If the point was visited from the direction that already recorded, skip
		if v, ok := mem[cur.coords]; ok {
			if slices.Contains(v, cur.dir) {
				continue
			}
		}

		// Check if not running out of bounds
		if cur.coords[0] < 0 || cur.coords[0] >= len(grid) || cur.coords[1] < 0 || cur.coords[1] >= len(grid[0]) {
			continue
		}

		mem[cur.coords] = append(mem[cur.coords], cur.dir) // Record visit

		// Determine next move
		switch grid[cur.coords[0]][cur.coords[1]] { // Check current point type
		case '.': // Current point is empty
			// Move through point preserving direction
			q = append(q, move{
				coords: [2]int{cur.coords[0] + moveMap[cur.dir][0], cur.coords[1] + moveMap[cur.dir][1]},
				dir:    cur.dir,
			})
		case '/': // Mirror of "/" type
			switch cur.dir {
			case '>':
				// Reflect up
				q = append(q, move{
					coords: [2]int{cur.coords[0] - 1, cur.coords[1]},
					dir:    '^',
				})
			case '<':
				// Reflect down
				q = append(q, move{
					coords: [2]int{cur.coords[0] + 1, cur.coords[1]},
					dir:    'v',
				})
			case 'v':
				// Reflect left
				q = append(q, move{
					coords: [2]int{cur.coords[0], cur.coords[1] - 1},
					dir:    '<',
				})
			case '^':
				// Reflect rught
				q = append(q, move{
					coords: [2]int{cur.coords[0], cur.coords[1] + 1},
					dir:    '>',
				})
			}
		case '\\': // Mirror of type "\"
			switch cur.dir {
			case '>': // Reflect down
				q = append(q, move{
					coords: [2]int{cur.coords[0] + 1, cur.coords[1]},
					dir:    'v',
				})
			case '<': // Reflect up
				q = append(q, move{
					coords: [2]int{cur.coords[0] - 1, cur.coords[1]},
					dir:    '^',
				})
			case 'v': // Reflect right
				q = append(q, move{
					coords: [2]int{cur.coords[0], cur.coords[1] + 1},
					dir:    '>',
				})
			case '^':
				// Reflect left
				q = append(q, move{
					coords: [2]int{cur.coords[0], cur.coords[1] - 1},
					dir:    '<',
				})
			}
		case '|': // Splitter of vertical type
			if cur.dir == '<' || cur.dir == '>' { // Split beam into vertical beams going up & down
				q = append(q, move{
					coords: [2]int{cur.coords[0] - 1, cur.coords[1]},
					dir:    '^',
				})
				q = append(q, move{
					coords: [2]int{cur.coords[0] + 1, cur.coords[1]},
					dir:    'v',
				})
			} else {
				// Move through point preserving direction
				q = append(q, move{
					coords: [2]int{cur.coords[0] + moveMap[cur.dir][0], cur.coords[1] + moveMap[cur.dir][1]},
					dir:    cur.dir,
				})
			}
		case '-': // Splitter of horizontal type
			if cur.dir == '^' || cur.dir == 'v' { // Split beam into horizontal beams goilg left and right
				q = append(q, move{
					coords: [2]int{cur.coords[0], cur.coords[1] - 1},
					dir:    '<',
				})
				q = append(q, move{
					coords: [2]int{cur.coords[0], cur.coords[1] + 1},
					dir:    '>',
				})
			} else {
				// Move through point preserving direction
				q = append(q, move{
					coords: [2]int{cur.coords[0] + moveMap[cur.dir][0], cur.coords[1] + moveMap[cur.dir][1]},
					dir:    cur.dir,
				})
			}
		}
	}

	return len(mem) // Get number of pointa in path
}

//go:embed test_data.txt
var TEST_INPUT string
