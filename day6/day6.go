package day6

import (
	"fmt"
	"strings"

	"roryj.ca/aoc2024/helpers"
)

type MapEntry int

const (
	Obstruction    MapEntry = 0
	Empty          MapEntry = 1
	NewObstruction MapEntry = 2
)

type Direction int

const (
	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

type MapPoint struct {
	t       MapEntry
	visited bool
}

func (m *MapPoint) visit() {
	m.visited = true
}

type Point struct {
	x int
	y int
}

type Guard struct {
	location    Point
	direction   Direction
	outOfBounds bool
}

func (g *Guard) deepClone() Guard {
	return Guard{
		location: Point{
			x: g.location.x,
			y: g.location.y,
		},
		direction:   g.direction,
		outOfBounds: g.outOfBounds,
	}
}

func (g *Guard) nextDirection() Direction {
	switch g.direction {
	case Up:
		return Right
	case Down:
		return Left
	case Right:
		return Down
	case Left:
		return Up
	default:
		panic("invalid direction change")
	}
}

func (g *Guard) changeDirection() {
	g.direction = g.nextDirection()
}

type Map struct {
	internalMap [][]MapPoint
	guard       Guard
}

func (m *Map) deepClone() Map {
	clone := Map{}
	clone.guard = m.guard.deepClone()
	clone.internalMap = [][]MapPoint{}
	for _, row := range m.internalMap {
		clonedRow := []MapPoint{}
		clearedRow := helpers.Map(row, func(p MapPoint) MapPoint {
			return MapPoint{
				t:       p.t,
				visited: false,
			}
		})

		clonedRow = append(clonedRow, clearedRow...)
		clone.internalMap = append(clone.internalMap, clonedRow)
	}

	return clone
}

func (m *Map) findLoops() {
	for _, row := range m.internalMap {
		for _, p := range row {
			if p.t == Empty {
				// empty space, no need
				continue
			}

			// checkForSquare()

		}
	}
}

func checkForLoop(m *Map) bool {

	m.internalMap[m.guard.location.y][m.guard.location.x].visit()

	// given a point, lets keep going and see if we get a loop, which we can calculate
	// by hitting the same obstacle in the same direction

	visitedLocation := map[string]bool{}

	for !m.GuardIsOutOfBounds() {
		// direction := m.guard.direction
		point := m.guard.location
		m.tick()
		// fmt.Println("we've moved")
		// m.debug()
		key := fmt.Sprintf("%v-%v", m.guard.location, m.guard.direction)
		// if we are at the same point, then we have turned. Lets see if we have gone the same
		// way and direction here
		if m.guard.location == point {
			_, ok := visitedLocation[key]
			if ok {
				// fmt.Println("ITS A LOOP!")
				// m.debug()
				return true
			}
		}

		visitedLocation[key] = true
	}

	return false
}

func (m *Map) visitedLocationCount() int {
	result := 0
	for _, row := range m.internalMap {
		for _, p := range row {
			if p.visited {
				result++
			}
		}
	}

	return result
}

// returns the next position of the guard. If the guard goes
// out of bounds this will return false
func (m *Map) getNextPosition() (Point, Direction, bool) {

	// if we are out of bounds, do nothing
	if m.GuardIsOutOfBounds() {
		return Point{}, Down, false
	}

	nextX := m.guard.location.x
	nextY := m.guard.location.y

	switch m.guard.direction {
	case Up:
		nextY--
	case Down:
		nextY++
	case Left:
		nextX--
	case Right:
		nextX++
	}

	outOfBounds := (nextX >= len(m.internalMap[0]) || nextX < 0 || nextY >= len(m.internalMap) || nextY < 0)
	direction := m.guard.direction

	if !outOfBounds {
		nextPoint := m.internalMap[nextY][nextX]
		if nextPoint.t == Obstruction || nextPoint.t == NewObstruction {
			direction = m.guard.nextDirection()
		}
	}

	return Point{x: nextX, y: nextY}, direction, !outOfBounds
}

func (m *Map) tick() {

	nextLocation, newDirection, ok := m.getNextPosition()
	if !ok {
		m.guard.outOfBounds = true
		m.guard.location.x = nextLocation.x
		m.guard.location.y = nextLocation.y
		return
	}

	if m.guard.direction != newDirection {
		m.guard.changeDirection()
		return
	}

	nextPoint := m.internalMap[nextLocation.y][nextLocation.x]

	nextPoint.visit()
	m.internalMap[nextLocation.y][nextLocation.x] = nextPoint

	m.guard.location.x = nextLocation.x
	m.guard.location.y = nextLocation.y
}

func (m *Map) GuardIsOutOfBounds() bool {
	return m.guard.outOfBounds
}

func (m *Map) debug() {
	for y, row := range m.internalMap {
		for x, p := range row {

			if m.guard.location.x == x && m.guard.location.y == y {
				// print out the guard based on their direction
				switch m.guard.direction {
				case Up:
					fmt.Print("^")
				case Down:
					fmt.Print("v")
				case Left:
					fmt.Print("<")
				case Right:
					fmt.Print(">")
				default:
					panic("invalid guard direction")
				}
				continue
			}

			// else we just print the basics
			switch p.t {
			case Obstruction:
				fmt.Print("#")
			case Empty:
				if p.visited {
					fmt.Print("X")
				} else {
					fmt.Print(".")
				}
			case NewObstruction:
				fmt.Printf("O")
			default:
				panic("invalid type")
			}
		}
		fmt.Println()
	}
}

func parse_input(input string) Map {
	internalMap := [][]MapPoint{}
	var guard Guard

	for y, line := range strings.Split(input, "\n") {
		row := []MapPoint{}

		for x, c := range line {
			switch c {
			case '.':
				// empty space
				row = append(row, MapPoint{t: Empty, visited: false})
			case '#':
				// block
				row = append(row, MapPoint{t: Obstruction, visited: false})
			case '^':
				row = append(row, MapPoint{t: Empty, visited: true})
				guard = Guard{
					location:  Point{x: x, y: y},
					direction: Up,
				}
			case '>':
				row = append(row, MapPoint{t: Empty, visited: true})
				guard = Guard{
					location:  Point{x: x, y: y},
					direction: Right,
				}
			case '<':
				row = append(row, MapPoint{t: Empty, visited: true})
				guard = Guard{
					location:  Point{x: x, y: y},
					direction: Left,
				}
			case 'v':
				row = append(row, MapPoint{t: Empty, visited: true})
				guard = Guard{
					location:  Point{x: x, y: y},
					direction: Down,
				}
			}
		}
		internalMap = append(internalMap, row)
	}

	return Map{
		internalMap: internalMap,
		guard:       guard,
	}
}

func Part_1_GuardPatrol(input string) int {

	m := parse_input(input)

	count := 0
	for !m.GuardIsOutOfBounds() {
		m.tick()
		// fmt.Printf("--- tick #%d ---\n", count)
		// m.debug()
		// fmt.Println()
		count++
	}

	return m.visitedLocationCount()
}

func Part_2_LoopGuard(input string) int {
	m := parse_input(input)

	loops := 0

	for !m.GuardIsOutOfBounds() {
		// fmt.Println("guard is not out of bounds, continuing")
		nextPosition, _, inBounds := m.getNextPosition()

		if !inBounds || nextPosition == m.guard.location {
			// in the same position, we can skip here
			m.tick()
			continue
		}

		nextPositionEntry := m.internalMap[nextPosition.y][nextPosition.x]
		if nextPositionEntry.t == Obstruction {
			// we already have an obstruction, cant do much here
			m.tick()
			continue
		}

		// fmt.Printf("checking for a loop when setting the position %v to an obstruction\n", nextPosition)
		clonedMap := m.deepClone()
		clonedMap.internalMap[nextPosition.y][nextPosition.x] = MapPoint{t: NewObstruction, visited: nextPositionEntry.visited}
		// clonedMap.debug()
		if checkForLoop(&clonedMap) {
			loops++
		}

		// lets keep going!
		m.tick()
	}

	return loops
}
