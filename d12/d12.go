package main

import (
	"aoc/utils"
	"log/slog"
	"strings"
	"time"
)

func p1(data []string, test bool) {
	start := time.Now()
	ans := 0

	// lol?
	var presents []int
	if test {
		presents = []int{7, 7, 7, 7, 7, 7}
	} else {
		presents = []int{6, 7, 5, 7, 7, 7}
	}

	for _, l := range data {
		parts := strings.Split(l, ":")
		areaParts := strings.Split(parts[0], "x")
		area := utils.Atoi(areaParts[0]) * utils.Atoi(areaParts[1])
		sizeGuess := 0
		for i, n := range strings.Fields(parts[1]) {
			sizeGuess += presents[i] * utils.Atoi(n)
		}
		if area > sizeGuess {
			ans += 1
		}
	}

	slog.Info("p1 ans", "value", ans)
	slog.Info("p1 took", "value", time.Since(start))
}

func main() {
	test := false
	day := "d12"
	if test {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	data, err := utils.ReadFile(day, test)
	if err != nil {
		slog.Error("failed to read file", "err", err)
	}
	p1(data, test)
}
