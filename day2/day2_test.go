package day2

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Part_1_CalculateSafeReports(t *testing.T) {

	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	result := Part_1_CalculateSafeReports(input)

	assert.Equal(t, result, 2)
}

func Test_Part_2_FindSimilarity(t *testing.T) {

	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	result := Part_2_CalculateSafeReportsWithDampener(input)

	assert.Equal(t, result, 4)
}

func Test_Part_2_FindSimilarity_Extra(t *testing.T) {

	input := `1 3 2 4
3 1 2 3 4
3 1 4 5 6
3 1 5 6 7
1 2 3 4 5 199 8
100 97 600 94 93 92
10 9 8 7 6 5 1`

	result := Part_2_CalculateSafeReportsWithDampener(input)

	assert.Equal(t, result, 7)
}

func Test_Part_2_FindSimilarity_RedditCornerCases(t *testing.T) {
	input := `48 46 47 49 51 54 56
1 1 2 3 4 5
1 2 3 4 5 5
5 1 2 3 4 5
1 4 3 2 1
1 6 7 8 9
1 2 3 4 3
9 8 7 6 7
7 10 8 10 11
29 28 27 25 26 25 22 20`

	result := Part_2_CalculateSafeReportsWithDampener(input)

	assert.Equal(t, result, 10)
}

func Test_Part_2_FindSimilarity_OutOfBounds(t *testing.T) {
	input := `45 45 10 10
49 48 47 40 33`

	result := Part_2_CalculateSafeReportsWithDampener(input)

	assert.Equal(t, result, 0)
}
