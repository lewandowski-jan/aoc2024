package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	sum := 0
	sumWithLogic := 0
	enabled := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pattern := `(mul\([0-9]+\,[0-9]+\)|do\(\)|don\'t\(\))`
		line := scanner.Text()

		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			operation := match[0]
			pattern := `[0-9]+`
			re := regexp.MustCompile(pattern)
			numbers := re.FindAllStringSubmatch(operation, -1)

			if numbers == nil {
				if operation == "do()" {
					enabled = true
				} else {
					enabled = false
				}

				continue
			}

			l := numbers[0][0]
			r := numbers[1][0]

			leftNumber, lErr := strconv.Atoi(l)
			rightNumber, rErr := strconv.Atoi(r)

			if lErr != nil || rErr != nil {
				continue
			}

			if enabled {
				sumWithLogic += leftNumber * rightNumber
			}
			sum += leftNumber * rightNumber
		}
	}

	fmt.Println(sum)
	fmt.Println(sumWithLogic)
}
