package main

import (
	"fmt"
	"io"
	"os"

	"roryj.ca/aoc2024/day1"
	"roryj.ca/aoc2024/day2"
)

// main function
func main() {
	{
		fmt.Println("--- Day 1 ---")
		file, err := os.Open("inputs/day1.txt")
		if err != nil {
			panic(err)
		}
		// close fi on exit and check for its returned error
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()
		input, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		result := day1.Part1_FindDifferenceDistanceBetweenLists(string(input))
		fmt.Println("Part 1 Result: ", result)

		result2 := day1.Part2_FindSimilarity(string(input))
		fmt.Println("Part 2 Result: ", result2)

	}

	{
		fmt.Println("--- Day 2 ---")
		file, err := os.Open("inputs/day2.txt")
		if err != nil {
			panic(err)
		}
		// close fi on exit and check for its returned error
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()
		input, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		result := day2.Part_1_CalculateSafeReports(string(input))
		fmt.Println("Part 1 Result: ", result)

		// result2 := day1.Part2_FindSimilarity(string(input))
		// fmt.Println("Part 2 Result: ", result2)
	}
}
