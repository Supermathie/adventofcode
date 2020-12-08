package main

import (
	"log"
	"regexp"
	"strconv"

	"supermathie.net/libadvent"
)

func testDay8b(origCode []gameInst, addr int) (int, bool) {
	code := make([]gameInst, len(origCode))
	copy(code, origCode)

	switch code[addr].op {
	case "nop":
		code[addr].op = "jmp"
	case "jmp":
		code[addr].op = "nop"
	default:
		return 0, false
	}

	pc := 0
	acc := 0
	instSeen := make(map[int]bool)
	for (instSeen[pc] != true) && pc < len(code) {
		instSeen[pc] = true
		inst := code[pc]
		switch inst.op {
		case "nop":
		case "acc":
			acc += inst.val
		case "jmp":
			pc += inst.val - 1
		default:
			log.Fatalf("unknown instruction: pc:%v", pc)
		}
		pc++
	}

	if pc == len(code) {
		return acc, true
	}
	return 0, false
}

func day8b(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}

	code := make([]gameInst, 0)

	gameInstParser := regexp.MustCompile(`^(nop|acc|jmp) ([+-]\d+)$`)

	for inst := range c {
		instParts := gameInstParser.FindStringSubmatch(inst)
		op := instParts[1]
		val, _ := strconv.Atoi(instParts[2])
		code = append(code, gameInst{op, val})
	}

	for i := 0; i < len(code); i++ {
		if acc, ok := testDay8b(code, i); ok {
			return acc, nil
		}
	}
	return -1, nil
}
