package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening a file")
		return
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) != 2 {
			continue
		}

		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			continue
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(left)
	sort.Ints(right)

	var distanceSum int
	for i, leftNum := range left {
		rightNum := right[i]
		distance := absInt(rightNum - leftNum)
		distanceSum += distance
	}

	fmt.Println(distanceSum)

	rightOccurances := make(map[int]int)
	leftSet := make(map[int]bool)
	for i, num := range right {
		leftSet[left[i]] = true
		rightOccurances[num]++
	}

	score := 0
	for leftNum := range leftSet {
		num, exists := rightOccurances[leftNum]
		if exists {
			score += leftNum * num
		}
	}

	fmt.Println(score)
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
