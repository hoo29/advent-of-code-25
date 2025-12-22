package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"time"
)

func p1(data []string) {

	start := time.Now()
	ans := 0
	var startInd int
	for ind, c := range data[0] {
		if c == 'S' {
			startInd = ind
			break
		}
	}
	queue := [][]int{{startInd, 1}}
	seen := map[string]bool{}
	for len(queue) > 0 {
		pos := queue[0]
		queue[0] = nil
		queue = queue[1:]

		if pos[1] >= len(data) {
			continue
		}
		line := data[pos[1]]
		if pos[0] < 0 || pos[0] >= len(line) {
			continue
		}
		key := fmt.Sprintf("%v_%v", pos[0], pos[1])
		if _, ok := seen[key]; ok {
			continue
		} else {
			seen[key] = true
		}

		char := utils.GetRuneFromString(line, pos[0])
		if char == '.' {
			queue = append(queue, []int{pos[0], pos[1] + 1})
		} else {
			ans += 1
			queue = append(queue, []int{pos[0] - 1, pos[1]})
			queue = append(queue, []int{pos[0] + 1, pos[1]})
		}
	}

	log.Printf("p1 took %s", time.Since(start))
	log.Printf("p1 ans %v", ans)
}

func p2(data []string) {

	start := time.Now()
	ans := 0

	log.Printf("p2 took %s", time.Since(start))
	log.Printf("p2 ans %v", ans)
}

func main() {
	test := false
	day := "d07"

	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	p1(data)
	p2(data)
}
