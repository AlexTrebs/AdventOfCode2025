package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func totalList(l []int) int {
	total := 0

	for _, num := range l {
		total += num
	}

	return total
}

func splitIntoN(s string, n int, totalLen int) []string {
	var parts []string

	newLen := totalLen / n

	for i := 0; i < totalLen; i += newLen {
		parts = append(parts, string(s[i:i+newLen]))
	}

	return parts
}

func findInvalid(start int, end int) ([]int, []int) {
	oneRepitition := []int{}
	oneOrMoreRepitition := []int{}

	for id := start; id <= end; id++ {
		strId := strconv.Itoa(id)
		strLen := len(strId)

		for n := 2; n <= strLen; n++ {
			if strLen%n == 0 {
				parts := splitIntoN(strId, n, strLen)
				set := []string{}

				for _, part := range parts {
					if !slices.Contains(set, part) {
						set = append(set, part)
					}
				}

				if len(set) == 1 {
					oneOrMoreRepitition = append(oneOrMoreRepitition, id)
					if n == 2 {
						oneRepitition = append(oneRepitition, id)
					}
					break
				}
			}
		}
	}

	return oneRepitition, oneOrMoreRepitition
}

func calculateInvalid(productIds []string) {
	part1Invalid := []int{}
	part2Invalid := []int{}
	for _, id := range productIds {
		nums := strings.Split(id, "-")
		start, _ := strconv.Atoi(strings.TrimSpace(nums[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(nums[1]))

		additionalPart1, additionalPart2 := findInvalid(start, end)

		part1Invalid = append(part1Invalid, additionalPart1...)
		part2Invalid = append(part2Invalid, additionalPart2...)
	}

	part1Total := totalList(part1Invalid)
	part2Total := totalList(part2Invalid)

	fmt.Println("Part 1 total is: \n", part1Total)
	fmt.Println("Part 2 total is: \n", part2Total)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(content)
	splitFn := func(c rune) bool {
		return c == ','
	}
	productIds := strings.FieldsFunc(fileContent, splitFn)
	calculateInvalid(productIds)
}
