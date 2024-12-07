package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func isPossibleCalibration(answer int, numbers []int, isPart2 bool) bool {
	if len(numbers) > 2 {
		summedArr := append([]int{numbers[0] + numbers[1]}, numbers[2:]...)
		multipliedArr := append([]int{numbers[0] * numbers[1]}, numbers[2:]...)

		if !isPart2 {
			return isPossibleCalibration(answer, summedArr, isPart2) || isPossibleCalibration(answer, multipliedArr, isPart2)
		}
		concatenatedNum, _ := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1]))
		concatenatedArr := append([]int{concatenatedNum}, numbers[2:]...)
		return isPossibleCalibration(answer, summedArr, isPart2) ||
			isPossibleCalibration(answer, multipliedArr, isPart2) ||
			isPossibleCalibration(answer, concatenatedArr, isPart2)
	} else if len(numbers) == 2 {
		if !isPart2 {
			return answer == numbers[0]+numbers[1] || answer == numbers[0]*numbers[1]
		}
		concatenatedNum, _ := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1]))
		return answer == numbers[0]+numbers[1] || answer == numbers[0]*numbers[1] || answer == concatenatedNum
	}
	fmt.Println("ERROR ARRAY LENGTH")
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sumP1 := 0
	sumP2 := 0
	for _, line := range lines {
		splits := strings.SplitN(line, ": ", 2)
		answer, _ := strconv.Atoi(splits[0])
		numbers := lo.Map(strings.Fields(splits[1]), func(str string, _ int) int {
			num, _ := strconv.Atoi(str)
			return num
		})
		if isPossibleCalibration(answer, numbers, false) {
			sumP1 += answer
			sumP2 += answer
		} else if isPossibleCalibration(answer, numbers, true) {
			sumP2 += answer
		}
	}

	fmt.Println("\npart1")
	fmt.Println(sumP1)
	fmt.Println("\npart2")
	fmt.Println(sumP2)
}
