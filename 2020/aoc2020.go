package main

import (
	"fmt"
	"log"
	"os"
)

type adventFunc func(string) (int, error)
type adventError string

func (e adventError) Error() string {
	return string(e)
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("need a day and inputfile")
	}
	day := os.Args[1]
	inputFile := os.Args[2]

	dayFunc := map[string]adventFunc{
		"1a":  day1a,
		"1b":  day1b,
		"2a":  day2a,
		"2b":  day2b,
		"3a":  day3a,
		"3b":  day3b,
		"4a":  day4a,
		"4b":  day4b,
		"5a":  day5a,
		"5b":  day5b,
		"6a":  day6a,
		"6b":  day6b,
		"7a":  day7a,
		"7b":  day7b,
		"8a":  day8a,
		"8b":  day8b,
		"9a":  day9a,
		"9b":  day9b,
		"10a": day10a,
		"10b": day10b,
		"11a": day11a,
		"11b": day11b,
		"12a": day12a,
		"12b": day12b,
		"13a": day13a,
		"13b": day13b,
		"14a": day14a,
		"14b": day14b,
		"15a": day15a,
		"15b": day15b,
		"16a": day16a,
		"16b": day16b,
		"17a": day17a,
		"17b": day17b,
		"18a": day18a,
		"18b": day18b,
		"19a": day19a,
		"19b": day19b,
		"20a": day20a,
		"20b": day20b,
		"21a": day21a,
		"21b": day21b,
		"22a": day22a,
		"22b": day22b,
		"23a": day23a,
		"23b": day23b,
		"24a": day24a,
		"24b": day24b,
		"25a": day25a,
		"25b": day25b,
	}[day]

	if dayFunc == nil {
		log.Fatalf("unknown day: %s", day)
	}

	result, err := dayFunc(inputFile)
	if err != nil {
		log.Fatalf("%s failed: %s", day, err)
	}
	fmt.Println(result)

}
