package day4

import "strings"

func parse_into_array(input string) [][]string {
	var result [][]string

	for _, r := range strings.Split(input, "\n") {
		var row []string

		for _, c := range r {
			row = append(row, string(c))
		}

		result = append(result, row)
	}

	return result
}

func Part_1_FindAllXmas(input string) int {

	wordsearch := parse_into_array(input)

	totalXmasCount := 0

	for y, row := range wordsearch {
		for x, char := range row {
			if char == "X" {
				totalXmasCount += check_all_directions(wordsearch, x, y)
			}
		}
	}

	return totalXmasCount
}

//	 (0,1) (0,3)
//		 |	   |
//
// [  	 V	   V
//
//	[ X, M, A, S ] row 0
//	[ X, A, A, A ] row 1
//	[ X, M, A, M ] row 2
//	[ X, A, A, X ] row 3
//
// ]
func check_all_directions(wordsearch [][]string, x int, y int) int {
	count := 0

	// check to the right
	// is it long enough?
	// check all letters
	row := wordsearch[y]
	if x+3 < len(row) && row[x+1] == "M" && row[x+2] == "A" && row[x+3] == "S" {
		count++
	}

	// check to the left
	// very similar to the right
	// can it go back enough?
	// check all letters
	if x-3 >= 0 && row[x-1] == "M" && row[x-2] == "A" && row[x-3] == "S" {
		count++
	}

	// check down
	// check that there are enough rows!
	if y+3 < len(wordsearch) && wordsearch[y+1][x] == "M" && wordsearch[y+2][x] == "A" && wordsearch[y+3][x] == "S" {
		count++
	}

	// check up
	// check that there are enough rows!
	if y-3 >= 0 && wordsearch[y-1][x] == "M" && wordsearch[y-2][x] == "A" && wordsearch[y-3][x] == "S" {
		count++
	}

	// check diagonal right + down
	if x+3 < len(row) && y+3 < len(wordsearch) && wordsearch[y+1][x+1] == "M" && wordsearch[y+2][x+2] == "A" && wordsearch[y+3][x+3] == "S" {
		count++
	}

	// check diagonal right + up
	if x+3 < len(row) && y-3 >= 0 && wordsearch[y-1][x+1] == "M" && wordsearch[y-2][x+2] == "A" && wordsearch[y-3][x+3] == "S" {
		count++
	}

	// check diagonal left + up
	if x-3 >= 0 && y-3 >= 0 && wordsearch[y-1][x-1] == "M" && wordsearch[y-2][x-2] == "A" && wordsearch[y-3][x-3] == "S" {
		count++
	}

	// check diagonal left + down
	if x-3 >= 0 && y+3 < len(wordsearch) && wordsearch[y+1][x-1] == "M" && wordsearch[y+2][x-2] == "A" && wordsearch[y+3][x-3] == "S" {
		count++
	}

	return count
}

func Part_2_FindAllX_mas(input string) int {

	wordsearch := parse_into_array(input)

	totalXmasCount := 0

	for y, row := range wordsearch {
		for x, char := range row {
			if char == "A" {
				if is_x_mas(wordsearch, x, y) {
					totalXmasCount++
				}
			}
		}
	}

	return totalXmasCount
}

func is_x_mas(wordsearch [][]string, x int, y int) bool {
	// false if we cant go right or left
	if x+1 >= len(wordsearch[0]) || x-1 < 0 {
		return false
	}

	// false if we cant go up or down
	if y+1 >= len(wordsearch) || y-1 < 0 {
		return false
	}

	// upper left, lower right
	first_mas := wordsearch[y-1][x-1] + wordsearch[y][x] + wordsearch[y+1][x+1]
	// upper right, lower left
	second_mas := wordsearch[y-1][x+1] + wordsearch[y][x] + wordsearch[y+1][x-1]

	return (first_mas == "MAS" || first_mas == "SAM") && (second_mas == "MAS" || second_mas == "SAM")

}
