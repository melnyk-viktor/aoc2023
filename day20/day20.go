package day20

import (
	_ "embed"
	"slices"
	"strings"
)

type module struct {
	id       string
	sid      string
	mType    string
	connects []string
	isFlip   bool
	flip     bool
	sig      int
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}

	// Solution 1 & 2
	var (
		modules           map[string]module         // Map of modules
		signalPropagation map[string]map[string]int // Signal progression map
		rxs               module                    // Source rx (parent module of rx)
		lps               int                       // Low signals count
		hps               int                       // High signals count
		p                 int                       // Button presses count
		rxsqs             map[string]int            // Button presses for rx
	)

	modules = map[string]module{}
	signalPropagation = map[string]map[string]int{}

	// Parse modules
	for _, l := range strings.Split(input, "\n") {
		sModule := strings.Split(l, " -> ")
		m := module{
			id:       sModule[0][1:],
			mType:    string(sModule[0][0]),
			connects: strings.Split(sModule[1], ", "),
			isFlip:   string(sModule[0][0]) == "%",
		}

		if sModule[0] == "broadcaster" {
			m.id = sModule[0]
		}
		if slices.Contains(m.connects, "rx") {
			rxs = m
		}

		modules[m.id] = m // Add module to mapping

		if m.mType == "&" {
			signalPropagation[m.id] = map[string]int{} // Add mapping for signal propagation of & module
		}
	}

	for k, m := range modules {
		for _, v := range m.connects {
			if _, ok := signalPropagation[v]; ok {
				signalPropagation[v][k] = 0 // Set default value for all combinations of & modules
			}
		}
	}

	rxsqs = map[string]int{}
	for i := 0; i < 1000; i++ {
		hps, lps, p, rxsqs = bp(hps, lps, p, modules, signalPropagation, rxs, rxsqs) // Push button
	}
	solution1 = hps * lps // Get solution of all low and all high signals

	// Continue for part 2 without zeroing values
	for len(rxsqs) < 4 {
		hps, lps, p, rxsqs = bp(hps, lps, p, modules, signalPropagation, rxs, rxsqs) // Push button
	}

	// Convert map to slice of ints
	rRssqs := []int{}
	for _, v := range rxsqs {
		rRssqs = append(rRssqs, v)
	}
	solution2 = LCM(rRssqs...) // Calculate solution by LCM

	return
}

func bp(
	hps, lps, p int,
	ms map[string]module,
	sp map[string]map[string]int,
	rxs module,
	rxsgs map[string]int,
) (int, int, int, map[string]int) {
	p += 1                                     // Increase push count
	lps += 1 + len(ms["broadcaster"].connects) // All pushes from broadcaster set to low
	q := []module{}                            // Slice used as queue

	// Start queue with broadcast point (start)
	for _, c := range ms["broadcaster"].connects {
		m := ms[c]
		m.sid = "broadcaster"
		m.sig = 0
		q = append(q, m)
	}

	for len(q) > 0 {
		// Pop 1st module
		m := q[0]
		q = q[1:]

		if m.id == "rx" {
			continue
		}

		propagate := 0 // Propagation of signal value

		// Check if propagate flips to 1
		if _, ok := sp[m.id]; ok {
			sp[m.id][m.sid] = m.sig
			for _, v := range sp[m.id] {
				if v == 0 {
					propagate = 1
					break
				}
			}
		}

		// Check if propagate flips to 1
		if m.isFlip {
			if m.sig == 1 {
				continue
			} else {
				tm := ms[m.id]
				tm.flip = !tm.flip
				ms[m.id] = tm
				if tm.flip {
					propagate = 1
				}
			}
		}

		// Set how many connections are propagated with which value
		if propagate == 1 {
			hps += len(m.connects)
		} else {
			lps += len(m.connects)
		}

		// Add further connections to queue
		for _, c := range m.connects {
			tm := ms[c]
			tm.sid = m.id
			tm.sig = propagate
			q = append(q, tm)
		}

		// Check if any rx connections conjuncting 1
		for k, v := range sp[rxs.id] {
			if _, ok := rxsgs[k]; v == 1 && !ok {
				rxsgs[k] = p
			}
		}

	}

	return hps, lps, p, rxsgs
}

// LCM implementation adjusted from https://go.dev/play/p/SmzvkDjYlb

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	if len(integers) == 0 {
		return 0
	} else if len(integers) == 1 {
		return integers[0]
	}

	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

//go:embed test_data.txt
var TEST_INPUT string
