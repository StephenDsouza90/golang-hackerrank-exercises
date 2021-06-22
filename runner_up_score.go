// https://www.hackerrank.com/challenges/find-second-maximum-number-in-a-list/problem
package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(listOfScores []int) []int {
	dict := make(map[int]int)
	for _, score := range listOfScores {
		dict[score] = 1
	}

	scoresWithoutDuplicate := make([]int, 0)
	for score, _ := range dict {
		scoresWithoutDuplicate = append(scoresWithoutDuplicate, score)
	}
	return scoresWithoutDuplicate
}

func sortScores(scoresWithoutDuplicate []int) []int {
	sort.Slice(scoresWithoutDuplicate, func(i, j int) bool {
		return scoresWithoutDuplicate[i] > scoresWithoutDuplicate[j]
	})

	var sortedListOfScores []int
	for _, score := range scoresWithoutDuplicate {
		sortedListOfScores = append(sortedListOfScores, score)
	}
	return sortedListOfScores
}

func getRunnerUpScore(listOfScores []int) {
	scoresWithoutDuplicate := removeDuplicates(listOfScores)
	sortedListOfScores := sortScores(scoresWithoutDuplicate)

	runnerUpScoreIndex := 1
	runnerUpScore := sortedListOfScores[runnerUpScoreIndex]
	fmt.Println("Runner up score: ", runnerUpScore)
}

func main() {
	listOfScores := []int{2, 3, 6, 6, 5}
	getRunnerUpScore(listOfScores)
}
