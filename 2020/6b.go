package main

import (
	"supermathie.net/libadvent"
)

func day6b(inputFile string) (int, error) {
	blocks, err := libadvent.ReadFileLinesSeparated(inputFile)
	if err != nil {
		return -1, err
	}

	totalAnswerCount := 0

	for _, group := range blocks {
		groupAnswerCount := make(map[string]int)
		for _, personAnswers := range group {
			for i := 0; i < len(personAnswers); i++ {
				groupAnswerCount[personAnswers[i:i+1]]++
			}
		}
		for _, x := range groupAnswerCount {
			// Did everyone in the group answer
			if x == len(group) {
				totalAnswerCount++
			}
		}
	}
	return totalAnswerCount, nil
}
