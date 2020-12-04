package main

import (
	"testing"
)

func Test4aCase1(t *testing.T) {
	data := "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"
	result := day4aIsValidPassport(data)
	want := true
	if result != want {
		t.Fatalf("day4a(%s) → %v, expected %v", data, result, want)
	}
}
func Test4aCase2(t *testing.T) {
	data := "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929"
	result := day4aIsValidPassport(data)
	want := false
	if result != want {
		t.Fatalf("day4a(%s) → %v, expected %v", data, result, want)
	}
}
func Test4aCase3(t *testing.T) {
	data := "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm"
	result := day4aIsValidPassport(data)
	want := true
	if result != want {
		t.Fatalf("day4a(%s) → %v, expected %v", data, result, want)
	}
}
func Test4aCase4(t *testing.T) {
	data := "hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in"
	result := day4aIsValidPassport(data)
	want := false
	if result != want {
		t.Fatalf("day4a(%s) → %v, expected %v", data, result, want)
	}
}
