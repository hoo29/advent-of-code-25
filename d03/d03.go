package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"math"
	"sync"
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

func p2(data string, onCount int, ind int, sum int, maxOnCount int, cache map[string]int, maxCache map[int]int) int {

	if onCount == maxOnCount {
		return sum
	}

	if ind == len(data) {
		return -1
	}

	if onCount+(len(data)-ind) < maxOnCount {
		return -1
	}

	if maxCache[onCount] > sum {
		return -1
	}

	key := fmt.Sprintf("%v_%v_%v", onCount, sum, ind)
	if val, ok := cache[key]; ok {
		// log.Println("cache hit")
		return val
	}

	sum1 := sum + utils.Atoi(string(data[ind]))*int(math.Pow10(maxOnCount-onCount-1))

	val1 := p2(data, onCount+1, ind+1, sum1, maxOnCount, cache, maxCache)
	val2 := p2(data, onCount, ind+1, sum, maxOnCount, cache, maxCache)
	maxVal := max(val1, val2)
	cache[key] = maxVal
	maxCache[onCount] = maxVal
	return maxVal
}

func main() {
	test := false
	day := "d03"

	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	start := time.Now()
	// p1(data)
	// log.Printf("p1 took %s", time.Since(start))

	start = time.Now()
	total := 0
	var wg sync.WaitGroup
	var mut sync.Mutex
	for i, line := range data {
		wg.Add(1)
		go func(l string) {
			log.Printf("%v/%v", i+1, len(data))
			cache := make(map[string]int)
			maxCache := make(map[int]int)

			res := p2(l, 0, 0, 0, 12, cache, maxCache)
			log.Printf("max %v", res)
			mut.Lock()
			total += res
			mut.Unlock()
		}(line)
	}
	wg.Wait()

	log.Printf("p2 %v", total)
	log.Printf("p2 took %s", time.Since(start))
}
