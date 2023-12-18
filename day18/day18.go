package day18

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

var moves = map[string][2]int{
	"U": {-1, 0},
	"D": {1, 0},
	"L": {0, -1},
	"R": {0, 1},
}

var hMoves = map[rune][2]int{
	'0': {0, 1},
	'1': {1, 0},
	'2': {0, -1},
	'3': {-1, 0},
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}

	plan := strings.Split(input, "\n")
	var (
		x   int
		y   int
		p   int      // Perimeter
		a   int      // Area
		eps [][2]int // Slice of corner points of a polygon
	)

	// Solution 1
	x, y, p, a = 0, 0, 0, 0
	eps = make([][2]int, 0)

	// Get corner coordinaetes of a polygon
	for i, s_instr := range plan {
		instr := strings.Split(s_instr, " ")
		m := moves[instr[0]]
		mlen, _ := strconv.Atoi(instr[1])

		x += mlen * m[0]
		y += mlen * m[1]

		p += mlen
		eps = append(eps, [2]int{x, y})

		if i > 0 {
			a += (eps[i-1][0] * eps[i][1]) - (eps[i-1][1] * eps[i][0])
		}
	}
	a = int(math.Abs(float64(a)) / 2) // Internal rea according to Gauss's formula

	solution1 = a + p/2 + 1 // Account for missing points (Pick's theorem https://en.wikipedia.org/wiki/Pick%27s_theorem)

	// Solution 2
	x, y, p, a = 0, 0, 0, 0
	eps = make([][2]int, 0)

	// Get all corner coordinated of polygon
	for i, s_instr := range plan {
		instr := strings.Split(s_instr, " ")
		m := hMoves[rune(instr[2][len(instr[2])-2])]
		mlen64, _ := strconv.ParseInt(instr[2][2:len(instr[2])-2], 16, 0)
		mlen := int(mlen64)

		x += mlen * m[0]
		y += mlen * m[1]

		p += mlen
		eps = append(eps, [2]int{x, y})

		if i > 0 {
			a += (eps[i-1][0] * eps[i][1]) - (eps[i-1][1] * eps[i][0])
		}
	}
	a = int(math.Abs(float64(a)) / 2) // Internal rea according to Gauss's formula

	solution2 = a + p/2 + 1 // Account for missing points (Pick's theorem https://en.wikipedia.org/wiki/Pick%27s_theorem)

	return
}

//go:embed test_data.txt
var TEST_INPUT string
