package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

func bevaluate(line string) int {
	total := 0
	tokens := strings.Split(line, " ")
	for i := 0; i < len(tokens); i++ {
		if i == 0 {
			total, _ = strconv.Atoi(tokens[i])
			continue
		}
		switch tokens[i] {
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

func bevaluateA(line string) int {
	aMatch := regexp.MustCompile(`(\d+) \+ (\d+)`)

	for match := aMatch.FindStringSubmatchIndex(line); match != nil; match = aMatch.FindStringSubmatchIndex(line) {
		val1, _ := strconv.Atoi(line[match[2]:match[3]])
		val2, _ := strconv.Atoi(line[match[4]:match[5]])
		if match[1] >= len(line) {
			line = fmt.Sprintf("%s%d", line[:match[0]], val1+val2)
		} else {
			line = fmt.Sprintf("%s%d%s", line[:match[0]], val1+val2, line[match[1]:])
		}

	}
	return bevaluate(line)
}

func bevaluateP(line string) int {
	pMatch := regexp.MustCompile(`\(([^())]+)\)`)

	for match := pMatch.FindStringSubmatchIndex(line); match != nil; match = pMatch.FindStringSubmatchIndex(line) {
		subExp := line[match[2]:match[3]]
		val := bevaluateA(subExp)
		if match[3]+1 > len(line) {
			line = fmt.Sprintf("%s%d", line[:match[0]], val)
		} else {
			line = fmt.Sprintf("%s%d%s", line[:match[0]], val, line[match[3]+1:])
		}
	}
	return bevaluateA(line)
}

func day18b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}
	total := 0
	for line := range c {
		total += bevaluateP(line)
	}

	return total, nil
}
