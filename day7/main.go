package main

import (
	"AdventOfCode2025/common"
	"fmt"
	"slices"
)

type Coordinates struct{ y, x int }

func findRuneLocation(grid [][]rune, target rune) (int, int) {
	for y, row := range grid {
		for x, r := range row {
			if r == target {
				return y, x
			}
		}
	}
	return -1, -1
}

func appendIfNotExists(lasers []Coordinates, coords Coordinates) []Coordinates {
	if !slices.Contains(lasers, coords) {
		lasers = append(lasers, coords)
	}

	return lasers
}

func completeLaser(matrix [][]rune) (int, []Coordinates) {
	totalSplit := 0
	var lasers []Coordinates

	yStart, xStart := findRuneLocation(matrix, 'S')
	lasers = append(lasers, Coordinates{x: xStart, y: yStart})

	for yIndex, line := range matrix {
		for _, laser := range lasers {
			if laser.y == yIndex-1 {
				if line[laser.x] == '^' {
					totalSplit += 1
					lasers = appendIfNotExists(lasers, Coordinates{x: laser.x - 1, y: yIndex})
					lasers = appendIfNotExists(lasers, Coordinates{x: laser.x + 1, y: yIndex})
				} else {
					lasers = appendIfNotExists(lasers, Coordinates{y: yIndex, x: laser.x})
				}
			}
		}
	}

	return totalSplit, lasers
}

func spawnTimeline(lasers []Coordinates, startY int, startX int, maxLen int, visited map[Coordinates]int) int {
	coords := Coordinates{y: startY, x: startX}

	timeLines, exists := visited[coords]

	if exists {
		return timeLines
	}

	timeLines = calculateTimelines(lasers, startY, startX, maxLen, visited)
	return timeLines
}

func calculateTimelines(lasers []Coordinates, startY int, startX int, maxLen int, visited map[Coordinates]int) int {
	total := 0
	for i := startY + 1; i < maxLen; i++ {
		if slices.Contains(lasers, Coordinates{y: i, x: startX}) {
			continue // Beam continues straight, no split yet
		} else {
			total += spawnTimeline(lasers, i, startX-1, maxLen, visited)
			total += spawnTimeline(lasers, i, startX+1, maxLen, visited)
			break
		}
	}

	if total == 0 {
		total = 1
	}

	visited[Coordinates{x: startX, y: startY}] = total
	return total
}

func main() {
	file, _ := common.ReadGridFromFile("input.txt")
	splits, lasers := completeLaser(file)
	fmt.Println("Part 1 is: ", splits)

	startY, startX := findRuneLocation(file, 'S')
	visited := make(map[Coordinates]int)
	timelinesTotal := calculateTimelines(lasers, startY, startX, len(file), visited)
	fmt.Println("Part 2 is: ", timelinesTotal)
}
