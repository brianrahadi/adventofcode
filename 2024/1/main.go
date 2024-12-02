package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for _, line := range lines {
		strs := strings.Split(line, "   ")

		num1, _ := strconv.Atoi(strs[0])
		num2, _ := strconv.Atoi(strs[1])

		leftList = append(leftList, num1)
		rightList = append(rightList, num2)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	res := 0

	for i := 0; i < len(leftList); i++ {
		maxNum := max(leftList[i], rightList[i])
		minNum := min(leftList[i], rightList[i])
		res += maxNum - minNum
	}

	return res
}

func part2(lines []string) int {
	leftList := make([]int, 0)
	rightCounter := map[int]int{}

	for _, line := range lines {
		strs := strings.Split(line, "   ")

		num1, _ := strconv.Atoi(strs[0])
		num2, _ := strconv.Atoi(strs[1])

		leftList = append(leftList, num1)
		rightCounter[num2]++
	}

	res := 0

	for _, num := range leftList {
		res += num * rightCounter[num]
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1")
	fmt.Println(part1(lines))

	fmt.Println("Part 2")
	fmt.Println(part2(lines))
}
