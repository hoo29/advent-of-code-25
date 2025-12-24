package main

import (
	"aoc/utils"
	"log/slog"
	"math"
	"runtime"
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

func inShape(pos coord, allEdges map[uint64]bool, edgesX map[coord]bool, edgesY map[coord]bool, minX, maxX, minY, maxY int) bool {
	// slip and slide off the wall if we are on one
	x := pos.x
	for {
		key := uint64(x)<<32 | uint64(pos.y)
		_, ok := allEdges[key]
		if ok {
			x++
		} else {
			break
		}
	}
	if x >= maxX {
		return true
	}
	crossCount := 0

	for ; x <= maxX; x++ {
		key := uint64(x)<<32 | uint64(pos.y)
		_, ok := allEdges[key]
		if ok {
			crossCount++
		}
		for {
			key := uint64(x)<<32 | uint64(pos.y)
			_, ok := allEdges[key]
			if ok {
				x++
			} else {
				break
			}
		}
		// x++

		// // then keep going until we are off wall
		// for {
		// 	_, ok := edges[coord{x: x, y: pos.y}]
		// 	if !ok {
		// 		break
		// 	}
		// 	x++
		// }
	}
	return crossCount%2 == 1
}

func p2(data []string) {
	start := time.Now()
	ans := 0

	coords := coordsStructFromData(data)
	edgesX := map[coord]bool{}
	edgesY := map[coord]bool{}
	// allEdges := map[coord]bool{}
	allEdges := make(map[uint64]bool)
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for i := range coords {
		c1 := coords[i]
		var c2 coord
		if i == len(coords)-1 {
			c2 = coords[0]
		} else {
			c2 = coords[i+1]
		}
		if c1.x > maxX {
			maxX = c1.x
		}
		if c1.x < minX {
			minX = c1.x
		}
		if c1.y > maxY {
			maxY = c1.y
		}
		if c1.y < minY {
			minY = c1.y
		}
		// edges[c1] = true
		// edges[c2] = true
		if c1.x == c2.x {
			for y := min(c1.y, c2.y); y <= max(c1.y, c2.y); y++ {
				// edgesY[coord{x: c1.x, y: y}] = true
				key := uint64(c1.x)<<32 | uint64(y)
				allEdges[key] = true

			}
		} else {
			for x := min(c1.x, c2.x); x <= max(c1.x, c2.x); x++ {
				// edgesX[coord{x: x, y: c1.y}] = true
				key := uint64(x)<<32 | uint64(c1.y)
				allEdges[key] = true
			}
		}
	}
	combs := utils.Combinations(coords)

	// for every combo
	maxGoroutines := runtime.NumCPU()
	guard := make(chan struct{}, maxGoroutines)
	var wg sync.WaitGroup
	var mu sync.Mutex
	done := 0
	for outerInd, outerC := range combs {
		guard <- struct{}{}
		wg.Add(1)
		go func(ind int, c []coord) {
			defer wg.Done()
			defer func() { <-guard }()

			// if ind%100 == 0 {
			// 	slog.Info("progress", "ind", ind, "len", len(combs))
			// 	// if ind == 100 {
			// 	// 	return
			// 	// }
			// }
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
				if !inShape(check, allEdges, edgesX, edgesY, minX, maxX, minY, maxY) {
					allInside = false
					break
				}
			}
			if !allInside {
				mu.Lock()
				done++
				slog.Info("progress", "done", done, "len", len(combs))
				mu.Unlock()
				return
			}
			// if so party

			area := calcAreaCord(c[0], c[1])
			mu.Lock()
			done++
			if area > ans {
				slog.Debug("new largest area", "c1", c[0], "c2", c[1])
				ans = area
			}
			slog.Info("progress", "done", done, "len", len(combs))
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
	// f, _ := os.Create("cpu2.prof")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	p2(data)
}
