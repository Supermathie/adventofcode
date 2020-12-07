package main

import (
	"regexp"
	"strings"

	"supermathie.net/libadvent"
)

func inSlice(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func day7a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	containedBy := make(map[string][]string)
	ruleParserC := regexp.MustCompile(`^(\w+ \w+) bags? contain (.*)\.$`)
	ruleParserI := regexp.MustCompile(`^(\d+) (\w+ \w+) bags?`)

	for rule := range c {
		// fmt.Println(rule)
		ruleParts := ruleParserC.FindStringSubmatch(rule)
		outerBag := ruleParts[1]
		innerRules := strings.Split(ruleParts[2], ", ")
		if innerRules[0] != "no other bags" {
			for _, innerRule := range innerRules {
				ruleParts := ruleParserI.FindStringSubmatch(innerRule)
				innerBag := ruleParts[2]
				// fmt.Printf(" %v contains %v\n", outerBag, innerBag)
				if containedBy[innerBag] == nil {
					containedBy[innerBag] = []string{outerBag}
				} else {
					containedBy[innerBag] = append(containedBy[innerBag], outerBag)
				}
			}
		}
	}

	possibleBags := make([]string, len(containedBy["shiny gold"]))
	seen := make(map[string]bool)
	seen["shiny gold"] = true
	copy(possibleBags, containedBy["shiny gold"])
	for len(seen)-1 < len(possibleBags) {
		for _, bag := range possibleBags {
			if seen[bag] != true {
				seen[bag] = true
				for _, outerBag := range containedBy[bag] {
					if !inSlice(possibleBags, outerBag) {
						possibleBags = append(possibleBags, outerBag)
					}
				}
			}
		}
	}

	return len(possibleBags), nil
}
