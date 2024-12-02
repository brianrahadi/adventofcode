package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/samber/lo"
)

func isSymbolAdjacent(matrix [][]rune, r int, c int) bool {
	adj := [8][2]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}

	for _, val := range adj {
		ax, ay := r+val[0], c+val[1]

		if ax < 0 || ax >= len(matrix) || ay < 0 || ay >= len(matrix) {
			continue
		}

		if !unicode.IsDigit(matrix[ax][ay]) && matrix[ax][ay] != '.' {
			// fmt.Println("FOUND ", string(matrix[ax][ay]))
			return true
		}
	}
	return false
}

func getNumber(matrix [][]rune, r int, c int) (int, string) {
	left, right := c, c
	for left >= 1 {
		if matrix[r][left-1] >= '0' && matrix[r][left-1] <= '9' {
			left -= 1
		} else {
			break
		}
	}

	for right <= len(matrix[r])-2 {
		if matrix[r][right+1] >= '0' && matrix[r][right+1] <= '9' {
			right += 1
		} else {
			break
		}
	}
	str := ""
	for i := left; i <= right; i++ {
		str += string(matrix[r][i])
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	key := fmt.Sprintf("%d,%d,%d", r, left, right)
	return num, key
}

func part1(lines []string) {
	matrix := [][]rune{}
	for _, line := range lines {
		splittedLine := strings.Split(line, "")
		runes := lo.Map(splittedLine, func(s string, _ int) rune { return rune(s[0]) })
		matrix = append(matrix, runes)
	}

	sum := 0
	coordinateSet := map[string]bool{}

	for r := range matrix {
		for c := range matrix[r] {
			if unicode.IsDigit(matrix[r][c]) && isSymbolAdjacent(matrix, r, c) {
				num, key := getNumber(matrix, r, c)
				if !coordinateSet[key] {
					// fmt.Println(num)
					sum += num
				}
				coordinateSet[key] = true
			}

		}
	}

	fmt.Println(sum)
}

func part2(lines []string) {
	for _, line := range lines {
		_ = line
	}

	fmt.Println("answer2")
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
