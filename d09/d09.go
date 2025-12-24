package main

import (
	"aoc/utils"
	"log/slog"
	"math"
	"strings"
	"time"
)

func coordsFromData(data []string) [][]int {
	coords := make([][]int, len(data))
	for i, d := range data {
		parts := strings.Split(d, ",")
		coords[i] = []int{utils.Atoi(parts[0]), utils.Atoi(parts[1])}
	}
	return coords
}

func calcArea(a []int, b []int) int {
	return (max(a[0], b[0]) - min(a[0], b[0]) + 1) * (max(a[1], b[1]) - min(a[1], b[1]) + 1)
}

func p1(data []string) {
	start := time.Now()
	ans := 0

	coords := coordsFromData(data)
	combs := utils.Combinations(coords)
	for _, c := range combs {
		area := calcArea(c[0], c[1])
		if area > ans {
			ans = area
		}
	}

	slog.Info("p1 ans", "value", ans)
	slog.Info("p1 took", "value", time.Since(start))
}

type coord struct {
	x int
	y int
}

func coordsStructFromData(data []string) []coord {
	coords := make([]coord, len(data))
	for i, d := range data {
		parts := strings.Split(d, ",")
		coords[i] = coord{
			x: utils.Atoi(parts[0]),
			y: utils.Atoi(parts[1]),
		}
	}
	return coords
}

func p2(data []string) {
	start := time.Now()
	ans := 0

	coords := coordsStructFromData(data)
	allData := map[coord]bool{}
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for i := 0; i < len(coords)-1; i++ {
		c1 := coords[i]
		c2 := coords[i+1]
		allData[c1] = true
		allData[c2] = true
		if c1.x == c2.x {
			for y := min(c1.y, c2.y); y < max(c1.y, c2.y); y++ {
				allData[coord{x: c1.x, y: y}] = true
			}
		} else {
			for x := min(c1.x, c2.x); x < max(c1.x, c2.x); x++ {
				allData[coord{x: x, y: c1.y}] = true
			}
		}
	}
	// combs := utils.Combinations(coords)

	slog.Info("p2 ans", "value", ans)
	slog.Info("p2 took", "value", time.Since(start))
}

func main() {
	test := false
	day := "d09"
	if !test {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	data, err := utils.ReadFile(day, test)
	if err != nil {
		slog.Error("failed to read file", "err", err)

	}
	p1(data)
	p2(data)
}
