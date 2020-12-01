package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("need a day and inputfile")
	}
	day := os.Args[1]
	inputFile := os.Args[2]

	switch day {
	case "1a":
		fmt.Println(day1a(inputFile))
	case "1b":
		fmt.Println(day1b(inputFile))
	default:
		log.Fatalf("unknown day: %s", day)
	}
}
