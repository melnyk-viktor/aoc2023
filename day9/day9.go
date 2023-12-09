package day9

import (
	"sort"
	"strconv"
	"strings"
)

func Solution(input string) (solution1, solution2 int) {
	// Solution 1 & 2
	histories := strings.Split(input, "\n")

	for _, h := range histories { // For every history input
		var his [][]int

		hfs := strings.Fields(h)
		his = append(his, []int{})
		for _, hf := range hfs {
			num, _ := strconv.Atoi(hf)
			his[0] = append(his[0], num)
		}

		i := 0
		for { // Until all zero diffs
			his = append(his, generate_diff(his[i]))

			// Check if non-zero diffs exist
			nis := generate_diff(his[i])
			sort.Ints(nis)
			if nis[len(nis)-1] == 0 && nis[0] == 0 {
				break
			} else {
				i++
			}
		}

		var (
			cur_f int
			cur_p int
		)
		for i = len(his) - 2; i > -1; i-- { // Reverse the diffs
			cur_f = his[i][len(his[i])-1] + cur_f // Move future prediction up
			cur_p = his[i][0] - cur_p             // Move past prediction up
		}

		solution1 += cur_f
		solution2 += cur_p
	}

	return
}

func generate_diff(seq []int) []int {
	var res []int
	for i := 0; i < len(seq)-1; i++ {
		res = append(res, seq[i+1]-seq[i])
	}
	return res
}
