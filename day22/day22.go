package day22

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

type brick struct {
	mapping    string // How brick is represented in initial mapping
	x1, y1, z1 int    // Start voxel coords
	x2, y2, z2 int    // End voxel coords
	// supports    []*brick
	// supportedBy []*brick
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}
	sBricks := strings.Split(input, "\n")

	// Solution 1, 2
	var (
		bricks []brick // Parsed brick structure -> Ordered by z position
		ns     []brick // New structure of bricks, used to check reaction to missing bricks
	)

	bricks = []brick{}
	for _, sBrick := range sBricks {
		sEnds := strings.Split(sBrick, "~") // Get brick start and brick end

		// Convert to ints
		bs := [3]int{}
		for i, c := range strings.Split(sEnds[0], ",") {
			ic, _ := strconv.Atoi(c)
			bs[i] = ic
		}

		// Convert to ints
		be := [3]int{}
		for i, c := range strings.Split(sEnds[1], ",") {
			ic, _ := strconv.Atoi(c)
			be[i] = ic
		}

		// Init brick
		b := brick{
			mapping: sBrick,
			x1:      bs[0],
			y1:      bs[1],
			z1:      bs[2],
			x2:      be[0],
			y2:      be[1],
			z2:      be[2],
		}

		bricks = append(bricks, b) // Append to unsorted bricks
	}

	// Sort bricks by pre-fall z
	slices.SortFunc(bricks, func(a, b brick) int {
		return a.z1 - b.z1
	})

	getFall(bricks) // Get stable brick structure

	ns = make([]brick, len(bricks)) // Init new bricks structure
	for i := range bricks {
		// Copy structure with missing brick i
		for j, b := range bricks {
			if j == i {
				ns[j] = brick{}
			} else {
				ns[j] = b
			}
		}

		movedBricks := getFall(ns) // Check the consequences (number of fallen bricks) due to removed brick
		if movedBricks == 0 {
			solution1 += 1
		} else {
			solution2 += movedBricks
		}
	}

	return
}

func getFall(bricks []brick) (res int) {
outer: // Loop marker
	for i := range bricks {
		cb := &bricks[i]
		fc := false // Check if lowest z possible
		for cb.z1 > 1 {
			for j := i - 1; j >= 0; j-- { // Check for colliding bricks from z & below
				b := &bricks[j]
				// Check if bricks collide
				if (cb.z2-1) >= b.z1 && (cb.z1-1) <= b.z2 && cb.x2 >= b.x1 && cb.x1 <= b.x2 && cb.y2 >= b.y1 && cb.y1 <= b.y2 {
					continue outer
				}
			}

			if !fc { // Brick reached steady z
				res += 1 // Count brick as moved
				fc = true
			}

			// Move down
			cb.z1 -= 1
			cb.z2 -= 1
		}
	}
	return
}

//go:embed test_data.txt
var TEST_INPUT string
