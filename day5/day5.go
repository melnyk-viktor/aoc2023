package day5

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func Solution(input string) (solution1, solution2 int) {
	// Solution 1
	s_stp := strings.Split(strings.Split(input, "\n")[0], " ")[1:]

	var (
		stp_in  []int
		stp_out []int
	)

	for _, s := range s_stp {
		i_stp, _ := strconv.Atoi(s)
		stp_in = append(stp_in, i_stp)
	}

	mappings := strings.Split(input, "\n\n")[1:]
	for _, mapping := range mappings { // Maps
		mapping_lines := strings.Split(mapping, "\n")[1:]
		for _, rn := range mapping_lines { // Every move in mapping
			s_vals := strings.Split(rn, " ")
			dest, _ := strconv.Atoi(s_vals[0])
			src, _ := strconv.Atoi(s_vals[1])
			rlen, _ := strconv.Atoi(s_vals[2])

			for i := 0; i < len(stp_in); i++ { // Translate via move if needed
				if stp_in[i] >= src && stp_in[i] < src+rlen {
					stp_out = append(stp_out, dest+(stp_in[i]-src))

					// Handle pop
					stp_in = append(stp_in[:i], stp_in[i+1:]...)
					i--
				}
			}
		}
		stp_in = append(stp_in, stp_out...)
		stp_out = []int{}
	}
	solution1 = slices.Min(stp_in)

	// Solution 2
	s_stp = strings.Split(strings.Split(input, "\n")[0], " ")[1:]

	// Group seed input into intervals, code below can be modified to run on indexing
	var ranges []Interval
	for i := 0; i < len(s_stp); i += 2 {
		start, _ := strconv.Atoi(s_stp[i])
		ilen, _ := strconv.Atoi(s_stp[i+1])
		end := start + ilen - 1
		ranges = append(ranges, Interval{Start: start, End: end})
	}

	// Mapping by ranges
	for _, mapping := range mappings { // Maps
		formed_moves := []Interval{}
		mapping_lines := strings.Split(mapping, "\n")[1:]
		for i := 0; i < len(ranges); i++ { // Every range in every map, simple loop for pop functionality
			rn := ranges[i]
			for _, moves := range mapping_lines { // Combine map moves with range
				s_vals := strings.Split(moves, " ")
				dest, _ := strconv.Atoi(s_vals[0])
				ms, _ := strconv.Atoi(s_vals[1])
				mlen, _ := strconv.Atoi(s_vals[2])

				if rn.Start <= (ms+mlen)-1 && ms <= rn.End {
					// Pre
					if rn.Start < ms {
						ranges = append(ranges, Interval{Start: rn.Start, End: ms - 1})
					}

					// Overlap
					formed_moves = append(
						formed_moves,
						Interval{
							Start: int(math.Max(float64(rn.Start), float64(ms))) - ms + dest,        // start at dest + dist_from_start (dist from start can be 0)
							End:   int(math.Min(float64(rn.End), float64((ms+mlen)-1))) - ms + dest, // end at dest + len_of_interval
						},
					)

					// Post
					if rn.End > (ms+mlen)-1 {
						ranges = append(ranges, Interval{Start: (ms + mlen), End: rn.End})
					}

					// Handle pop
					ranges = append(ranges[:i], ranges[i+1:]...)
					i -= 1
					break
				}
			}
		}
		ranges = append(ranges, formed_moves...) // append new ranges to be processed by other mappings
	}

	solution2 = -1
	for _, r := range ranges {
		if solution2 == -1 || solution2 > r.Start {
			solution2 = r.Start
		}
	}

	return
}
