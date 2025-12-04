package main

import (
	"AdventOfCode2025/common"
	"fmt"
	"log"
)

func normaliseVal(maxLen int, i int) int {
	if i < 0 {
		return maxLen - 1
	}
	if i == maxLen {
		return 0
	}

	return i
}

func checkMovable(y_index int, x_index int, y_len int, x_len int, warehouse [][]rune) bool {
	adjacentStacks := 0

	for y := y_index - 1; y <= y_index+1; y++ {
		if y >= 0 && y < y_len {
			for x := x_index - 1; x <= x_index+1; x++ {
				if x >= 0 && x < x_len {
					if x == x_index && y == y_index {
						continue
					}
					if warehouse[y][x] == '@' {
						adjacentStacks += 1
						if adjacentStacks >= 4 {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

func forkliftTime(warehouse [][]rune) {
	count := 0

	y_len := len(warehouse)
	x_len := len(warehouse[0])

	for y_index, row := range warehouse {
		for x_index, char := range row {
			if char == '@' {
				if checkMovable(y_index, x_index, y_len, x_len, warehouse) {
					count += 1
				}
			}
		}
	}

	fmt.Println("Part 1 is: ", count)
	count = 0
	total := 0

	for {
		for y_index, row := range warehouse {
			for x_index, char := range row {
				if char == '@' {
					if checkMovable(y_index, x_index, y_len, x_len, warehouse) {
						count += 1
						warehouse[y_index][x_index] = '.'
					}
				}
			}
		}

		if count == 0 {
			break
		}

		total += count
		count = 0
	}

	fmt.Println("Part 2 is: ", total)
}

func main() {
	warehouse, err := common.ReadGridFromFile("input.txt")

	if err != nil {
		log.Fatal("Cannot open file")
	}

	forkliftTime(warehouse)
}
