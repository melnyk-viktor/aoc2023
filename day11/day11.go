package day11

import (
	"math"
	"slices"
	"strings"
)

var P2EXP = 999999 // 1000000 - 1 accounting for already existing row

func Solution(input string) (solution1, solution2 int) {
	// Solution 1
	var (
		gl        [][2]int // Galaxies
		grs       []int    // Rows with galaxies
		gcs       []int    // Columns with galaxies
		grow_rows []int    // Rows that will expand
		grow_cols []int    // Columns that will expand
	)
	gm := strings.Split(input, "\n")

	// Get coords of all galaxies and mark their rows & columns
	for x := range gm {
		for y := range gm[x] {
			if string(gm[x][y]) == "#" {
				gl = append(gl, [2]int{x, y})
				grs = append(grs, x)
				gcs = append(gcs, y)
			}
		}
	}

	for x := range gm { // Get all rows that will expand
		if !slices.Contains(grs, x) {
			grow_rows = append(grow_rows, x)
		}
	}
	for y := range gm[0] { // Get all columns that will expand
		if !slices.Contains(gcs, y) {
			grow_cols = append(grow_cols, y)
		}
	}

	// Calculate new coordinates for galaxies, accounting for expansion
	for i, gc := range gl {
		count := 0
		for j := 0; j < len(grow_rows); j++ {
			if grow_rows[j] < gc[0] { // Only "prior" rows that expand will change coordinate of galaxy
				count++
			} else {
				break
			}
		}
		gl[i][0] += count // Shift galaxy on x

		count = 0
		for j := 0; j < len(grow_cols); j++ {
			if grow_cols[j] < gc[1] { // Only "prior" columns that expand will change coordinate of galaxy
				count++
			} else {
				break
			}
		}
		gl[i][1] += count // Shift galaxy on y
	}

	for i := 0; i < len(gl)-1; i++ { // Last galaxy has no pairings
		for j := i + 1; j < len(gl); j++ { // Pair every next galaxies
			/*
				x1 - x2 is distance by x coordinate
				y1 - y2 is distance by y coordinate
				Use absolute values, sign is irrelevant, it depends on order
			*/
			solution1 += int(math.Abs(float64(gl[i][0]-gl[j][0])) + math.Abs(float64(gl[i][1]-gl[j][1])))
		}
	}

	// Solution 2
	gl = [][2]int{}     // Galaxies
	grs = []int{}       // Rows with galaxies
	gcs = []int{}       // Columns with galaxies
	grow_rows = []int{} // Rows that will expand
	grow_cols = []int{} // Columns that will expand
	gm = strings.Split(input, "\n")

	// Get coords of all galaxies and mark their rows & columns
	for x := range gm {
		for y := range gm[x] {
			if string(gm[x][y]) == "#" {
				gl = append(gl, [2]int{x, y})
				grs = append(grs, x)
				gcs = append(gcs, y)
			}
		}
	}

	for x := range gm { // Get all rows that will expand
		if !slices.Contains(grs, x) {
			grow_rows = append(grow_rows, x)
		}
	}
	for y := range gm[0] { // Get all columns that will expand
		if !slices.Contains(gcs, y) {
			grow_cols = append(grow_cols, y)
		}
	}

	// Calculate new coordinates for galaxies, accounting for expansion and P2EXP
	for i, gc := range gl {
		count := 0
		for j := 0; j < len(grow_rows); j++ {
			if grow_rows[j] < gc[0] { // Only "prior" rows that expand will change coordinate of galaxy
				count++
			} else {
				break
			}
		}
		gl[i][0] += count * P2EXP

		count = 0
		for j := 0; j < len(grow_cols); j++ {
			if grow_cols[j] < gc[1] { // Only "prior" columns that expand will change coordinate of galaxy
				count++
			} else {
				break
			}
		}
		gl[i][1] += count * P2EXP
	}

	for i := 0; i < len(gl)-1; i++ { // Last galaxy has no pairings
		for j := i + 1; j < len(gl); j++ {
			/*
				x1 - x2 is distance by x coordinate
				y1 - y2 is distance by y coordinate
				Use absolute values, sign is irrelevant, it depends on order
			*/
			solution2 += int(math.Abs(float64(gl[i][0]-gl[j][0])) + math.Abs(float64(gl[i][1]-gl[j][1])))
		}
	}

	return
}
