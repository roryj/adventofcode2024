package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func parse_input_file(input string) ([][]int, map[int][]int) {

	rules := make(map[int][]int)
	var printQueues [][]int

	splitInput := strings.Split(input, "\n\n")

	// first half is the rules
	for _, r := range strings.Split(splitInput[0], "\n") {
		splitRow := strings.Split(r, "|")
		beforePage, _ := strconv.Atoi(splitRow[0])
		afterPage, _ := strconv.Atoi(splitRow[1])

		existingRules, ok := rules[beforePage]
		if !ok {
			existingRules = []int{}
		}

		existingRules = append(existingRules, afterPage)
		rules[beforePage] = existingRules
	}

	for _, queue := range strings.Split(splitInput[1], "\n") {
		var q []int
		for _, entry := range strings.Split(queue, ",") {
			i, _ := strconv.Atoi(entry)
			q = append(q, i)
		}

		printQueues = append(printQueues, q)
	}

	return printQueues, rules
}

func Part_1_validate_print_queue(input string) int {
	queues, rules := parse_input_file(input)

	fmt.Printf("queues: %v & rules: %v", queues, rules)

	return 0
}
