package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	rows := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		chars := []string{}
		for _, rune := range runes {
			chars = append(chars, string(rune))
		}

		rows = append(rows, chars)
	}

	if err := scanner.Err(); err != nil {
		return
	}

	counter := 0

	var horizontal []string
	for i := 0; i < len(rows); i++ {
		line := ""
		for j := 0; j < len(rows[0]); j++ {
			line += rows[i][j]
		}
		horizontal = append(horizontal, line)
	}

	var vertical []string
	for i := 0; i < len(rows); i++ {
		line := ""
		for j := 0; j < len(rows[0]); j++ {
			line += rows[j][i]
		}
		vertical = append(vertical, line)
	}

	var diagonalRight []string
	for i := 0; i < len(rows)*2; i++ {
		line := ""
		for j := -len(rows[0]); j < len(rows[0]); j++ {
			effectiveI := len(rows[0]) + j
			effectiveJ := j + i
			if !isIn(effectiveI, effectiveJ, len(rows), len(rows[0])) {
				continue
			}

			line += rows[effectiveI][effectiveJ]
		}

		if line == "" {
			continue
		}
		diagonalRight = append(diagonalRight, line)
	}

	var diagonalLeft []string
	for i := 0; i <= len(rows)*2; i++ {
		line := ""
		for j := len(rows[0]) * 2; j >= 0; j-- {
			effectiveI := len(rows[0])*2 - j
			effectiveJ := j - i
			if !isIn(effectiveI, effectiveJ, len(rows), len(rows[0])) {
				continue
			}
			line += rows[effectiveI][effectiveJ]
		}

		if line == "" {
			continue
		}
		diagonalLeft = append(diagonalLeft, line)
	}

	allStrings := []string{}
	allStrings = append(allStrings, horizontal...)
	allStrings = append(allStrings, vertical...)
	allStrings = append(allStrings, diagonalRight...)
	allStrings = append(allStrings, diagonalLeft...)

	reForward := regexp.MustCompile("XMAS")
	reBackward := regexp.MustCompile("SAMX")

	for _, str := range allStrings {
		forward := reForward.FindAllString(str, -1)
		backward := reBackward.FindAllString(str, -1)
		counter += len(forward) + len(backward)
	}

	fmt.Println(counter)

	crossMasCounter := 0
	for i := 1; i < len(rows)-1; i++ {
		for j := 1; j < len(rows[0])-1; j++ {
			rightDiagonal := rows[i-1][j-1] + rows[i][j] + rows[i+1][j+1]
			leftDiagonal := rows[i+1][j-1] + rows[i][j] + rows[i-1][j+1]

			if rightDiagonal != "SAM" && rightDiagonal != "MAS" {
				continue
			}
			if leftDiagonal != "SAM" && leftDiagonal != "MAS" {
				continue
			}

			crossMasCounter++
		}
	}

	fmt.Println(crossMasCounter)
}

func isIn(i int, j int, rows int, cols int) bool {
	if i < 0 || i >= rows {
		return false
	}

	if j < 0 || j >= cols {
		return false
	}

	return true
}
