package main

import (
	"fmt"
	"log"

	"supermathie.net/libadvent"
)

type result struct {
	len   int
	start int
	rule  int
}

func search19(grammar map[int]rule19, s string) bool {
	possible := make(map[result]bool)
	for i := 0; i < len(s); i++ {
		for _, rule := range grammar {
			if rule.terminal == rune(s[i]) {
				possible[result{1, i, rule.num}] = true
			}
		}
	}

	for i := 0; i < len(s); i++ { // start of substring
		changed := false
		for _, rule := range grammar {
			for _, option := range rule.options {
				if len(option) == 1 {
					if possible[result{1, i, option[0]}] {
						if !possible[result{1, i, rule.num}] {
							changed = true
						}
						possible[result{1, i, rule.num}] = true
					}
				}
			}
		}
		if changed {
			i-- // run it again
		}
	}

	for l := 2; l <= len(s); l++ { // length of substring
		changed := false
		for i := 0; i <= len(s)-l+1; i++ { // start of substring
			for p := 1; p <= l-1; p++ { // partition point of substring
				for _, rule := range grammar {
					if possible[result{l, i, rule.num}] {
						continue // we already know this rule works
					}
					for _, option := range rule.options {
						if len(option) == 1 {
							if possible[result{l, i, option[0]}] {
								possible[result{l, i, rule.num}] = true
								changed = true
							}
						} else {
							if len(option) != 2 {
								log.Fatalf("rule:%d option:%v error 2", rule.num, option)
							}
							if possible[result{p, i, option[0]}] && possible[result{l - p, i + p, option[1]}] {
								possible[result{l, i, rule.num}] = true
								changed = true
							}
						}
					}
				}
			}
		}
		if changed {
			l-- // run it again
		}
	}

	return possible[result{len(s), 0, 0}]
}

// 285 too low

func day19b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	grammar := readGrammar(c)

	validMessages := 0
	for line := range c {
		if search19(grammar, line) {
			fmt.Printf("+:%s\n", line)
			validMessages++
		} else {
			fmt.Printf("-:%s\n", line)
		}
	}
	return validMessages, nil
}
