package day4

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Part_1_FindAllXmasSimple(t *testing.T) {

	input := `XMAS
XAAA
XMAM
XAAX`

	result := Part_1_FindAllXmas(input)
	assert.Equal(t, result, 3)
}

func Test_Part_1_FindAllXmas(t *testing.T) {

	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	result := Part_1_FindAllXmas(input)
	assert.Equal(t, result, 18)
}

func Test_Part_2_FindAllX_mas(t *testing.T) {

	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	result := Part_2_FindAllX_mas(input)
	assert.Equal(t, result, 9)
}
