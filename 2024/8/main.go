package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

type Point struct {
	r int
	c int
}

func part1(lines []string) {
	matrix := [][]rune{}
	for _, line := range lines {
		splittedLine := strings.Split(line, "")
		runes := lo.Map(splittedLine, func(str string, _ int) rune { return rune(str[0]) })
		matrix = append(matrix, runes)
	}

	runePoints := map[rune][]Point{}
	for r := range matrix {
		for c := range matrix[r] {
			point := Point{r, c}
			if matrix[r][c] != '.' {
				runePoints[matrix[r][c]] = append(runePoints[matrix[r][c]], point)
			}
		}
	}

	freqs := make([][]rune, len(matrix))
	for i := range freqs {
		freqs[i] = make([]rune, len(matrix[i]))
	}
	freqCount := 0

	for _, points := range runePoints {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1 := points[i]
				p2 := points[j]
				dist := Point{p2.r - p1.r, p2.c - p1.c}
				freq1 := Point{p1.r - dist.r, p1.c - dist.c}
				freq2 := Point{p2.r + dist.r, p2.c + dist.c}

				if min(freq1.r, freq1.c) >= 0 && max(freq1.r, freq1.c) < len(matrix) && freqs[freq1.r][freq1.c] != 'X' {
					freqs[freq1.r][freq1.c] = 'X'
					freqCount++
				}

				if min(freq2.r, freq2.c) >= 0 && max(freq2.r, freq2.c) < len(matrix[0]) && freqs[freq2.r][freq2.c] != 'X' {
					freqs[freq2.r][freq2.c] = 'X'
					freqCount++
				}

			}
		}
	}

	fmt.Println(freqCount)
}

func part2(lines []string) {
	matrix := [][]rune{}
	for _, line := range lines {
		splittedLine := strings.Split(line, "")
		runes := lo.Map(splittedLine, func(str string, _ int) rune { return rune(str[0]) })
		matrix = append(matrix, runes)
	}

	runePoints := map[rune][]Point{}

	for r := range matrix {
		for c := range matrix[r] {
			point := Point{r, c}
			if matrix[r][c] != '.' {
				runePoints[matrix[r][c]] = append(runePoints[matrix[r][c]], point)
			}
		}
	}

	freqs := make([][]rune, len(matrix))
	for i := range freqs {
		freqs[i] = make([]rune, len(matrix[i]))
	}

	for _, points := range runePoints {
		for _, point := range points {
			freqs[point.r][point.c] = 'X'
		}

		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1 := points[i]
				p2 := points[j]
				dist := Point{p2.r - p1.r, p2.c - p1.c}
				freq1 := Point{p1.r - dist.r, p1.c - dist.c}
				freq2 := Point{p2.r + dist.r, p2.c + dist.c}

				for {
					if min(freq1.r, freq1.c) >= 0 && max(freq1.r, freq1.c) < len(matrix) {
						freqs[freq1.r][freq1.c] = 'X'
					} else {
						break
					}
					freq1.r -= dist.r
					freq1.c -= dist.c
				}

				for {
					if min(freq2.r, freq2.c) >= 0 && max(freq2.r, freq2.c) < len(matrix[0]) {
						freqs[freq2.r][freq2.c] = 'X'
					} else {
						break
					}
					freq2.r += dist.r
					freq2.c += dist.c
				}

			}
		}
	}

	tot := 0
	for r := range freqs {
		for c := range freqs {
			if freqs[r][c] == 'X' {
				tot += 1
			}
		}
	}

	fmt.Println(tot)
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
