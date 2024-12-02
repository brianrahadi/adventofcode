package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(lines []string) {
	leftList := make([]int, 0, len(lines))
	rightList := make([]int, 0, len(lines))
	for _, line := range lines {
		strs := strings.Fields(line)

		num1, _ := strconv.Atoi(strs[0])
		num2, _ := strconv.Atoi(strs[1])

		leftList = append(leftList, num1)
		rightList = append(rightList, num2)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	res := 0

	for i := range leftList {
		maxNum := max(leftList[i], rightList[i])
		minNum := min(leftList[i], rightList[i])
		res += maxNum - minNum
	}

	fmt.Println(res)
}

func part2(lines []string) {
	leftList := make([]int, 0, len(lines))
	rightCounter := map[int]int{}
	for _, line := range lines {
		strs := strings.Fields(line)

		num1, _ := strconv.Atoi(strs[0])
		num2, _ := strconv.Atoi(strs[1])

		leftList = append(leftList, num1)
		rightCounter[num2]++
	}

	res := 0

	for _, num := range leftList {
		res += num * rightCounter[num]
	}

	fmt.Println(res)
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
