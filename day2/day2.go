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

func Part_1_CalculateSafeReports(input string) int {
	reports := parse_part_1_input(input)

	safeReportCount := 0

	for _, report := range reports {

		fmt.Println("Starting report check")
		fmt.Printf("Looking at %v\n", report.levels)

		firstLevel := report.levels[0]
		secondLevel := report.levels[1]

		shouldIncrease := firstLevel < secondLevel
		ok := true

		lastLevel := firstLevel

		for i := 1; i < len(report.levels); i++ {

			currLevel := report.levels[i]

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

			levelDifference := int(math.Abs(float64(currLevel - lastLevel)))

			fmt.Printf("difference between %d -> %d is %d\n", lastLevel, currLevel, levelDifference)

			if levelDifference != 1 && levelDifference != 2 && levelDifference != 3 {
				ok = false
				break
			}

			lastLevel = currLevel
		}

		if ok {
			safeReportCount++
		}
	}

	return safeReportCount
}
