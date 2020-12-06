package main

import (
	"supermathie.net/libadvent"
)

func day6a(inputFile string) (int, error) {
	blocks, err := libadvent.ReadFileLinesSeparated(inputFile)
	if err != nil {
		return -1, err
	}

	totalAnswerCount := 0

	for _, group := range blocks {
		groupAnswerPresent := make(map[string]int)
		for _, personAnswers := range group {
			for i := 0; i < len(personAnswers); i++ {
				groupAnswerPresent[personAnswers[i:i+1]] = 1
			}
		}
		for range groupAnswerPresent {
			totalAnswerCount++
		}
	}
	return totalAnswerCount, nil
}
