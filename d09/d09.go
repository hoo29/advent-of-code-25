package main

import (
	"aoc/utils"
	"log/slog"
	"runtime"
	"slices"
	"sort"
	"strings"
	"sync"
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

func calcAreaCord(a, b coord) int {
	return (max(a.x, b.x) - min(a.x, b.x) + 1) * (max(a.y, b.y) - min(a.y, b.y) + 1)
}

func inShape(pos coord, yCoords map[int][]int) bool {

	row := yCoords[pos.y]
	// too far to the left
	if pos.x < row[0] {
		return false
	}
	// too far to the right
	if pos.x > row[len(row)-1] {
		return false
	}
	rowInd := sort.SearchInts(row, pos.x)
	// if we are on a wall must be in shape
	if row[rowInd] == pos.x {
		return true
	}

	// this seems like it should be wrong but it works so...
	crossCount := 1
	for ; rowInd < len(row)-1; rowInd++ {
		if row[rowInd+1] != row[rowInd]+1 {
			crossCount++
		}
	}
	return crossCount%2 == 1
}

func p2(data []string) {
	start := time.Now()
	ans := 0

	coords := coordsStructFromData(data)
	yCoords := map[int][]int{}
	for i := range coords {
		c1 := coords[i]
		var c2 coord
		if i == len(coords)-1 {
			c2 = coords[0]
		} else {
			c2 = coords[i+1]
		}
		if c1.x == c2.x {
			if c1.y < c2.y {
				for y := c1.y; y < c2.y; y++ {
					yCoords[y] = append(yCoords[y], c1.x)
				}
			} else {
				for y := c1.y; y > c2.y; y-- {
					yCoords[y] = append(yCoords[y], c1.x)
				}
			}
		} else {
			if c1.x < c2.x {
				for x := c1.x; x < c2.x; x++ {
					yCoords[c1.y] = append(yCoords[c1.y], x)
				}
			} else {
				for x := c1.x; x > c2.x; x-- {
					yCoords[c1.y] = append(yCoords[c1.y], x)
				}
			}
		}
	}
	for _, v := range yCoords {
		slices.Sort(v)
	}
	combs := utils.Combinations(coords)

	// for every combo
	maxGoroutines := runtime.NumCPU()
	guard := make(chan bool, maxGoroutines)
	// guard := make(chan bool, 1)
	defer close(guard)
	var wg sync.WaitGroup
	var mu sync.Mutex
	done := 0
	for outerInd, outerC := range combs {
		guard <- true
		wg.Add(1)
		go func(ind int, c []coord) {
			defer wg.Done()
			defer func() { <-guard }()

			// ignore straight line matches
			if c[0].x == c[1].x || c[0].y == c[1].y {
				return
			}

			toCheck := []coord{}
			// for every point on both perimeters
			for x := min(c[0].x, c[1].x) + 1; x <= max(c[0].x, c[1].x); x++ {
				for _, y := range []int{min(c[0].y, c[1].y), max(c[0].y, c[1].y)} {
					toCheck = append(toCheck, coord{
						x: x,
						y: y,
					})
				}
			}
			for y := min(c[0].y, c[1].y) + 1; y <= max(c[0].y, c[1].y); y++ {
				for _, x := range []int{min(c[0].x, c[1].x), max(c[0].x, c[1].x)} {
					toCheck = append(toCheck, coord{
						x: x,
						y: y,
					})
				}
			}
			// check if inside shape
			allInside := true
			for _, check := range toCheck {
				if !inShape(check, yCoords) {
					allInside = false
					break
				}
			}
			if allInside {
				// if so party
				area := calcAreaCord(c[0], c[1])
				mu.Lock()
				if area > ans {
					slog.Debug("new largest area", "c1", c[0], "c2", c[1])
					ans = area
				}
				mu.Unlock()
			}
			mu.Lock()
			done++
			if done%1000 == 0 {
				slog.Info("progress", "done", done, "len", len(combs))
			}
			mu.Unlock()
		}(outerInd, outerC)
	}
	wg.Wait()
	slog.Info("p2 ans", "value", ans)
	slog.Info("p2 took", "value", time.Since(start))
}

func main() {
	test := false
	day := "d09"
	if test {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	data, err := utils.ReadFile(day, test)
	if err != nil {
		slog.Error("failed to read file", "err", err)
	}
	p1(data)
	// takes 3 minutes on 16 threads but close enough
	p2(data)
}
