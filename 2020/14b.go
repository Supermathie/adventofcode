package main

import (
	"log"
	"regexp"
	"strconv"

	"supermathie.net/libadvent"
)

func day14b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	mem := map[int]int{}
	orMask := 0
	floatingBits := []int{}
	matcher := regexp.MustCompile(`^(mask|mem)(?:\[(\d+)\])? = (.*)$`)

	for line := range c {
		cmd := matcher.FindStringSubmatch(line)
		if cmd == nil {
			log.Fatalf("regex does not match line: %v", line)
		}
		switch cmd[1] {
		case "mask":
			orMask = 0
			floatingBits = make([]int, 0)
			for i, v := range cmd[3] {
				switch v {
				case '0':
					// do nothing
				case '1':
					orMask |= 1 << (35 - i)
				case 'X':
					floatingBits = append(floatingBits, (35 - i))
				default:
					log.Fatalf("bad mask bit %v", v)
				}
			}
		case "mem":
			memPos, _ := strconv.Atoi(cmd[2])
			value, _ := strconv.Atoi(cmd[3])
			for combo := range libadvent.AllCombinations(floatingBits) {
				actualMemPos := memPos | orMask
				for _, i := range floatingBits {
					actualMemPos &= ^(1 << i)
				}
				for _, i := range combo {
					actualMemPos |= 1 << i
				}
				// fmt.Printf("writing %v to %v\n", value, actualMemPos)
				mem[actualMemPos] = value
			}
		}
	}

	total := 0
	for i := range mem {
		total += mem[i]
	}
	return total, nil
}
