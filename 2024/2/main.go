package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func isSafeArray(nums []int) bool {
	isIncreasing := nums[0] < nums[1]

	for i := range len(nums) - 1 {
		if isIncreasing && (nums[i+1] <= nums[i] || nums[i+1] > nums[i]+3) {
			return false
		}
		if !isIncreasing && (nums[i+1] >= nums[i] || nums[i+1] < nums[i]-3) {
			return false
		}
	}

	return true
}

func part1(lines []string) {
	safeCount := 0
	for _, line := range lines {
		numStrs := strings.Fields(line)
		nums := lo.Map(numStrs, func(numStr string, _ int) int {
			num, _ := strconv.Atoi(numStr)
			return num
		})
		if isSafeArray(nums) {
			safeCount += 1
		}
	}

	fmt.Println(safeCount)
}

func isSafeArrayPt2(nums []int) bool {
	if isSafeArray(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		arrayWithDeletedIndex := lo.Filter(nums, func(_ int, index int) bool { return index != i })

		if isSafeArray(arrayWithDeletedIndex) {
			return true
		}
	}

	return false
}

func part2(lines []string) {
	safeCount := 0
	for _, line := range lines {
		numStrs := strings.Fields(line)
		nums := lo.Map(numStrs, func(numStr string, _ int) int {
			num, _ := strconv.Atoi(numStr)
			return num
		})

		if isSafeArrayPt2(nums) {
			safeCount += 1
		}
	}

	fmt.Println(safeCount)
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
