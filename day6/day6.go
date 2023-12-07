package day6

import (
	"math"
	"strconv"
	"strings"
)

func Solution(input string) (solution1, solution2 int) {
	// Solution 1
	s_rt := strings.Fields(strings.Split(input, "\n")[0])[1:]
	s_rd := strings.Fields(strings.Split(input, "\n")[1])[1:]

	solution1 = 1

	// Left here for honesty XD
	// for i := 0; i < len(s_rd); i++ {
	// 	rt, _ := strconv.Atoi(s_rt[i])
	// 	rd, _ := strconv.Atoi(s_rd[i])

	// 	var better_races int

	// 	for t := 1; t < rt; t++{
	// 		if t * (rt - t) > rd {
	// 			better_races++
	// 		}
	// 	}

	// 	solution1 *= better_races
	// }

	// Better solution I had no brain capacity to analyze in first 15 min
	// (rd = (rt - x) * x) equasion for longer distance than rd, if x is a period to hold the button

	for i := 0; i < len(s_rt); i++ {
		rt_i, _ := strconv.Atoi(s_rt[i])
		rt := float64(rt_i)
		rd_i, _ := strconv.Atoi(s_rd[i])
		rd := float64(rd_i)

		dsc := rt*rt - 4*(-1)*(-rd)
		x1 := ((-rt) + math.Sqrt(dsc)) / (2 * -1)
		x2 := ((-rt) - math.Sqrt(dsc)) / (2 * (-1))

		solution1 *= int(math.Floor(x2)) - int(math.Ceil(x1)) + 1
	}

	// Solution 2
	rt_s := strings.Split(strings.ReplaceAll(strings.Split(input, "\n")[0], " ", ""), ":")[1]
	rd_s := strings.Split(strings.ReplaceAll(strings.Split(input, "\n")[1], " ", ""), ":")[1]

	// Left here for honesty XD
	// solution2 = 0
	// rt, _ := strconv.Atoi(rt_s)
	// rd, _ := strconv.Atoi(rd_s)

	// fmt.Println(rt, rd)

	// for t := 1; t < rt; t++{
	// 	if t * (rt - t) > rd {
	// 		solution2++
	// 	}
	// }

	// Better solution I had no brain capacity to analyze in first 15 min
	// (rd = (rt - x) * x) equasion for longer distance than rd, if x is a period to hold the button
	rt_i, _ := strconv.Atoi(rt_s)
	rt := float64(rt_i)
	rd_i, _ := strconv.Atoi(rd_s)
	rd := float64(rd_i)

	dsc := rt*rt - 4*(-1)*(-rd)
	x1 := ((-rt) + math.Sqrt(dsc)) / (2 * -1)
	x2 := ((-rt) - math.Sqrt(dsc)) / (2 * (-1))

	solution2 = int(math.Floor(x2)) - int(math.Ceil(x1)) + 1

	return
}
