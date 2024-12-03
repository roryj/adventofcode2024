package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type report struct {
	levels []int
}

func map_list[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func parse_part_1_input(input string) []report {
	var reports []report

	for _, r := range strings.Split(input, "\n") {
		levels := map_list(strings.Fields(r), func(lvlStr string) int {
			result, _ := strconv.Atoi(lvlStr)
			return result
		})

		reports = append(reports, report{levels: levels})
	}

	return reports
}

func is_level_safe(levels []int) bool {
	fmt.Println("Starting report check")
	fmt.Printf("Looking at %v\n", levels)

	firstLevel := levels[0]
	secondLevel := levels[1]

	shouldIncrease := firstLevel < secondLevel
	ok := true

	lastLevel := firstLevel

	for i := 1; i < len(levels); i++ {

		currLevel := levels[i]
		levelDifference := int(math.Abs(float64(currLevel - lastLevel)))

		fmt.Printf("difference between %d -> %d is %d\n", lastLevel, currLevel, levelDifference)

		if currLevel == lastLevel {
			ok = false
			fmt.Println("same level. Unsafe")
			break
		}

		if shouldIncrease && currLevel < lastLevel {
			ok = false
			fmt.Println("decreasing but should be increasing. Unsafe")
			break
		}

		if !shouldIncrease && currLevel > lastLevel {
			ok = false
			fmt.Println("increasing but should be decreasing. Unsafe")
			break
		}

		if levelDifference != 1 && levelDifference != 2 && levelDifference != 3 {
			ok = false
			break
		}

		lastLevel = currLevel
	}
	return ok
}

func Part_1_CalculateSafeReports(input string) int {
	reports := parse_part_1_input(input)

	safeReportCount := 0

	for _, report := range reports {
		if is_level_safe(report.levels) {
			safeReportCount++
		}
	}

	return safeReportCount
}

func Part_2_CalculateSafeReportsWithDampener(input string) int {
	reports := parse_part_1_input(input)

	safeReportCount := 0

	for _, report := range reports {

		foundSafeReport := false

		original_levels := make([]int, len(report.levels))
		copy(original_levels, report.levels)
		index_to_remove := 0

		levels_to_test := make([]int, len(report.levels))
		copy(levels_to_test, original_levels)

		for !foundSafeReport {
			copy(original_levels, report.levels)
			fmt.Printf("report levels: %v\n", report.levels)
			fmt.Printf("og_levels: %v\n", original_levels)
			fmt.Printf("curr testing levels: %v\n", levels_to_test)

			if is_level_safe(levels_to_test) {
				fmt.Println("levels are safe!")
				safeReportCount++
				break
			}

			// check to see if we are at the end
			if index_to_remove >= len(original_levels) {
				break
			}

			fmt.Printf("array %v did not pass. Going to remove item at index %d\n", levels_to_test, index_to_remove)
			fmt.Printf("left side of og: %v\n", original_levels[:index_to_remove])
			fmt.Printf("right side of og: %v\n", original_levels[index_to_remove+1:])

			tmp_array := make([]int, len(original_levels))
			copy(tmp_array, report.levels)

			levels_to_test = append(tmp_array[:index_to_remove], tmp_array[(index_to_remove+1):]...)
			fmt.Printf("the new levels we are testing: %v\n", levels_to_test)
			index_to_remove++
		}

	}

	return safeReportCount
}
