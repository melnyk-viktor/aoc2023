package day19

import (
	_ "embed"
	"encoding/json"
	"strconv"
	"strings"
)

type rule struct {
	key      string
	cmpFunc  func(int) bool
	dest     string
	operator string
	i_val    int
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}

	// Solution 1
	var (
		s_workflows []string          // Separate workflow strings
		s_parts     []string          // Separate part descriprion strings
		workflows   map[string][]rule // Map of workflow names to their rulesets
		res         string            // Current resulting workflow or choice A or R
	)

	in_blocks := strings.Split(input, "\n\n") // Split into workflows and parts

	s_workflows = strings.Split(in_blocks[0], "\n")
	s_parts = strings.Split(in_blocks[1], "\n")

	workflows = make(map[string][]rule)
	for _, wf := range s_workflows {
		// Extract rule definitions of workflow
		wfbs := strings.Split(wf, "{")
		key := wfbs[0]
		s_rules := wfbs[1][:len(wfbs[1])-1]

		// Map workflow name to its rule objects
		workflows[key] = []rule{}
		for _, s_rule := range strings.Split(s_rules, ",") {
			workflows[key] = append(workflows[key], parseRule(s_rule))
		}
	}

	for _, s_part := range s_parts {
		// Extract/morph part definition compatible with JSON format
		s_part = strings.ReplaceAll(s_part, "{", "{\"")
		s_part = strings.ReplaceAll(s_part, "=", "\"=")
		s_part = strings.ReplaceAll(s_part, ",", ",\"")
		s_part = strings.ReplaceAll(s_part, "=", ":")

		// Part definition to map conversion
		part := map[string]int{}
		if err := json.Unmarshal([]byte(s_part), &part); err != nil {
			panic(err)
		}

		res = "in" // Starting value
		for res != "A" && res != "R" {
			// Check part against workflow rules and move to next workflow or A/R
			for _, rule := range workflows[res] {
				if rule.cmpFunc(part[rule.key]) {
					res = rule.dest
					break
				}
			}
		}

		// If part is acceptable calculate sum of part values and add to result
		if res == "A" {
			solution1 += part["x"] + part["m"] + part["a"] + part["s"]
		}
	}

	// Solution 2

	// Default part values ranges
	ranges := map[string][2]int{
		"x": {1, 4001},
		"m": {1, 4001},
		"a": {1, 4001},
		"s": {1, 4001},
	}
	solution2 = recurCombinations(ranges, workflows, "in")

	return
}

func parseRule(s_rule string) rule {
	var (
		cf   func(int) bool // Comparison function of the rule
		cbs  []string       // Rule key for value to check
		cv   int            // Value used in comparison function
		dest string         // Destination workflow name or A/R on positive comparison
		op   string         // Operation type
		rbs  []string       // Temporary holder for rule definition blocks
	)
	if strings.Contains(s_rule, ":") {
		rbs = strings.Split(s_rule, ":")
		dest = rbs[1]
		op = ""

		if strings.Contains(rbs[0], "<") {
			op = "<"
			cbs = strings.Split(rbs[0], "<")
			cv, _ = strconv.Atoi(cbs[1])
			cf = func(i int) bool { return i < cv }

		} else if strings.Contains(rbs[0], ">") {
			op = ">"
			cbs = strings.Split(rbs[0], ">")
			cv, _ = strconv.Atoi(cbs[1])
			cf = func(i int) bool { return i > cv }
		}

		return rule{
			key:      cbs[0],
			cmpFunc:  cf,
			dest:     dest,
			operator: op,
			i_val:    cv,
		}

	} else {
		return rule{
			key:      "x",
			cmpFunc:  func(_ int) bool { return true },
			dest:     s_rule,
			operator: "",
			i_val:    0,
		}
	}
}

func recurCombinations(ranges map[string][2]int, wfs map[string][]rule, rkey string) int {
	var (
		res int
		ir  [2]int // Inner range
		or  [2]int // Outer range
	)

	if rkey == "A" { // Compute and return result for accepted ranges
		res = 1
		for _, v := range ranges {
			res *= v[1] - v[0]
		}
		return res
	} else if rkey == "R" { // Dismiss regected
		return 0
	}

	// Deepcopy
	ranges_c := make(map[string][2]int)
	for k, v := range ranges {
		ranges_c[k] = v
	}

	// Adjust ranges of workflow based on comparison operation they use
	for _, r := range wfs[rkey][:len(wfs[rkey])-1] {
		if r.operator == "<" {
			ir = [2]int{ranges_c[r.key][0], r.i_val}
			or = [2]int{r.i_val, ranges_c[r.key][1]}
		} else {
			ir = [2]int{r.i_val + 1, ranges_c[r.key][1]}
			or = [2]int{ranges_c[r.key][0], r.i_val + 1}
		}

		ranges_c[r.key] = ir                            // Set inner range and compute next destination
		res += recurCombinations(ranges_c, wfs, r.dest) // Compute next destination
		ranges_c[r.key] = or                            // Return original range value
	}
	res += recurCombinations(ranges_c, wfs, wfs[rkey][len(wfs[rkey])-1].dest) // Account for last non A/R destination in workflow

	return res
}

//go:embed test_data.txt
var TEST_INPUT string
