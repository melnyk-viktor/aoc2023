package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var day_mapping = []func(){
	day1.Solution,
	day2.Solution,
	day3.Solution,
}

func main() {

	// Decoration
	fmt.Println(`
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
	`)

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
