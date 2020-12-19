package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

type rule19 struct {
	num      int
	terminal rune
	options  [][]int
}

func expand19(grammar map[int]rule19, ruleNum int) (c chan string) {
	c = make(chan string, 8)
	go func() {
		defer close(c)
		rule := grammar[ruleNum]
		if rule.terminal != 0 {
			c <- string(rule.terminal)
		} else {
			for _, option := range rule.options {
				for s1 := range expand19(grammar, option[0]) {
					if len(option) > 1 {
						for s2 := range expand19(grammar, option[1]) {
							c <- s1 + s2
						}
					} else {
						c <- s1
					}
				}
			}
		}
	}()
	return
}

func readGrammar(c chan string) map[int]rule19 {
	ruleMatcher := regexp.MustCompile(`^(\d+): (?:"(a)"|"(b)"|(\d.*))$`)

	rules := make(map[int]rule19)

	for line := range c {
		if line == "" {
			break
		}
		match := ruleMatcher.FindStringSubmatch(line)
		rule := rule19{}
		if match == nil {
			log.Fatalf("matcher did not match line:\n%v", line)
		}
		rule.num, _ = strconv.Atoi(match[1])
		if match[2] == "a" {
			rule.terminal = 'a'
		} else if match[3] == "b" {
			rule.terminal = 'b'
		} else {
			options := strings.Split(match[4], " | ")
			rule.options = make([][]int, 0)
			for i := 0; i < len(options); i++ {
				rule.options = append(rule.options, make([]int, 0))
				for _, s := range strings.Split(options[i], " ") {
					ruleNum, _ := strconv.Atoi(s)
					rule.options[i] = append(rule.options[i], ruleNum)
				}
			}
		}
		rules[rule.num] = rule
	}
	return rules
}

func day19a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	grammar := readGrammar(c)

	possibleMessages := make(map[string]bool, 0)
	for possibleMessage := range expand19(grammar, 0) {
		possibleMessages[possibleMessage] = true
	}

	validMessages := 0
	for line := range c {
		if possibleMessages[line] {
			validMessages++
		}
	}
	return validMessages, nil
}
