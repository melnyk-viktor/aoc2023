package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"aoc2023/day4"
	"aoc2023/day5"
	"aoc2023/day6"
	"flag"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
)

// Solutions mapping
var day_mapping = []func()(int, int){
	day1.Solution,
	day2.Solution,
	day3.Solution,
	day4.Solution,
	day5.Solution,
	day6.Solution,
}

func main() {
	// Decoration
	fmt.Println(ART[rand.Intn(len(ART))])

	// NOTE: --input and --inputs can be added, but passing list of files for --days and --all flags is too much for simple cli, that is why input is hardcoded.
	// Also it is not elegant to just restrict input files to some folder and/or naming scheme to parse them automatically, but that is also an option.

	day_option := flag.String("day", "", "Run solution for given day")
	days_option := flag.String("days", "", "Comma-separated list of days to run solutions for")
	all_flag := flag.Bool("all", false, "Run all solutions")
	flag.Parse()

	// Check exclusivity of flags
	if (*day_option != "" && *days_option != "") || (*day_option != "" && *all_flag) || (*all_flag && *days_option != "") {
		fmt.Println("Flags are exclusive")
		return // NOTE: maybe use os.Exit(1)
	}

	// Build
	var days_to_display []string
	if *day_option != "" { // NOTE: additional check possible for duplicates
		days_to_display = append(days_to_display, *day_option)
	}
	if *days_option != "" { // NOTE: additional check possible for duplicates
		days_to_display = append(days_to_display, strings.Split(*days_option, ",")...)
	}

	// Message on no given flags
	if len(days_to_display) == 0 && !*all_flag {
		fmt.Println("To display solutions use --day=N, --days=X,Y,Z, or --all")
	}

	// Solutions
	for n_day := 0; n_day < len(day_mapping); n_day++ {
		if slices.Contains(days_to_display, strconv.Itoa(n_day+1)) || *all_flag {
			s1, s2 := day_mapping[n_day]()
			fmt.Println("AoC 2023 Day", n_day+1)
			fmt.Printf("\tSolution 1: %d\n\tSolution 2: %d\n\n", s1, s2)
		}
	}
}

var ART = []string {
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
