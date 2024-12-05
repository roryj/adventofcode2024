package day3

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Part_1_parse_and_run_corrupted(t *testing.T) {

	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := Part_1_parse_and_run_corrupted(input)

	assert.Equal(t, result, 100)
}
