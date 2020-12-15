package main

import (
	"fmt"
	"strconv"
	"strings"

	"supermathie.net/libadvent"
)

func day15a(inputFile string) (int, error) {
	c, err := libadvent.ReadFileLines(inputFile)
	if err != nil {
		return -1, err
	}
	input := make([]int, 0)
	history := map[int]int{}

	cur := 0
	for i, s := range strings.Split(<-c, ",") {
		num, _ := strconv.Atoi(s)
		input = append(input, num)
		history[num] = i
		cur = num
	}

	for t := len(input); t < 2020; t++ {
		lastTime, spokenBefore := history[cur]
		history[cur] = t - 1
		if spokenBefore {
			cur = t - 1 - lastTime
		} else {
			cur = 0
		}
		fmt.Printf("t=%d, cur=%d\n", t, cur)
	}

	return cur, nil
}
