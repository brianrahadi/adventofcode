package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func part1(lines []string) {
	sum := 0
	for i, line := range lines {
		line = strings.Replace(line, fmt.Sprintf("Card %d:", i+1), "", 1)
		splittedLine := strings.SplitN(line, "|", 2)
		winningNumbers := lo.Map(strings.Fields(splittedLine[0]), func(numStr string, _ int) int {
			num, _ := strconv.Atoi(numStr)
			return num
		})

		winningNumbersSet := map[int]bool{}
		for _, num := range winningNumbers {
			winningNumbersSet[num] = true
		}

		playingNumbers := lo.Map(strings.Fields(splittedLine[1]), func(numStr string, _ int) int {
			num, _ := strconv.Atoi(numStr)
			return num
		})

		acc := 0
		for _, num := range playingNumbers {
			if !winningNumbersSet[num] {
				continue
			}
			if acc == 0 {
				acc = 1
			} else {
				acc *= 2
			}
		}
		sum += acc
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	sum := 0
	cardCounter := map[int]int{}
	for i, line := range lines {
		line = strings.Replace(line, fmt.Sprintf("Card %d:", i+1), "", 1)
		splittedLine := strings.SplitN(line, "|", 2)
		winningNumbers := lo.Map(strings.Fields(splittedLine[0]), func(numStr string, _ int) int {
			num, _ := strconv.Atoi(numStr)
			return num
		})

		winningNumbersSet := map[int]bool{}
		for _, num := range winningNumbers {
			winningNumbersSet[num] = true
		}

		playingNumbers := lo.Map(strings.Fields(splittedLine[1]), func(numStr string, _ int) int {
			num, _ := strconv.Atoi(numStr)
			return num
		})

		currentCardIndex := i + 1
		cardCounter[currentCardIndex] += 1 // starts with 1

		addedCardIndex := i + 2 // after i + 1
		for _, num := range playingNumbers {
			if !winningNumbersSet[num] {
				continue
			}
			cardCounter[addedCardIndex] += cardCounter[currentCardIndex]
			addedCardIndex++
		}
		sum += cardCounter[i+1]
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
