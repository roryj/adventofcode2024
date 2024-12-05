package day3

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ParseStart = iota
	ParseArguments
	ParseStartBracket
	ParseLeftInput
	ParseInputSeperation
	ParseRightInput
	ParseEndBracket
)

const (
	MultiplyInstruction = "mul"
)

type Instruction struct {
	instructionType string
	arguments       []int
}

func (i *Instruction) addArgument(arg int) {
	i.arguments = append(i.arguments, arg)
}

func (i *Instruction) run() int {
	result := 1
	switch i.instructionType {
	case MultiplyInstruction:
		for _, arg := range i.arguments {
			result *= arg
		}
	}

	return result
}

func Part_1_parse_and_run_corrupted(input string) int {

	// we get a single line input, and go through character by character
	// building an operation to run and then adding it to the total.
	// Anytime we run into any invalid characters, we dump anything we've seen
	// and keep going

	result := 0

	currentInstruction := Instruction{}
	currentState := ParseStart
	var currentInput []string
	done := false

	count := 0

	fmt.Printf("INPUT: %v\n", input)

	for !done && count < len(input) {
		r := input[count]
		currChar := string(r)
		fmt.Printf("Currently @ %v\n", count)
		switch currentState {
		case ParseStart:
			// if we are at the start, then we need to check the next three characters for a valid
			// instruction
			if count+3 >= len(input) {
				// ran out of room
				// we are done
				done = true
				break
			}
			possibleInstruction := input[count : count+3]
			fmt.Printf("[%d/%d] instruction? -> %v\n", count+1, len(input), possibleInstruction)

			// TODO: add better check if we need more instructions
			if possibleInstruction != MultiplyInstruction {
				fmt.Println("invalid instruction for ParseStart. NEXT!")
				// also clear the queue and reset the instruction
				currentState = ParseStart
				currentInstruction = Instruction{}
				count++
				continue
			}

			currentInstruction.instructionType = MultiplyInstruction
			currentState = ParseStartBracket
			count += 3
		case ParseStartBracket:
			if currChar != "(" {
				fmt.Println("invalid instruction in ParseStartBracket. NEXT!")
				// also clear the queue and reset the instruction
				currentInstruction = Instruction{}
				currentState = ParseStart
				count++
				continue
			}

			currentState = ParseArguments
			count++
		case ParseArguments:
			// if we have a comma, add the last argument to the list
			if currChar == "," {
				numStr := strings.Join(currentInput, "")
				currentInput = nil
				argument, err := strconv.Atoi(numStr)
				if err != nil {
					fmt.Printf("invalid number argument: %v\n", err)
					currentInstruction = Instruction{}
					currentState = ParseStart
					continue
				}
				currentInstruction.addArgument(argument)
				count++
				continue
			}

			if currChar == ")" {
				numStr := strings.Join(currentInput, "")
				currentInput = nil
				argument, err := strconv.Atoi(numStr)
				if err != nil {
					fmt.Printf("invalid number argument: %v\n", err)
					currentInstruction = Instruction{}
					currentState = ParseStart
					count++
					continue
				}
				currentInstruction.addArgument(argument)
				if len(currentInstruction.arguments) != 2 {
					fmt.Printf("invalid number argument: %v\n", currentInstruction)
					currentInstruction = Instruction{}
					currentState = ParseStart
					count++
					continue
				}

				fmt.Printf("Valid instruction! Running: %v\n", currentInstruction)
				result += currentInstruction.run()
				currentInstruction = Instruction{}
				currentState = ParseStart
				count++
				continue
			}

			// try parsing the char to a number
			_, err := strconv.Atoi(currChar)
			// if it is not a valid number, then we are done here
			if err != nil {
				fmt.Println("invalid number for argument")
				currentInstruction = Instruction{}
				currentState = ParseStart
				currentInput = nil
				count++
				continue
			}

			// otherwise add it to the array to process and keep going
			currentInput = append(currentInput, currChar)
			count++

		default:
			fmt.Println("unhandled instruction")
			count++
		}
	}

	return result
}
