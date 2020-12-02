package main

import (
	"log"
	"regexp"
	"strconv"

	"supermathie.net/libadvent"
)

// The policy (before the : ) describes two one-indexed positions. Exactly one
// of these positions in the password must be the specified character to be valid.
func day2bIsValidPassword(pass string) bool {
	// 1-7 h: hhlnhfhzxhhphhdhh
	r := regexp.MustCompile(`^(\d+)-(\d+)\s+(.):\s+(.*)$`)
	result := r.FindStringSubmatch(pass)
	if result == nil {
		log.Fatalf("regex did not match string:\n%s", pass)
	}
	pos1, _ := strconv.Atoi(result[1])
	pos2, _ := strconv.Atoi(result[2])
	letter := result[3][0]
	password := result[4]
	return (password[pos1-1] == letter) != (password[pos2-1] == letter)
}

func day2b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	var count int
	for line := range c {
		if day2bIsValidPassword(line) {
			count++
		}
	}

	return count, nil
}
