package day5

import (
	"fmt"
	"strconv"
	"strings"

	"roryj.ca/aoc2024/helpers"
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

func is_valid_queue(printQ []int, rules map[int][]int) bool {
	var seenItems []int

	for _, page := range printQ {
		fmt.Printf("Checking page: %v\n", page)

		// if the seen items is 0, then we can just add it, it will pass all rules
		if len(seenItems) == 0 {
			fmt.Println("empty queue, all good")
			seenItems = append(seenItems, page)
			continue
		}

		// now, for the current page, check to see if there are any rules that says it needs to be
		// before any pages before it. This isnt efficient, but meh
		allRequiredAfterPages, ok := rules[page]
		if !ok {
			// no rule for page, page is ok!
			fmt.Println("no rules for the current page, all good")
			seenItems = append(seenItems, page)
			continue
		}

		for _, afterPage := range allRequiredAfterPages {
			for _, seen := range seenItems {
				// fmt.Printf("checking if %v is in %v\n", afterPage, seen)
				if afterPage == seen {
					// the order is wrong, a page that needs to be after is before the current page
					fmt.Println("there is a page that breaks the printing order. Skipping")
					return false
				}
			}
		}
		seenItems = append(seenItems, page)
	}

	return true
}

func isValidRule(first int, second int, rules map[int][]int) bool {

	secondRules, ok := rules[second]

	if !ok {
		return true
	}

	for _, r := range secondRules {
		if first == r {
			return false
		}
	}

	return true
}

func getActiveRulesForQ(printQ []int, rules map[int][]int) map[int][]int {
	activeRules := make(map[int][]int, 0)

	for _, page := range printQ {
		rule := rules[page]

		var activeRule []int

		for _, r := range rule {
			for _, p := range printQ {
				if r == p {
					activeRule = append(activeRule, r)
				}
			}
		}

		fmt.Printf("%s[%v] -> %v\n", strings.Repeat(" ", 4), page, activeRule)
		// _, ok := activeRules[page]
		// if !ok {
		// 	activeRules[page] = make([]int, 0)
		// }
		activeRules[page] = activeRule
	}

	return activeRules
}

func Part_1_validate_print_queue(input string) int {
	queues, rules := parse_input_file(input)

	fmt.Printf("queues: %v & rules: %v", queues, rules)

	var okQs [][]int

	for _, printQ := range queues {
		fmt.Printf("Checking queue: %v\n", printQ)
		// all rules for current queue
		fmt.Println("Rules for queue:")
		for _, page := range printQ {
			rule := rules[page]

			var activeRule []int

			for _, r := range rule {
				for _, p := range printQ {
					if r == p {
						activeRule = append(activeRule, r)
					}
				}
			}

			fmt.Printf("%s[%v] -> %v\n", strings.Repeat(" ", 4), page, activeRule)
		}

		var seenItems []int
		validQ := true

		for _, page := range printQ {
			fmt.Printf("Checking page: %v\n", page)

			// if the seen items is 0, then we can just add it, it will pass all rules
			if len(seenItems) == 0 {
				fmt.Println("empty queue, all good")
				seenItems = append(seenItems, page)
				continue
			}

			// now, for the current page, check to see if there are any rules that says it needs to be
			// before any pages before it. This isnt efficient, but meh
			allRequiredAfterPages, ok := rules[page]
			if !ok {
				// no rule for page, page is ok!
				fmt.Println("no rules for the current page, all good")
				seenItems = append(seenItems, page)
				continue
			}

			for _, afterPage := range allRequiredAfterPages {
				for _, seen := range seenItems {
					// fmt.Printf("checking if %v is in %v\n", afterPage, seen)
					if afterPage == seen {
						// the order is wrong, a page that needs to be after is before the current page
						validQ = false
						fmt.Println("there is a page that breaks the printing order. Skipping")
						break
					}
				}
				if !validQ {
					break
				}
			}
			if !validQ {
				break
			}

			seenItems = append(seenItems, page)
		}

		// if we have seen all items, then we are ok and the page order is valid
		if validQ {
			okQs = append(okQs, printQ)
		}
	}

	// now calculate the result needed, which is the middle item in each ok queue and sum them together

	result := 0

	for _, q := range okQs {
		result += q[len(q)/2]
	}

	return result
}

func Part_2_incorrect_only_updates(input string) int {
	queues, rules := parse_input_file(input)

	invalidQs := helpers.Filter(queues, func(t []int) bool {
		return !is_valid_queue(t, rules)
	})

	fmt.Printf("the invalid queues to process: %v\n", invalidQs)

	fixedQs := [][]int{}

	for _, q := range invalidQs {

		activeRules := getActiveRulesForQ(q, rules)

		fmt.Printf("%v\n", activeRules)

		count := 0

		for {
			if count > 100 {
				break
			}
			valid := true
			// now we loop through each item in the queue, and if there is a rule saying that the
			// item is in the wrong spot, we swap their spots!
			for i := range len(q) - 1 {

				fmt.Printf("queue state: %v\n", q)

				j := i + 1
				for j < len(q) {
					currentEntry := q[i]
					compared := q[i+1]

					if !isValidRule(currentEntry, compared, rules) {
						q[i] = compared
						q[i+1] = currentEntry
						valid = false
						break
					}

					j++
				}
			}
			if valid {
				fixedQs = append(fixedQs, q)
				break
			}
		}

	}

	result := 0
	for _, q := range fixedQs {
		result += q[len(q)/2]
	}

	return result
}
