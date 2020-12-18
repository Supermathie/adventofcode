package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

func aevaluate(line string) int {
	total := 0
	tokens := strings.Split(line, " ")
	for i := 0; i < len(tokens); i++ {
		if i == 0 {
			total, _ = strconv.Atoi(tokens[i])
			continue
		}
		switch tokens[i] {
		case "+":
			val, err := strconv.Atoi(tokens[i+1])
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			total += val
			i++
		case "*":
			val, err := strconv.Atoi(tokens[i+1])
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			total *= val
			i++
		default:
			log.Fatalf("unexpected token: %v", tokens[i])
		}
	}
	return total
}

func aevaluateP(line string) int {
	pMatch := regexp.MustCompile(`\(([^())]+)\)`)

	for match := pMatch.FindStringSubmatchIndex(line); match != nil; match = pMatch.FindStringSubmatchIndex(line) {
		subExp := line[match[2]:match[3]]
		val := aevaluate(subExp)
		line = fmt.Sprintf("%s%d%s", line[:match[0]], val, line[match[3]+1:])
	}
	return aevaluate(line)
}

func day18a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}
	total := 0
	for line := range c {
		total += aevaluateP(line)
	}

	return total, nil
}
