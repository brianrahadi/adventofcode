package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

func checkXmasCount(matrix [][]rune, r int, c int) int {
	count := 0

	dirs := [8][2]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}

	for _, dir := range dirs {
		ar, ac := r, c
		for i := range 4 {
			if ar < 0 || ar >= len(matrix) || ac < 0 || ac >= len(matrix[ar]) {
				break
			}
			if i == 0 && matrix[ar][ac] != 'X' {
				break
			}
			if i == 1 && matrix[ar][ac] != 'M' {
				break
			}
			if i == 2 && matrix[ar][ac] != 'A' {
				break
			}
			if i == 3 && matrix[ar][ac] == 'S' {
				count += 1
			}
			ar += dir[0]
			ac += dir[1]
		}
	}

	return count
}

func part1(lines []string) {
	matrix := [][]rune{}
	for _, line := range lines {
		fields := strings.Split(line, "")
		runes := lo.Map(fields, func(str string, _ int) rune { return rune(str[0]) })
		matrix = append(matrix, runes)
	}

	count := 0

	for r := range matrix {
		for c := range matrix[r] {
			count += checkXmasCount(matrix, r, c)
		}
	}

	fmt.Println(count)
}

// M.S
// .A.
// M.S
func checkIsMAS(matrix [][]rune, r int, c int) bool {
	if matrix[r][c] != 'A' {
		return false
	}

	dirs := [5][2]int{
		{-1, 1},
		{1, -1},
		{1, 1},
		{-1, -1},
	}

	previouslyM := true

	for i, dir := range dirs {
		ar, ac := r+dir[0], c+dir[1]
		if ar < 0 || ar >= len(matrix) || ac < 0 || ac >= len(matrix[ar]) {
			break
		}
		if (i == 0 || i == 2) && (matrix[ar][ac] == 'M' || matrix[ar][ac] == 'S') {
			if matrix[ar][ac] == 'M' {
				previouslyM = true
			} else {
				previouslyM = false
			}
		} else if (i == 1 || i == 3) && (matrix[ar][ac] == 'M' || matrix[ar][ac] == 'S') {
			if previouslyM && matrix[ar][ac] == 'S' || (!previouslyM && matrix[ar][ac] == 'M') {
				if i == 3 {
					return true
				}
			} else {
				break
			}
		} else {
			break
		}
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

	count := 0

	for r := range matrix {
		for c := range matrix[r] {
			if checkIsMAS(matrix, r, c) {
				count++
			}
		}
	}

	fmt.Println(count)
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
