package main

import (
	"aoc/utils"
	"log"
	"strings"
	"time"
)

func main() {
	test := false
	day := "d06"

	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var numbers [][]int
	var ops []string
	for _, line := range data {
		parts := strings.Fields(line)
		numberLine := []int{}
		for _, part := range parts {
			if part == "+" || part == "*" {
				ops = append(ops, part)
			} else {
				numberLine = append(numberLine, utils.Atoi(part))
			}
		}
		if len(numberLine) > 0 {
			numbers = append(numbers, numberLine)
		}
	}

	start := time.Now()
	ans := 0
	for i := 0; i < len(ops); i++ {
		val := numbers[0][i]
		for j := 1; j < len(numbers); j++ {
			if ops[i] == "+" {
				val += numbers[j][i]
			} else {
				val *= numbers[j][i]
			}
		}
		ans += val
	}

	log.Printf("p1 took %s", time.Since(start))
	log.Printf("p1 ans %v", ans)

	start = time.Now()
	for i := 0; i < len(ops); i++ {
		maxDigitCount := 0
		for j := 0; j < len(numbers); j++ {
			// for
			if ops[i] == "+" {
				val += numbers[j][i]
			} else {
				val *= numbers[j][i]
			}
		}
		ans += val
	}

	log.Printf("p2 took %s", time.Since(start))
	log.Printf("p2 %v", ans)
}
