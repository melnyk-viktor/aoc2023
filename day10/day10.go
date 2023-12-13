package day10

import (
	_ "embed"
	"slices"
	"strings"
)

type node struct {
	coords      [2]int
	ptype       string
	connections []node
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}

	// Solution 1 & 2
	pipe_map := strings.Split(input, "\n")

	var sn node
	var final_path [][2]int
	for i, l := range pipe_map { // Find and define start
		if strings.Contains(l, "S") {
			sn = node{ // Create start node
				coords: [2]int{i, strings.Index(l, "S")},
				ptype:  string(l[strings.Index(l, "S")]),
			}
			// Check connection on top of start
			if sn.coords[0] > 0 && slices.Contains([]string{"|", "7", "F"}, string(pipe_map[sn.coords[0]-1][sn.coords[1]])) {
				sn.connections = append(sn.connections, node{coords: [2]int{sn.coords[0] - 1, sn.coords[1]}, ptype: string(pipe_map[sn.coords[0]-1][sn.coords[1]])})
			}
			// Check connection under the start
			if sn.coords[0] < len(pipe_map) && slices.Contains([]string{"|", "J", "L"}, string(pipe_map[sn.coords[0]+1][sn.coords[1]])) {
				sn.connections = append(sn.connections, node{coords: [2]int{sn.coords[0] + 1, sn.coords[1]}, ptype: string(pipe_map[sn.coords[0]+1][sn.coords[1]])})
			}
			// Check connection to the left of start
			if sn.coords[1] > 0 && slices.Contains([]string{"-", "F", "L"}, string(pipe_map[sn.coords[0]][sn.coords[1]-1])) {
				sn.connections = append(sn.connections, node{coords: [2]int{sn.coords[0], sn.coords[1] - 1}, ptype: string(pipe_map[sn.coords[0]][sn.coords[1]-1])})
			}
			// Check connection to the right of start
			if sn.coords[1] < len(pipe_map[sn.coords[0]]) && slices.Contains([]string{"-", "J", "7"}, string(pipe_map[sn.coords[0]][sn.coords[1]+1])) {
				sn.connections = append(sn.connections, node{coords: [2]int{sn.coords[0], sn.coords[1] + 1}, ptype: string(pipe_map[sn.coords[0]][sn.coords[1]+1])})
			}
			break
		}
	}

	for i := 0; i < len(sn.connections); i++ { // Check every connection of start
		end_conn, res_l, path := get_connected_nodes(sn.connections[i], &sn, pipe_map) // Recurse to get path

		if end_conn.coords != [2]int{-1, -1} { // Check for dangling path
			for i, conn := range sn.connections {
				if conn.coords == end_conn.coords {
					// Remove strart connection that was the end of the other loop
					sn.connections = append(sn.connections[:i], sn.connections[i+1:]...)
				}
			}
			if res_l/2 > solution1 {
				solution1 = res_l / 2                               // Furthest point is in the middle of the path
				final_path = append(path, sn.connections[i].coords) // Append last connection to path (needed for part 2)
			}
		}
	}

	// Solution 2
	for x, line := range pipe_map { // Scanning every line
		crossed_first := false // Checker if crossed first part of our path
		vls := 0               // Number of vertical connectors from our path in line
		for y := range line {
			if slices.Contains(final_path, [2]int{x, y}) {
				crossed_first = true // Set checker of crossion the first part of our path in line

				// If encountered vertical connector count
				if string(line[y]) == "|" || string(line[y]) == "J" || string(line[y]) == "L" || string(line[y]) == "S" {
					vls++
				}
			} else if crossed_first && vls%2 == 1 { // Tile in between paired vertical connectors is "inside"
				solution2++
			}
		}
	}

	return
}

// Rec to get path of connected nodes
func get_connected_nodes(c node, from *node, pipe_map []string) (*node, int, [][2]int) {
	switch c.ptype {
	case "|": // Vertical connection
		if from.coords[0] > c.coords[0] { // Going down
			step := node{
				coords: [2]int{c.coords[0] - 1, c.coords[1]},
				ptype:  string(pipe_map[c.coords[0]-1][c.coords[1]]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		} else { // Going up
			step := node{
				coords: [2]int{c.coords[0] + 1, c.coords[1]},
				ptype:  string(pipe_map[c.coords[0]+1][c.coords[1]]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		}
	case "-": // Horizontal connection
		if from.coords[1] > c.coords[1] { // Right to left
			step := node{
				coords: [2]int{c.coords[0], c.coords[1] - 1},
				ptype:  string(pipe_map[c.coords[0]][c.coords[1]-1]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		} else { // Left to right
			step := node{
				coords: [2]int{c.coords[0], c.coords[1] + 1},
				ptype:  string(pipe_map[c.coords[0]][c.coords[1]+1]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		}
	case "L": // Top to right connection
		if from.coords[0] < c.coords[0] { // Top to right
			step := node{
				coords: [2]int{c.coords[0], c.coords[1] + 1},
				ptype:  string(pipe_map[c.coords[0]][c.coords[1]+1]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		} else { // Right to up
			step := node{
				coords: [2]int{c.coords[0] - 1, c.coords[1]},
				ptype:  string(pipe_map[c.coords[0]-1][c.coords[1]]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		}
	case "J": // Top to left connection
		if from.coords[0] < c.coords[0] { // Top to left
			step := node{
				coords: [2]int{c.coords[0], c.coords[1] - 1},
				ptype:  string(pipe_map[c.coords[0]][c.coords[1]-1]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		} else { // Left to up
			step := node{
				coords: [2]int{c.coords[0] - 1, c.coords[1]},
				ptype:  string(pipe_map[c.coords[0]-1][c.coords[1]]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		}
	case "7": // Left to bottom connection
		if from.coords[1] < c.coords[1] { // Left to down
			step := node{
				coords: [2]int{c.coords[0] + 1, c.coords[1]},
				ptype:  string(pipe_map[c.coords[0]+1][c.coords[1]]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		} else { // Down to left
			step := node{
				coords: [2]int{c.coords[0], c.coords[1] - 1},
				ptype:  string(pipe_map[c.coords[0]][c.coords[1]-1]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		}
	case "F": // Right to bottom connection
		if from.coords[1] > c.coords[1] { // Right to down
			step := node{
				coords: [2]int{c.coords[0] + 1, c.coords[1]},
				ptype:  string(pipe_map[c.coords[0]+1][c.coords[1]]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_p = append(res_p, step.coords)
				res_l++
			}
			return res_c, res_l, res_p
		} else { // Down to right
			step := node{
				coords: [2]int{c.coords[0], c.coords[1] + 1},
				ptype:  string(pipe_map[c.coords[0]][c.coords[1]+1]),
			}
			if from.ptype != "S" {
				from.connections = append(from.connections, step) // Register connection in node
			}
			res_c, res_l, res_p := get_connected_nodes(step, &c, pipe_map) // Rec further
			if res_l != -1 {
				res_l++
			}
			res_p = append(res_p, step.coords)
			return res_c, res_l, res_p
		}
	case "S": // Encountered Start/End Point
		return from, 1, [][2]int{}
	default: // Encountered Dangling path
		return &node{coords: [2]int{-1, -1}}, -1, [][2]int{{-1, -1}}
	}
}

//go:embed test_data.txt
var TEST_INPUT string
