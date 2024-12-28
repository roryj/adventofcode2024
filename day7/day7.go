package day7

import (
	"fmt"
	"strconv"
	"strings"

	"roryj.ca/aoc2024/helpers"
)

type equation struct {
	answer  int
	numbers []int
}

func parse_input(input string) []equation {

	result := []equation{}

	for _, row := range strings.Split(input, "\n") {
		firstSplit := strings.Split(row, ":")
		answer, _ := strconv.Atoi(firstSplit[0])
		numbers := helpers.Map(strings.Fields(firstSplit[1]), func(s string) int {
			r, _ := strconv.Atoi(s)
			return r
		})

		result = append(result, equation{answer: answer, numbers: numbers})
	}

	return result
}

type Operator int

const (
	Multiply Operator = 0
	Addition Operator = 1
	Concat   Operator = 2
)

func (o Operator) op(a int, b int) int {
	switch o {
	case Multiply:
		return a * b
	case Addition:
		return a + b
	case Concat:
		v, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		if err != nil {
			panic("somehow invalid number?")
		}
		return v
	default:
		panic("invalid operator")
	}
}

func combinations(operators []Operator, k int) [][]Operator {

	result := [][]Operator{}

	// Helper function to generate combinations recursively.
	var backtrack func(combination []Operator)
	backtrack = func(combination []Operator) {
		if len(combination) == k {
			// Make a copy of the combination and add it to the result.
			combinationCopy := make([]Operator, len(combination))
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
			return
		}

		for _, op := range operators {
			// Add operator to the current combination.
			combination = append(combination, op)
			// Recursively generate the next operator.
			backtrack(combination)
			// Backtrack by removing the last added operator.
			combination = combination[:len(combination)-1]
		}
	}

	backtrack([]Operator{})
	return result

	// results := [][]Operator{}

	// cs := combin.Combinations(len(operators), k)
	// for _, c := range cs {

	// 	combo := []Operator{}
	// 	for x := range k {
	// 		combo = append(combo, operators[c[x]])
	// 	}
	// 	results = append(results, combo)
	// 	// fmt.Printf("%s%s\n", operators[c[0]], operators[c[1]])
	// }

	// return results

	// [1, 2, 3, 4]
	//   +  +  +
	//   *  +  +
	//   +  *  +
	//   +  +  *
	//   *  +  *
	//   *  *  +
	//   +  *  *
	//   *  *  *
	// if k == 0 {
	// 	return [][]Operator{{}}
	// }
	// if len(operators) == 0 {
	// 	return [][]Operator{}
	// }

	// result := [][]Operator{}

	// // Include the first element
	// for _, c := range combinations(operators[1:], k-1) {
	// 	result = append(result, append([]Operator{operators[0]}, c...))
	// }

	// // Exclude the first element
	// for _, c := range combinations(operators[1:], k) {
	// 	result = append(result, c)
	// }

	// return result
}

func Part_1_CalibrateEquations(input string) (int, int) {

	equations := parse_input(input)

	part1Result := 0

	for _, row := range equations {
		// fmt.Printf("details => %v and %v\n", row.answer, row.numbers)
		combos := combinations([]Operator{Multiply, Addition}, len(row.numbers)-1)
		// fmt.Printf("combos: %v\n", combos)

		for _, combo := range combos {
			iteration := 0
			thisAnswer := row.numbers[0]
			for i := 1; i < len(row.numbers); i++ {
				thisAnswer = combo[iteration].op(thisAnswer, row.numbers[i])
				iteration++
			}

			if thisAnswer == row.answer {
				part1Result += thisAnswer
				break
			}
		}
	}

	part2Result := 0
	for _, row := range equations {
		combos := combinations([]Operator{Multiply, Addition, Concat}, len(row.numbers)-1)
		// fmt.Printf("combos: %v\n", combos)

		for _, combo := range combos {
			iteration := 0
			thisAnswer := row.numbers[0]
			for i := 1; i < len(row.numbers); i++ {
				thisAnswer = combo[iteration].op(thisAnswer, row.numbers[i])
				iteration++
			}

			if thisAnswer == row.answer {
				part2Result += thisAnswer
				break
			}
		}
	}

	return part1Result, part2Result
}
