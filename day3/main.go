package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateLargest(row []string, size int) string {
	largest := 0
	largestIndex := 0

	totalLen := len(row)

	for index, numStr := range row {
		num, _ := strconv.Atoi(numStr)

		if index > totalLen-size {
			break
		}

		if num > largest {
			largest = num
			largestIndex = index
		}
	}

	largestStr := strconv.Itoa(largest)
	restOfStr := ""

	if size != 1 {
		restOfStr = calculateLargest(row[largestIndex+1:], size-1)
	}
	return largestStr + restOfStr
}

func calculateJoltage(rows []string) {
	sumPartOne := 0
	sumPartTwo := 0
	for _, r := range rows {
		numPartOne, _ := strconv.Atoi(calculateLargest(strings.Split(strings.TrimSpace(r), ""), 2))
		numPartTwo, _ := strconv.Atoi(calculateLargest(strings.Split(strings.TrimSpace(r), ""), 12))

		sumPartOne += numPartOne
		sumPartTwo += numPartTwo
	}

	fmt.Println("Result part 1: ", sumPartOne)
	fmt.Println("Result part 2: ", sumPartTwo)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(content)
	splitFn := func(c rune) bool {
		return c == '\n'
	}
	lines := strings.FieldsFunc(fileContent, splitFn)
	calculateJoltage(lines)
}
