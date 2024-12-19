package day5

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Part_1_validate_print_queue(t *testing.T) {

	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

	result := Part_1_validate_print_queue(input)

	assert.Equal(t, result, 143)
}

func Test_Part_2_incorrect_only_updates(t *testing.T) {

	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

	result := Part_2_incorrect_only_updates(input)

	assert.Equal(t, result, 123)
}
