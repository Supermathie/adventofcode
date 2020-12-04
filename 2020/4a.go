package main

import (
	"fmt"
	"regexp"
	"strings"

	"supermathie.net/libadvent"
)

func day4aValidPassport(passport string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	fields := regexp.MustCompile("[\n ]+").Split(passport, -1)
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
			if day4aValidPassport(curPassport) {
				fmt.Println("valid:")
				fmt.Println(curPassport)
				validPassports++
			} else {
				fmt.Println("invalid:")
				fmt.Println(curPassport)
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
	if day4aValidPassport(curPassport) {
		fmt.Println("valid:")
		fmt.Println(curPassport)
		validPassports++
	} else {
		fmt.Println("invalid:")
		fmt.Println(curPassport)
	}

	return validPassports, nil
}
