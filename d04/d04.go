package main

import (
	"aoc/utils"
	"log"
	"strings"
	"time"
)

func checkBoundary(grid [][]string, x int, y int) bool {
	toCheck := [][]int{
		{x, y + 1},
		{x + 1, y + 1},
		{x + 1, y},
		{x + 1, y - 1},
		{x, y - 1},
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},
	}
	paperCount := 0
	for _, check := range toCheck {
		if check[1] >= len(grid) || check[1] < 0 {
			continue
		}
		line := grid[check[1]]
		if check[0] >= len(line) || check[0] < 0 {
			continue
		}
		if line[check[0]] == "@" {
			paperCount++
		}
	}
	return paperCount < 4
}

func main() {
	test := false
	day := "d04"

	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	grid := make([][]string, len(data))
	for y, line := range data {
		gridLine := make([]string, len(line))
		for x, c := range strings.Split(line, "") {
			gridLine[x] = string(c)
		}
		grid[y] = gridLine
	}

	start := time.Now()
	ans := 0
	for y := range len(grid) {
		for x := range len(grid[y]) {
			if grid[y][x] == "." {
				continue
			}
			if checkBoundary(grid, x, y) {
				ans += 1
			}
		}
	}
	log.Printf("p1 took %s", time.Since(start))
	log.Printf("p1 ans %v", ans)

	start = time.Now()
	total := 0
	removed := 0
	for removed != 0 {
		removed = 0
		var toRemove [][]int
		for y := range len(grid) {
			for x := range len(grid[y]) {
				if grid[y][x] == "." {
					continue
				}
				if checkBoundary(grid, x, y) {
					removed += 1
					total += 1
					toRemove = append(toRemove, []int{x, y})
				}

			}
		}
		for _, remove := range toRemove {
			grid[remove[1]][remove[0]] = "."
		}
	}
	log.Printf("p2 took %s", time.Since(start))
	log.Printf("p2 %v", total)

}
