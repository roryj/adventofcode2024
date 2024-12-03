package day1

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Part1_FindDifferenceDistanceBetweenLists(t *testing.T) {

	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	result := Part1_FindDifferenceDistanceBetweenLists(input)

	assert.Equal(t, result, 11)
}

func Test_Part2_FindSimilarity(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	result := Part2_FindSimilarity(input)

	assert.Equal(t, result, 31)
}
