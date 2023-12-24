package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"

	"aoc2023/aoc_utils"
	"aoc2023/day1"
	"aoc2023/day10"
	"aoc2023/day11"
	"aoc2023/day12"
	"aoc2023/day13"
	"aoc2023/day14"
	"aoc2023/day15"
	"aoc2023/day16"
	"aoc2023/day17"
	"aoc2023/day18"
	"aoc2023/day19"
	"aoc2023/day2"
	"aoc2023/day20"
	"aoc2023/day21"
	"aoc2023/day22"
	"aoc2023/day23"
	"aoc2023/day3"
	"aoc2023/day4"
	"aoc2023/day5"
	"aoc2023/day6"
	"aoc2023/day7"
	"aoc2023/day8"
	"aoc2023/day9"
)

// Solutions mapping
var day_mapping = []func(string) (int, int){
	day1.Solution,
	day2.Solution,
	day3.Solution,
	day4.Solution,
	day5.Solution,
	day6.Solution,
	day7.Solution,
	day8.Solution,
	day9.Solution,
	day10.Solution,
	day11.Solution,
	day12.Solution,
	day13.Solution,
	day14.Solution,
	day15.Solution,
	day16.Solution,
	day17.Solution,
	day18.Solution,
	day19.Solution,
	day20.Solution,
	day21.Solution,
	day22.Solution,
	day23.Solution,
}

var URL = "https://adventofcode.com/2023"

func main() {
	var (
		d2d   []string
		se    string
		input string
	)

	// Decoration
	fmt.Println(ART[rand.Intn(len(ART))])

	// NOTE: --input and --inputs can be added, but passing list of files for --days and --all flags is too much for simple cli, that is why input is hardcoded.
	// Also it is not elegant to just restrict input files to some folder and/or naming scheme to parse them automatically, but that is also an option.

	session := flag.String("session", "", "Session cookie for AoC website, can be set via AOC_SESSION env variable")
	day_option := flag.String("day", "", "Run solution for given day")
	days_option := flag.String("days", "", "Comma-separated list of days to run solutions for")
	all_flag := flag.Bool("all", false, "Run all solutions")
	flag.Parse()

	if *session != "" {
		se = *session
	} else {
		se = os.Getenv("AOC_SESSION")
	}

	if se == "" {
		fmt.Printf("\nWARNING: Rinning on predefined tests inputs.\n\n")
	}

	// Check exclusivity of flags
	if (*day_option != "" && *days_option != "") || (*day_option != "" && *all_flag) || (*all_flag && *days_option != "") {
		fmt.Println("Flags are exclusive")
		os.Exit(1)
	}

	// Build
	if *day_option != "" { // NOTE: additional check possible for duplicates
		d2d = append(d2d, *day_option)
	}
	if *days_option != "" { // NOTE: additional check possible for duplicates
		d2d = append(d2d, strings.Split(*days_option, ",")...)
	}

	// Message on no given flags
	if len(d2d) == 0 && !*all_flag {
		fmt.Println("To display solutions use --day=N, --days=X,Y,Z, or --all")
	}

	// Solutions
	// NOTE: input athering can be optimized using goroutines
	for n_day := 0; n_day < len(day_mapping); n_day++ {
		if slices.Contains(d2d, strconv.Itoa(n_day+1)) || *all_flag {
			if se != "" {
				input = aoc_utils.GetInputData(URL, strconv.Itoa(n_day+1), se)
			}
			s1, s2 := day_mapping[n_day](input)
			fmt.Println("AoC 2023 Day", n_day+1)
			fmt.Printf("\tSolution 1: %d\n\tSolution 2: %d\n\n", s1, s2)
		}
	}
}

var ART = []string{
	`
	    _\/_
	     /\
	     /\
	    /  \
	    /~~\o
	   /o   \
	  /~~*~~~\
	 o/    o \
	 /~~~~~~~~\~'
	/__*_______\
	     ||
	   \====/
	    \__/
	`,
	`
	   *        *        *        __o    *       *
	*      *       *        *    /_| _     *
	   K  *     K      *        O'_)/ \  *    *
	  <')____  <')____    __*   V   \  ) __  *
	   \ ___ )--\ ___ )--( (    (___|__)/ /*     *
	 *  |   |    |   |  * \ \____| |___/ /  *
		|*  |    |   | aos \____________/       *
	`,
	`
	            .-~~\
	           /     \ _
	           ~x   .-~_)_
	             ~x".-~   ~-.
	         _   ( /         \   _
	         ||   T  o  o     Y  ||
	       ==:l   l   <       !  I;==
	          \\   \  .__/   /  //
	           \\ ,r"-,___.-'r.//
	            }^ \.( )   _.'//.
	           /    }~Xi--~  //  \
	          Y    Y I\ \    "    Y
	          |    | |o\ \        |
	          |    l_l  Y T       |  -Row
	          l      "o l_j       !
	           \                 /
	    ___,.---^.     o       .^---.._____
	"~~~          "           ~            ~~~"
	`,
}
