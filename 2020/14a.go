package main

import (
	"log"
	"regexp"
	"strconv"

	"supermathie.net/libadvent"
)

func day14a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	mem := map[int]int{}
	orMask := 0
	andMask := 0
	matcher := regexp.MustCompile(`^(mask|mem)(?:\[(\d+)\])? = (.*)$`)

	for line := range c {
		cmd := matcher.FindStringSubmatch(line)
		if cmd == nil {
			log.Fatalf("regex does not match line: %v", line)
		}
		switch cmd[1] {
		case "mask":
			orMask = 0
			andMask = (1 << 36) - 1
			for i, v := range cmd[3] {
				switch v {
				case '0':
					andMask ^= 1 << (35 - i)
				case '1':
					orMask |= 1 << (35 - i)
				case 'X':
					// do nothing
				default:
					log.Fatalf("bad mask bit %v", v)
				}
			}
		case "mem":
			memPos, _ := strconv.Atoi(cmd[2])
			value, _ := strconv.Atoi(cmd[3])
			mem[memPos] = (value | orMask) & andMask
		}
	}
	total := 0
	for i := range mem {
		total += mem[i]
	}

	return total, nil
}
