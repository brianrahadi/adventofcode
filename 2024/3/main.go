// regex tested from https://regex101.com/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|`)
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	fmt.Println(sum)
}

func part2(lines []string) {
	sum := 0
	shouldDo := true
	for _, line := range lines {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				shouldDo = true
			} else if match[0] == "don't()" {
				shouldDo = false
			} else if shouldDo {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				sum += num1 * num2
			}
		}
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
