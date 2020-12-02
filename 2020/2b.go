package main

import (
	"log"
	"regexp"
	"strconv"

	"supermathie.net/libadvent"
)

/*
Each policy actually describes two positions in the password, where 1 means the
first character, 2 means the second character, and so on. (Be careful; Toboggan
Corporate Policies have no concept of "index zero"!) Exactly one of these
positions must contain the given letter. Other occurrences of the letter are
irrelevant for the purposes of policy enforcement.
*/
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
