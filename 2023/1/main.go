package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		runes := []rune(line)
		numStr := ""
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				numStr = string(runes[i])
				break
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				numStr = numStr + string(runes[i])
				break
			}
		}

		num, _ := strconv.Atoi(numStr)
		sum += num
	}

	fmt.Println(sum)
}

func firstDigit(str string) int {
	acc := ""

	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			return int(str[i] - '0')
		}

		acc += string(str[i])

		for i, d := range digits {
			if strings.Contains(acc, d) {
				return i + 1
			}
		}
	}
	println("ERROR FIRST DIGIT")
	return 0
}

func lastDigit(str string) int {
	acc := ""

	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] >= '0' && str[i] <= '9' {
			return int(str[i] - '0')
		}

		acc = string(str[i]) + acc

		for i, d := range digits {
			if strings.Contains(acc, d) {
				return i + 1
			}
		}
	}
	println("ERROR LAST DIGIT")
	return 0
}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		num := firstDigit(line)*10 + lastDigit(line)
		sum += num
	}

	fmt.Println(sum)
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
