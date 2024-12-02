package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		nums := make([]int, 0, len(numStrs))
		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

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
		arrayWithDeletedIndex := append([]int{}, nums[:i]...)
		arrayWithDeletedIndex = append(arrayWithDeletedIndex, nums[i+1:]...)

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
		nums := make([]int, 0, len(numStrs))
		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

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
