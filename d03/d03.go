package main

import (
	"aoc/utils"
	"log"
	"math"
	"strings"
	"time"
)

func p1(data []string) {
	p1 := 0
	for _, v := range data {
		maxInd := 0
		maxVal1 := -1
		for i := 0; i < len(v)-1; i++ {
			val := utils.Atoi(string(v[i]))
			if val > maxVal1 {
				maxVal1 = val
				maxInd = i
			}
		}
		maxVal2 := -1
		for i := maxInd + 1; i < len(v); i++ {
			val := utils.Atoi(string(v[i]))
			if val > maxVal2 {
				maxVal2 = val
				maxInd = i
			}
		}
		p1 += (maxVal1 * 10) + maxVal2
	}
	log.Printf("p1 %v", p1)
}

func findMax(data []int, start int, end int) (int, int) {
	max := -1
	ind := -1
	for i := start; i < end; i++ {
		if data[i] > max {
			max = data[i]
			ind = i
		}
	}
	return max, ind
}

func main() {
	test := false
	day := "d03"

	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	start := time.Now()
	p1(data)
	log.Printf("p1 took %s", time.Since(start))

	start = time.Now()
	total := 0
	maxOn := 12
	for _, line := range data {
		ints := make([]int, len(line))
		for ind, c := range strings.Split(line, "") {
			ints[ind] = utils.Atoi(c)
		}
		startInd := 0
		lineAns := 0
		for onCount := range maxOn {
			endInd := len(ints) - (maxOn - onCount - 1)
			val, valInd := findMax(ints, startInd, endInd)
			lineAns += val * int(math.Pow10(maxOn-onCount-1))
			startInd = valInd + 1
		}
		total += lineAns
		// log.Printf("line %v", line)
		// log.Printf("max %v", lineAns)
	}

	log.Printf("p2 %v", total)
	log.Printf("p2 took %s", time.Since(start))
}
