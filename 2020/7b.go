package main

import (
	"regexp"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

type bagQuant struct {
	num    int
	colour string
}

func day7b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	contains := make(map[string][]bagQuant)
	ruleParserC := regexp.MustCompile(`^(\w+ \w+) bags? contain (.*)\.$`)
	ruleParserI := regexp.MustCompile(`^(\d+) (\w+ \w+) bags?`)

	for rule := range c {
		// fmt.Println(rule)
		ruleParts := ruleParserC.FindStringSubmatch(rule)
		outerBag := ruleParts[1]
		innerRules := strings.Split(ruleParts[2], ", ")
		contains[outerBag] = make([]bagQuant, 0)
		if innerRules[0] != "no other bags" {
			for _, innerRule := range innerRules {
				ruleParts := ruleParserI.FindStringSubmatch(innerRule)
				num, _ := strconv.Atoi(ruleParts[1])
				innerBag := ruleParts[2]
				// fmt.Printf(" %v contains %v\n", outerBag, innerBag)
				contains[outerBag] = append(contains[outerBag], bagQuant{num, innerBag})
			}
		}
	}

	numBagsInsideMemo := make(map[string]int)
	var numBagsInside func(string) int
	numBagsInside = func(colour string) int {
		if qty, ok := numBagsInsideMemo[colour]; ok {
			return qty
		}
		total := 0
		for _, bagQ := range contains[colour] {
			total += bagQ.num * (numBagsInside(bagQ.colour) + 1)
		}
		return total
	}

	return numBagsInside("shiny gold"), nil
}
