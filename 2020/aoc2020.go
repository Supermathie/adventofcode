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
		"1a": day1a,
		"1b": day1b,
		"2a": day2a,
		"2b": day2b,
		// "3a": day1b,
		// "3b": day1b,
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
