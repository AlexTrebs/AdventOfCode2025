package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateDial(start int, diff int) (int, int, int) {
	newVal := start + (diff % 100)
	crossings := int(math.Abs(float64(diff) / 100))
	crossed := false

	if newVal < 0 {
		newVal = 100 + newVal
		crossed = true
	} else if newVal > 99 {
		newVal = newVal - 100
		crossed = true
	}

	if newVal == 0 {
		return newVal, 1, crossings + 1
	} else if crossed && start != 0 {
		return newVal, 0, crossings + 1
	}

	return newVal, 0, crossings
}

func singleInstruction(line []byte, start int) (int, int, int) {
	direction := string(line[0])
	num, err := strconv.Atoi(string(line[1:]))

	if err != nil {
		fmt.Println("Error converting string:", err)
		return start, 0, 0
	}

	if direction == "R" {
		return calculateDial(start, num)
	} else {
		return calculateDial(start, -num)
	}
}

func calcInstructions(lines []string, start int) {
	countDay1 := 0
	countDay2 := 0
	addCount1 := 0
	addCount2 := 0
	number := start

	for _, line := range lines {
		number, addCount1, addCount2 = singleInstruction([]byte(line), number)

		countDay1 += addCount1
		countDay2 += addCount2
	}

	fmt.Println("Final count for part one is: ", countDay1)
	fmt.Println("Final count for part two is: ", countDay2)
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
	calcInstructions(lines, 50)
}
