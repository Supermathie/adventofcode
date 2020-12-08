package main

import (
	"log"
	"regexp"
	"strconv"

	"supermathie.net/libadvent"
)

type gameInst struct {
	op  string
	val int
}

func day8a(inputFile string) (int, error) {
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

	pc := 0
	acc := 0
	instSeen := make(map[int]bool)
	for instSeen[pc] != true {
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
	return acc, nil
}
