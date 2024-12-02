package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func subsetsPass(subsets []string) bool {
	cubeCounterThreshold := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, subset := range subsets {
		cubeCounter := map[string]int{}

		strs := strings.Split(subset, ", ") // "5 Red"
		for _, s := range strs {
			parts := strings.SplitN(s, " ", 2)
			count, _ := strconv.Atoi(parts[0])
			color := parts[1]
			cubeCounter[color] = count
		}
		for colour, count := range cubeCounter {
			if count > cubeCounterThreshold[colour] {
				return false
			}
		}
	}

	return true
}
func part1(lines []string) {
	sum := 0
	for i, line := range lines {
		_ = line
		removed := fmt.Sprintf("Game %d: ", i+1)
		updatedLine := strings.Replace(line, removed, "", 1)
		subsets := strings.Split(updatedLine, "; ")
		if subsetsPass(subsets) {
			sum += i + 1
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
