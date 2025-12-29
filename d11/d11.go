package main

import (
	"aoc/utils"
	"fmt"
	"log/slog"
	"maps"
	"strings"
	"time"
)

func solveP1(node string, history map[string]bool, tree map[string][]string) int {
	if node == "out" {
		return 1
	}
	if ok, _ := history[node]; ok {
		return 0
	}
	history[node] = true

	ans := 0
	for _, n := range tree[node] {
		nHistory := maps.Clone(history)
		ans += solveP1(n, nHistory, tree)
	}
	return ans
}

func p1(data []string) {
	start := time.Now()
	ans := 0
	tree := map[string][]string{}
	for _, l := range data {
		parts := strings.Fields(l)
		tree[parts[0][:len(parts[0])-1]] = parts[1:]
	}

	ans = solveP1("you", map[string]bool{}, tree)

	slog.Info("p1 ans", "value", ans)
	slog.Info("p1 took", "value", time.Since(start))
}

func solveP2(node string, fft bool, dac bool, tree map[string][]string, cache map[string]int) int {
	if node == "out" {
		if fft && dac {
			return 1
		} else {
			return 0
		}
	}
	key := fmt.Sprintf("%s_%v_%v", node, fft, dac)
	if v, ok := cache[key]; ok {
		return v
	}

	fft = fft || node == "fft"
	dac = dac || node == "dac"
	ans := 0
	for _, n := range tree[node] {
		nodeAns := solveP2(n, fft, dac, tree, cache)
		ans += nodeAns
	}
	cache[key] = ans
	return ans
}

func p2(data []string) {
	start := time.Now()
	ans := 0
	tree := map[string][]string{}
	for _, l := range data {
		parts := strings.Fields(l)
		tree[parts[0][:len(parts[0])-1]] = parts[1:]
	}

	ans = solveP2("svr", false, false, tree, map[string]int{})

	slog.Info("p2 ans", "value", ans)
	slog.Info("p2 took", "value", time.Since(start))
}

func main() {
	test := false
	day := "d11"
	if test {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	data, err := utils.ReadFile(day, test)
	if err != nil {
		slog.Error("failed to read file", "err", err)
	}
	p1(data)
	p2(data)
}
