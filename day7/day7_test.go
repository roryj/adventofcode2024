package day7

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Part_1_CalibrateEquations(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	resultOne, resultTwo := Part_1_CalibrateEquations(input)
	assert.Equal(t, 3749, resultOne)
	assert.Equal(t, 11387, resultTwo)
}
