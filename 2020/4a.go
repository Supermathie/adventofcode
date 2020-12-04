package main

import (
	"fmt"
	"regexp"
	"strings"

	"supermathie.net/libadvent"
)

func day4aIsValidPassport(passport string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	fields := regexp.MustCompile(`\s+`).Split(passport, -1)
	fieldMap := make(map[string]string)

	for _, field := range fields {
		fdata := strings.SplitN(field, ":", 2)
		fieldMap[fdata[0]] = fdata[1]
	}
	for _, field := range requiredFields {
		if _, exists := fieldMap[field]; !exists {
			return false
		}
	}
	return true
}

func day4a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	curPassport := ""
	validPassports := 0

	for line := range c {
		if line == "" {
			if day4aIsValidPassport(curPassport) {
				// fmt.Printf("valid: %v\n", curPassport)
				validPassports++
			} else {
				// fmt.Printf("INVALID: %v\n", curPassport)
			}
			curPassport = ""
		} else {
			if curPassport == "" {
				curPassport = line
			} else {
				curPassport += " " + line
			}
		}
	}
	if day4aIsValidPassport(curPassport) {
		// fmt.Printf("valid: %v\n", curPassport)
		validPassports++
	} else {
		// fmt.Printf("INVALID: %v\n", curPassport)
		fmt.Println(curPassport)
	}

	return validPassports, nil
}
