package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FreshIdsRange struct{ start, end int }

func createFreshIndexList(ranges []string) []FreshIdsRange {
	var freshIds []FreshIdsRange

	for _, r := range ranges {
		splitR := strings.Split(r, "-")
		start, _ := strconv.Atoi(splitR[0])
		end, _ := strconv.Atoi(splitR[1])
		freshIds = append(freshIds, FreshIdsRange{start: start, end: end})
	}

	return freshIds
}

func findFresh(freshIds []FreshIdsRange, foods []string) []int {
	var freshFoods []int

	for _, food := range foods {
		foodId, _ := strconv.Atoi(food)

		for _, idsRange := range freshIds {
			if foodId >= idsRange.start && foodId <= idsRange.end {
				freshFoods = append(freshFoods, foodId)

				break
			}
		}
	}

	return freshFoods
}

func checkMergedOverlap(merged []FreshIdsRange, index int) []FreshIdsRange {
	for j := index + 1; j < len(merged); j++ {
		if merged[index].start <= merged[j].end+1 && merged[index].end >= merged[j].start-1 {
			if merged[j].start < merged[index].start {
				merged[index].start = merged[j].start
			}
			if merged[j].end > merged[index].end {
				merged[index].end = merged[j].end
			}
			merged = append(merged[:j], merged[j+1:]...)
		}
	}

	return merged
}

func countFreshIds(freshIds []FreshIdsRange) int {
	if len(freshIds) == 0 {
		return 0
	}

	var merged []FreshIdsRange

	for _, current := range freshIds {
		if len(merged) == 0 {
			merged = append(merged, current)
			continue
		}

		overlapped := false
		for i := 0; i < len(merged); i++ {
			if current.start <= merged[i].end+1 && current.end >= merged[i].start-1 {
				if current.start < merged[i].start {
					merged[i].start = current.start
				}
				if current.end > merged[i].end {
					merged[i].end = current.end
				}
				overlapped = true

				merged = checkMergedOverlap(merged, i)

				break
			}
		}

		if !overlapped {
			merged = append(merged, current)
		}
	}

	total := 0
	for _, r := range merged {
		total += r.end - r.start + 1
	}

	return total
}

func main() {
	file, _ := os.ReadFile("input.txt")
	contents := strings.Split(string(file), "\n\n")
	ranges := strings.Split(contents[0], "\n")
	foods := strings.Split(contents[1], "\n")

	freshIds := createFreshIndexList(ranges)
	freshFoods := findFresh(freshIds, foods)

	fmt.Println("Part one is: ", len(freshFoods))
	fmt.Println("Part two is: ", countFreshIds(freshIds))
}
