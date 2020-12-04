package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

func day4bValidPassport(passport string) bool {
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
	{ // byr (Birth Year) - four digits; at least 1920 and at most 2002.
		byr, err := strconv.Atoi(fieldMap["byr"])
		if err != nil || byr < 1920 || byr > 2002 {
			return false
		}
	}
	{ // iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		iyr, err := strconv.Atoi(fieldMap["iyr"])
		if err != nil || iyr < 2010 || iyr > 2020 {
			return false
		}
	}
	{ // eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		eyr, err := strconv.Atoi(fieldMap["eyr"])
		if err != nil || eyr < 2020 || eyr > 2030 {
			return false
		}
	}
	{ // hgt (Height) - a number followed by either cm or in:
		result := regexp.MustCompile(`^(\d+)(cm|in)$`).FindStringSubmatch(fieldMap["hgt"])
		if result == nil {
			return false
		}
		height, _ := strconv.Atoi(result[1])
		switch result[2] {
		case "cm": // If cm, the number must be at least 150 and at most 193.
			if height < 150 || height > 193 {
				return false
			}
		case "in": // If in, the number must be at least 59 and at most 76.
			if height < 59 || height > 76 {
				return false
			}
		default:
			return false
		}
	}
	{ // hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		result := regexp.MustCompile(`^#[0-9a-f]{6}$`).FindStringSubmatch(fieldMap["hcl"])
		if result == nil {
			return false
		}
	}
	{ // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		switch fieldMap["ecl"] {
		case "amb":
		case "blu":
		case "brn":
		case "gry":
		case "grn":
		case "hzl":
		case "oth":
		default:
			return false
		}
	}
	{ // pid (Passport ID) - a nine-digit number, including leading zeroes.
		result := regexp.MustCompile(`^[0-9]{9}$`).FindStringSubmatch(fieldMap["pid"])
		if result == nil {
			return false
		}
	}
	{ // cid (Country ID) - ignored, missing or not.
	}
	return true
}

func day4b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	curPassport := ""
	validPassports := 0

	for line := range c {
		if line == "" {
			if day4bValidPassport(curPassport) {
				fmt.Printf("valid: %v\n", curPassport)
				validPassports++
			} else {
				fmt.Printf("INVALID: %v\n", curPassport)
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
	if day4bValidPassport(curPassport) {
		fmt.Printf("valid: %v\n", curPassport)
		validPassports++
	} else {
		fmt.Printf("INVALID: %v\n", curPassport)
	}

	return validPassports, nil
}
