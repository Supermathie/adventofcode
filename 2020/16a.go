package main

import (
	"regexp"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

func validateTicket(validators map[string]func(int) bool, ticket []int) (checksum int, ticketValid bool) {
	checksum = 0
	ticketValid = true
	for _, v := range ticket {
		valid := false
		for _, f := range validators {
			if f(v) {
				valid = true
				break
			}
		}
		if !valid {
			ticketValid = false
			checksum += v
		}
	}
	return
}

func validValidators(validators map[string]func(int) bool, values []int) (valid []string) {
	valid = make([]string, 0)

	for name, f := range validators {
		ok := true
		for _, value := range values {
			if !f(value) {
				ok = false
				break
			}
		}
		if ok {
			valid = append(valid, name)
		}
	}
	return
}

func parseTicketLine(line string) (ticket []int) {
	ticket = make([]int, 0)
	for _, v := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(v)
		ticket = append(ticket, num)
	}
	return
}

func day16(inputFile string) (int, int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, -1, err
	}

	validators := map[string]func(int) bool{}
	validTickets := make([][]int, 0)
	invalidTotal := 0

	validatorMatcher := regexp.MustCompile(`^([\w ]+): (\d+)-(\d+) or (\d+)-(\d+)$`)

	for line := range c {
		if line == "" {
			break
		}
		match := validatorMatcher.FindStringSubmatch(line)
		validatorName := match[1]
		min1, _ := strconv.Atoi(match[2])
		max1, _ := strconv.Atoi(match[3])
		min2, _ := strconv.Atoi(match[4])
		max2, _ := strconv.Atoi(match[5])

		validators[validatorName] = func(val int) bool {
			return (min1 <= val && val <= max1) || (min2 <= val && val <= max2)
		}
	}

	<-c // dump 'your ticket' line
	myTicket := parseTicketLine(<-c)

	<-c // dump blank line
	<-c // dump 'nearby tickets' line

	for line := range c {
		ticket := parseTicketLine(line)
		checksum, ok := validateTicket(validators, ticket)
		if ok {
			validTickets = append(validTickets, ticket)
		} else {
			invalidTotal += checksum
		}
	}

	validColumnValidators := make([][]string, len(myTicket))
	for col := 0; col < len(myTicket); col++ {
		values := make([]int, len(validTickets))
		for i, ticket := range validTickets {
			values[i] = ticket[col]
		}
		validColumnValidators[col] = validValidators(validators, values)
	}

	columnMap := make(map[int]string)
	for len(columnMap) < len(validators) {
		for col, options := range validColumnValidators {
			if len(options) == 1 {
				colName := options[0]
				columnMap[col] = colName
				for i := 0; i < len(validColumnValidators); i++ {
					index, found := libadvent.IndexOfS(validColumnValidators[i], colName)
					if found {
						if index == len(validColumnValidators[i])-1 {
							validColumnValidators[i] = validColumnValidators[i][:index]
						} else {
							validColumnValidators[i] = append(validColumnValidators[i][:index], validColumnValidators[i][index+1:]...)
						}
					}
				}
			}
		}
	}

	ticketValue := 1
	for i, name := range columnMap {
		if strings.Contains(name, "departure") {
			ticketValue *= myTicket[i]
		}
	}
	return invalidTotal, ticketValue, nil
}

func day16a(inputFile string) (int, error) {
	a, _, err := day16(inputFile)
	return a, err
}

// validators = {}
// validators['departure_location'] = lambda x: (x >= 42 and x <= 322) or (x >= 347 and x <= 954)
// validators['departure_station']  = lambda x: (x >= 49 and x <= 533) or (x >= 555 and x <= 966)
// validators['departure_platform'] = lambda x: (x >= 28 and x <= 86 ) or (x >= 101 and x <= 974)
// validators['departure_track']    = lambda x: (x >= 50 and x <= 150) or (x >= 156 and x <= 950)
// validators['departure_date']     = lambda x: (x >= 30 and x <= 117) or (x >= 129 and x <= 957)
// validators['departure_time']     = lambda x: (x >= 31 and x <= 660) or (x >= 678 and x <= 951)
// validators['arrival_location']   = lambda x: (x >= 26 and x <= 482) or (x >= 504 and x <= 959)
// validators['arrival_station']    = lambda x: (x >= 29 and x <= 207) or (x >= 220 and x <= 971)
// validators['arrival_platform']   = lambda x: (x >= 28 and x <= 805) or (x >= 829 and x <= 964)
// validators['arrival_track']      = lambda x: (x >= 48 and x <= 377) or (x >= 401 and x <= 964)
// validators['tclass']             = lambda x: (x >= 28 and x <= 138) or (x >= 145 and x <= 959)
// validators['duration']           = lambda x: (x >= 33 and x <= 182) or (x >= 205 and x <= 966)
// validators['price']              = lambda x: (x >= 25 and x <= 437) or (x >= 449 and x <= 962)
// validators['route']              = lambda x: (x >= 41 and x <= 403) or (x >= 428 and x <= 968)
// validators['row']                = lambda x: (x >= 33 and x <= 867) or (x >= 880 and x <= 960)
// validators['seat']               = lambda x: (x >= 40 and x <= 921) or (x >= 930 and x <= 955)
// validators['train']              = lambda x: (x >= 47 and x <= 721) or (x >= 732 and x <= 955)
// validators['ttype']              = lambda x: (x >= 33 and x <= 243) or (x >= 265 and x <= 964)
// validators['wagon']              = lambda x: (x >= 31 and x <= 756) or (x >= 768 and x <= 973)
// validators['zone']               = lambda x: (x >= 50 and x <= 690) or (x >= 713 and x <= 967)
