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
