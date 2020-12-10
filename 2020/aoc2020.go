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
