package day6

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Part_1_GuardPatrol(t *testing.T) {

	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	result := Part_1_GuardPatrol(input)

	assert.Equal(t, result, 41)
}

func Test_Part_2_LoopGuard(t *testing.T) {

	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	result := Part_2_LoopGuard(input)

	assert.Equal(t, result, 6)
}

func Test_Part_2_LoopGuard_other(t *testing.T) {

	input := `#...#.....
^........#
.#........
..#.......
.......#..
..........
.#........
........#.
#.........
......#...`

	result := Part_2_LoopGuard(input)

	assert.Equal(t, result, 6)
}
