package main

import (
	"fmt"
	"io"
	"os"

	"roryj.ca/aoc2024/day1"
	"roryj.ca/aoc2024/day2"
	"roryj.ca/aoc2024/day3"
	"roryj.ca/aoc2024/day4"
	"roryj.ca/aoc2024/day5"
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

		result2 := day2.Part_2_CalculateSafeReportsWithDampener(string(input))
		fmt.Println("Part 2 Result: ", result2)
	}
	{
		fmt.Println("--- Day 3 ---")
		file, err := os.Open("inputs/day3.txt")
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
		result := day3.Part_1_parse_and_run_corrupted(string(input))
		fmt.Println("Part 1 Result: ", result)

		result2 := day3.Part_2_parse_and_run_corrupted_enable_disable(string(input))
		fmt.Println("Part 2 Result: ", result2)
	}
	{
		fmt.Println("--- Day 4 ---")
		file, err := os.Open("inputs/day4.txt")
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
		result := day4.Part_1_FindAllXmas(string(input))
		fmt.Println("Part 1 Result: ", result)

		result2 := day4.Part_2_FindAllX_mas(string(input))
		fmt.Println("Part 2 Result: ", result2)
	}
	{
		fmt.Println("--- Day 5 ---")
		file, err := os.Open("inputs/day5.txt")
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
		result := day5.Part_1_validate_print_queue(string(input))
		fmt.Println("Part 1 Result: ", result)

		result2 := day5.Part_2_incorrect_only_updates(string(input))
		fmt.Println("Part 2 Result: ", result2)
	}
}
