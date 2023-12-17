package day17

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"
)

type point struct {
	coords [2]int // (x, y)
	dir    [2]int // (dx, dy)
	cost   int
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}
	grid := strings.Split(input, "\n")

	var (
		hp    *pointHeap     // Min heap based on cost
		ec    [2]int         // Exit coordinates (x, y)
		dists map[string]int // Map of visited points with directions of visit, to cost of visit
	)

	// Solution 1
	hp = &pointHeap{
		{coords: [2]int{0, 0}, dir: [2]int{0, 1}, cost: 0}, // Starting at (0, 0), going right
		{coords: [2]int{0, 0}, dir: [2]int{1, 0}, cost: 0}, // Starting at (0, 0), going down
	}
	heap.Init(hp)

	ec = [2]int{len(grid) - 1, len(grid[0]) - 1}
	dists = make(map[string]int)

	for hp.Len() > 0 {
		cur := heap.Pop(hp).(point) // Pop current point

		if cur.coords[0] == ec[0] && cur.coords[1] == ec[1] {
			solution1 = cur.cost
			goto endS1 // Stop iterating on 1st encounter of end point
		}

		if v, ok := dists[fmt.Sprint(cur.coords, cur.dir)]; ok && cur.cost > v { // Point already encountered with better result
			continue
		}

		for _, i := range [2][2]int{{-cur.dir[1], cur.dir[0]}, {cur.dir[1], -cur.dir[0]}} { // Swap directions for turn/split
			nc := cur.cost
			for m := 1; m < 4; m++ { // Can't move more than 3 points in same direction
				nx := cur.coords[0] + (i[0] * m) // Calculate new x
				ny := cur.coords[1] + (i[1] * m) // Calculate new y

				if (nx >= 0 && nx < len(grid)) && (ny >= 0 && ny < len(grid[nx])) {
					nc += int(grid[nx][ny] - '0') // Increase cost of moving to new point on grid
					n := point{coords: [2]int{nx, ny}, dir: [2]int{i[0], i[1]}, cost: nc}
					if v, ok := dists[fmt.Sprint(n.coords, n.dir)]; !ok || nc < v { // If not encountered before or if cost is more optimal
						dists[fmt.Sprint(n.coords, n.dir)] = nc // Register new cost in map
						heap.Push(hp, n)                        // Push to min heap
					}
				}
			}
		}
	}
endS1:

	// Solution 2
	hp = &pointHeap{
		{coords: [2]int{0, 0}, dir: [2]int{0, 1}, cost: 0}, // Starting at (0, 0), going right
		{coords: [2]int{0, 0}, dir: [2]int{1, 0}, cost: 0}, // Starting at (0, 0), going down
	}
	heap.Init(hp)

	ec = [2]int{len(grid) - 1, len(grid[0]) - 1}
	dists = make(map[string]int)

	for hp.Len() > 0 {
		cur := heap.Pop(hp).(point) // Pop current point

		if cur.coords[0] == ec[0] && cur.coords[1] == ec[1] {
			solution2 = cur.cost
			goto endS2 // Stop iteration on 1st encounter of endpoint
		}

		if v, ok := dists[fmt.Sprint(cur.coords, cur.dir)]; ok && cur.cost > v { // Point already encountered with better result
			continue
		}

		for _, i := range [2][2]int{{-cur.dir[1], cur.dir[0]}, {cur.dir[1], -cur.dir[0]}} { // Swapped directions for turn/split
			nc := cur.cost
			for m := 1; m < 11; m++ { // Can't move more than 10 points in same direction
				nx := cur.coords[0] + (i[0] * m) // Calculate new x
				ny := cur.coords[1] + (i[1] * m) // Calculate new y
				if (nx >= 0 && nx < len(grid)) && (ny >= 0 && ny < len(grid[nx])) {
					nc += int(grid[nx][ny] - '0') // Increase cost of moving to new point on grid
					n := point{coords: [2]int{nx, ny}, dir: [2]int{i[0], i[1]}, cost: nc}
					if v, ok := dists[fmt.Sprint(n.coords, n.dir)]; (!ok || nc < v) && m >= 4 { // If not encountered before or if cost is more optimal, Has to be at least 4 moves
						dists[fmt.Sprint(n.coords, n.dir)] = nc // Register new cost in map
						heap.Push(hp, n)                        // Push to min heap
					}
				}
			}
		}
	}
endS2:

	return
}

//go:embed test_data.txt
var TEST_INPUT string

// Min heap implementation from official docs: https://pkg.go.dev/container/heap#example-package-IntHeap
// An IntHeap is a min-heap of ints.
type pointHeap []point

func (h pointHeap) Len() int           { return len(h) }
func (h pointHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h pointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pointHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(point))
}

func (h *pointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
