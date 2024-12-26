package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}

	defer file.Close()

	reports := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levelsStr := strings.Fields(line)
		levels := []int{}
		for _, levelStr := range levelsStr {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				continue
			}
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}

	if err := scanner.Err(); err != nil {
		return
	}

	count := 0
	countWithDampener := 0

	for _, report := range reports {
		if isSafe(report) {
			count++
			countWithDampener++
			continue
		}

		if isSafeWithDampener(report) {
			countWithDampener++
		}
	}

	fmt.Println(count)
	fmt.Println(countWithDampener)
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func skipAt(slice []int, skipIndex int) []int {
	if skipIndex < 0 || skipIndex >= len(slice) {
		return slice
	}
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:skipIndex]...)
	newSlice = append(newSlice, slice[skipIndex+1:]...)
	return newSlice
}

func isSafeWithDampener(levels []int) bool {
	for i := range levels {
		if isSafe(skipAt(levels, i)) {
			return true
		}
	}

	return false
}

func isSafe(levels []int) bool {
	lastDiff := 0
	levelsLen := len(levels)

	for i := range levels {
		if i == levelsLen-1 {
			break
		}

		l := levels[i]
		r := levels[i+1]

		diff := r - l
		if diff == 0 {
			return false
		}

		absDiff := absInt(diff)
		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if lastDiff < 0 && diff > 0 {
			return false
		}

		if lastDiff > 0 && diff < 0 {
			return false
		}

		lastDiff = diff
	}

	return true
}
