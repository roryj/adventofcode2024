package day1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, []int) {
	var listOne []int
	var listTwo []int

	for _, r := range strings.Split(input, "\n") {
		entries := strings.Fields(r)
		entryOne, _ := strconv.Atoi(entries[0])
		entryTwo, _ := strconv.Atoi(entries[1])
		listOne = append(listOne, entryOne)
		listTwo = append(listTwo, entryTwo)
	}

	return listOne, listTwo
}

func parseInputToFrequencyMap(input string) (map[int]int, map[int]int) {
	var leftListMap map[int]int = make(map[int]int)
	var rightListMap map[int]int = make(map[int]int)

	for _, r := range strings.Split(input, "\n") {
		entries := strings.Fields(r)
		entryOne, _ := strconv.Atoi(entries[0])
		entryTwo, _ := strconv.Atoi(entries[1])

		val, ok := leftListMap[entryOne]
		if !ok {
			leftListMap[entryOne] = 1
		} else {
			leftListMap[entryOne] = val + 1
		}

		val2, ok := rightListMap[entryTwo]
		if !ok {
			rightListMap[entryTwo] = 1
		} else {
			rightListMap[entryTwo] = val2 + 1
		}
	}

	return leftListMap, rightListMap
}

// Find the total distance between lists
func Part1_FindDifferenceDistanceBetweenLists(input string) int {
	numberListOne, numberListTwo := parseInput(input)

	sort.Slice(numberListOne, func(i, j int) bool {
		return numberListOne[i] < numberListOne[j]
	})
	sort.Slice(numberListTwo, func(i, j int) bool {
		return numberListTwo[i] < numberListTwo[j]
	})

	totalDistance := 0

	for x := range len(numberListOne) {
		rowDistance := int(math.Abs(float64(numberListOne[x] - numberListTwo[x])))
		totalDistance += rowDistance
	}

	return totalDistance
}

func Part2_FindSimilarity(input string) int {
	leftListMap, rightListMap := parseInputToFrequencyMap(input)

	similarityScore := 0

	for k, v := range leftListMap {
		// for the current key, check to see how many times it appears in the right list
		numTimes := rightListMap[k]

		// get the similarity score for the entry by multiplying the number by how many times
		// it appears in the right list
		// NOTE: we are also multiplying by the number of times the item appears in the left list, as
		// that is also important to do here
		currSimilarityScore := k * numTimes * v

		// add it to the total
		similarityScore += currSimilarityScore
	}

	return similarityScore
}
