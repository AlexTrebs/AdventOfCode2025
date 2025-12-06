package main

import (
	"fmt"
	"os"
	"strconv"

	"strings"
)

func calculateEquation(nums []string, ops rune) int {
	total := 0

	for _, num := range nums {
		numInt, _ := strconv.Atoi(strings.TrimSpace(num))
		if total == 0 {
			total = numInt
		} else {
			if ops == '*' {
				total = total * numInt
			} else {
				total += numInt
			}
		}
	}

	return total
}

func doMath(nums [][]string, ops []rune, part int) {
	total := 0
	for i, numRow := range nums {
		total += calculateEquation(numRow, ops[i])
	}

	fmt.Printf("Part %d is: %d\n", part, total)
}

func partTwoMatrixManipulation(nums [][]string, lengths []int) [][]string {
	var allNums [][]string
	for i, l := range lengths {
		var newNums []string

		for j := l; j > 0; j-- {
			tempNum := ""
			for _, num := range nums[i] {
				if len(num) >= j {
					tempNum += string(num[j-1])
				}
			}

			if len(strings.TrimSpace(tempNum)) == 0 {
				continue
			}

			newNums = append(newNums, tempNum)
		}
		allNums = append(allNums, newNums)
	}

	return allNums
}

func splitRows(rows []string) ([]int, [][]string, []rune) {
	colLength := 0

	var lengths []int
	var ops []rune

	for i, char := range rows[4] {
		if char != ' ' {
			if i != 0 {
				lengths = append(lengths, colLength+1)
				colLength = 0
			}
			ops = append(ops, char)
		} else {
			colLength++
		}
	}
	lengths = append(lengths, colLength+1)

	nums := make([][]string, len(lengths))
	for i := range len(lengths) {
		nums[i] = make([]string, 4)
	}

	for colIndex, row := range rows[:4] {
		baseIndex := 0

		for rowIndex, l := range lengths {
			nums[rowIndex][colIndex] = row[baseIndex : baseIndex+l]
			baseIndex += l
		}
	}

	return lengths, nums, ops
}

func main() {
	file, _ := os.ReadFile("input.txt")

	rows := strings.Split(string(file), "\n")
	lens, nums, ops := splitRows(rows)

	doMath(nums, ops, 1)
	doMath(partTwoMatrixManipulation(nums, lens), ops, 2)
}
