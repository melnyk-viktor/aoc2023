package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"flag"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"time"
)

var day_mapping = []func(){
	day1.Solution,
	day2.Solution,
	day3.Solution,
}

func main() {

	// Decoration
	rand.Seed(time.Now().Unix())
	fmt.Println(ART[rand.Intn(len(ART))])

	// NOTE: --input and --inputs can be added, but passing list of files for --days and --all flags is too much for simple cli, that is why input is hardcoded.
	// Also it is not elegant to just restrict input files to some folder and/or naming scheme to parse them automatically, but that is also an option.

	day_option := flag.String("day", "", "Run solution for given day")
	days_option := flag.String("days", "", "Comma-separated list of days to run solutions for")
	all_flag := flag.Bool("all", false, "Run all solutions")
	flag.Parse()

	if (*day_option != "" && *days_option != "") || (*day_option != "" && *all_flag) || (*all_flag && *days_option != "") {
		fmt.Println("Flags are exclusive")
		return // NOTE: maybe use os.Exit(1)
	}

	var days_to_display []string
	if *day_option != "" {
		days_to_display = append(days_to_display, *day_option)
	}
	if *days_option != "" {
		days_to_display = append(days_to_display, strings.Split(*days_option, ",")...)
	}

	if len(days_to_display) == 0 && !*all_flag {
		fmt.Println("To display solutions use --day=N, --days=X,Y,Z, or --all")
	}

	// Solutions
	for n_day := 0; n_day < len(day_mapping); n_day++ {
		if slices.Contains(days_to_display, strconv.Itoa(n_day+1)) || *all_flag {
			day_mapping[n_day]()
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
