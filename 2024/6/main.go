package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

type Coordinate struct {
	x int
	y int
}

func moveForward(c rune) Coordinate {
	switch c {
	case '^':
		return Coordinate{-1, 0}
	case '>':
		return Coordinate{0, 1}
	case 'V':
		return Coordinate{1, 0}
	case '<':
		return Coordinate{0, -1}
	}

	// error here
	fmt.Println("ERROR MOVE FORWARD")
	return Coordinate{0, 0}
}

func switchDirection(c rune) rune {
	switch c {
	case '^':
		return '>'
	case '>':
		return 'V'
	case 'V':
		return '<'
	case '<':
		return '^'
	}

	// error here
	fmt.Println("ERROR SWITCHING DIRECTION")
	return '^'
}

func countPath(matrix [][]rune) int {

	visited := map[Coordinate]bool{}
	visitedCount := 0
	gPos := Coordinate{0, 0}
	gDir := '^'

	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			if matrix[i][j] != '.' && matrix[i][j] != '#' {
				gPos = Coordinate{i, j}
				gDir = matrix[i][j]
				matrix[i][j] = '.'
				break
			}
		}
	}

	for {
		x, y := gPos.x, gPos.y

		if !visited[Coordinate{x, y}] {
			visited[Coordinate{x, y}] = true
			visitedCount += 1
		}

		movedForward := moveForward(gDir)
		newX, newY := x+movedForward.x, y+movedForward.y

		if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[newX]) {
			break
		}

		if matrix[newX][newY] == '#' {
			gDir = switchDirection(gDir)
			continue
		}

		gPos = Coordinate{newX, newY}

	}
	return visitedCount
}

func part1(lines []string) {
	matrix := [][]rune{}
	for _, line := range lines {
		fields := strings.Split(line, "")
		runes := lo.Map(fields, func(str string, _ int) rune { return rune(str[0]) })
		matrix = append(matrix, runes)
	}

	fmt.Println(countPath(matrix))
}

func hasObstacleCycle(matrix [][]rune) bool {
	visited := map[Coordinate]bool{}

	visitedCount := 0
	gPos := Coordinate{0, 0}
	gDir := '^'
	visitedState := map[string]bool{}
	visitedState[fmt.Sprintf("%d,%d,%d", gDir, gPos.x, gPos.y)] = true

	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			if matrix[i][j] != '.' && matrix[i][j] != '#' {
				gPos = Coordinate{i, j}
				gDir = matrix[i][j]
				break
			}
		}
	}

	for {
		x, y := gPos.x, gPos.y

		if visitedState[fmt.Sprintf("%d,%d,%d", gDir, gPos.x, gPos.y)] {
			return true
		}
		visitedState[fmt.Sprintf("%d,%d,%d", gDir, gPos.x, gPos.y)] = true

		if !visited[Coordinate{x, y}] {
			visited[Coordinate{x, y}] = true
			visitedCount += 1
		}

		movedForward := moveForward(gDir)
		newX, newY := x+movedForward.x, y+movedForward.y

		if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[newX]) {
			break
		}

		if matrix[newX][newY] == '#' {
			gDir = switchDirection(gDir)
			continue
		}

		gPos = Coordinate{newX, newY}

	}

	return false
}

func part2(lines []string) {
	matrix := [][]rune{}
	for _, line := range lines {
		fields := strings.Split(line, "")
		runes := lo.Map(fields, func(str string, _ int) rune { return rune(str[0]) })
		matrix = append(matrix, runes)
	}

	sumObstacleCycle := 0

	for r := range len(matrix) {
		for c := range len(matrix[r]) {
			if matrix[r][c] == '.' {
				matrix[r][c] = '#'
				if hasObstacleCycle(matrix) {
					sumObstacleCycle += 1
				}
				matrix[r][c] = '.'
			}
		}
	}

	fmt.Println(sumObstacleCycle)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println()
	fmt.Println("part 1")
	part1(lines)

	fmt.Println()
	fmt.Println("part 2")
	part2(lines)
}
