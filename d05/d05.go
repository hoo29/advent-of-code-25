package main

import (
	"aoc/utils"
	"log"
	"strings"
	"time"
)

func main() {
	test := false
	day := "d05"

	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	beforeBlank := true
	var ranges [][]int
	var ids []int
	for _, line := range data {
		if line == "" {
			beforeBlank = false

		} else if beforeBlank {
			parts := strings.Split(line, "-")
			ranges = append(ranges, []int{utils.Atoi(parts[0]), utils.Atoi(parts[1])})
		} else {
			ids = append(ids, utils.Atoi(line))
		}
	}

	start := time.Now()
	ans := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				ans += 1
				break
			}
		}
	}

	log.Printf("p1 took %s", time.Since(start))
	log.Printf("p1 ans %v", ans)

	start = time.Now()

	changed := true
	for changed {
		changed = false
		var (
			r1Ind, r2Ind int
			newRange     []int
		)
		for r1Ind = 0; r1Ind < len(ranges); r1Ind++ {
			for r2Ind = 0; r2Ind < len(ranges); r2Ind++ {
				if r1Ind == r2Ind {
					continue
				}
				r1 := ranges[r1Ind]
				r2 := ranges[r2Ind]
				if r1[0] >= r2[0] && r1[0] <= r2[1] {
					newRange = []int{min(r1[0], r2[0]), max(r1[1], r2[1])}
					changed = true
					break
				}
			}
			if changed {
				break
			}
		}
		if changed {
			ranges = append(ranges[:max(r1Ind, r2Ind)], ranges[max(r1Ind, r2Ind)+1:]...)
			ranges = append(ranges[:min(r1Ind, r2Ind)], ranges[min(r1Ind, r2Ind)+1:]...)
			ranges = append(ranges, newRange)
		}
	}
	ans = 0
	for _, r := range ranges {
		ans += r[1] - r[0] + 1
	}

	log.Printf("p2 took %s", time.Since(start))
	log.Printf("p2 %v", ans)
}
