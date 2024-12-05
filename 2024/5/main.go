package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func part1(lines []string) {
	isBefore := map[int][]int{}

	isBeforeOnePass := map[int]map[int]bool{}
	sum := 0

	isFirstSection := true
	for _, line := range lines {
		if len(line) == 0 {
			isFirstSection = false
			for key, values := range isBefore {
				isBeforeOnePass[key] = make(map[int]bool)
				for _, value := range values {
					isBeforeOnePass[key][value] = true
				}
			}
			continue
		}
		if isFirstSection {
			intStrs := strings.Split(line, "|")
			num1, _ := strconv.Atoi(intStrs[0])
			num2, _ := strconv.Atoi(intStrs[1])

			arr, ok := isBefore[num1]
			if ok {
				isBefore[num1] = append(arr, num2)
			} else {
				isBefore[num1] = []int{num2}
			}
		} else {
			numStrs := strings.Split(line, ",")
			nums := lo.Map(numStrs, func(str string, _ int) int {
				num, _ := strconv.Atoi(str)
				return num
			})
			isValid := true
			for i := range len(nums) - 1 {
				if !isBeforeOnePass[nums[i]][nums[i+1]] {
					isValid = false
					break
				}
			}

			if isValid {
				sum += nums[len(nums)/2]
			}
		}
	}

	fmt.Println(sum)
}

func part2(lines []string) {
	isBefore := map[int][]int{}

	isBeforeOnePass := map[int]map[int]bool{}
	sum := 0

	isFirstSection := true
	for _, line := range lines {
		if len(line) == 0 {
			isFirstSection = false
			for key, values := range isBefore {
				isBeforeOnePass[key] = make(map[int]bool)
				for _, value := range values {
					isBeforeOnePass[key][value] = true
				}
			}
			continue
		}
		if isFirstSection {
			intStrs := strings.Split(line, "|")
			num1, _ := strconv.Atoi(intStrs[0])
			num2, _ := strconv.Atoi(intStrs[1])

			arr, ok := isBefore[num1]
			if ok {
				isBefore[num1] = append(arr, num2)
			} else {
				isBefore[num1] = []int{num2}
			}
		} else {
			numStrs := strings.Split(line, ",")
			nums := lo.Map(numStrs, func(str string, _ int) int {
				num, _ := strconv.Atoi(str)
				return num
			})
			isValid := true
			for i := range len(nums) - 1 {
				if !isBeforeOnePass[nums[i]][nums[i+1]] {
					isValid = false
					break
				}
			}

			if !isValid {
				slices.SortFunc(nums, func(num1 int, num2 int) int {
					if isBeforeOnePass[num1][num2] {
						return 1
					}
					return -1
				})
				sum += nums[len(nums)/2]
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
