package main

import (
	"aoc/utils"
	"log"
	"strings"
	"time"
)

func p1(data []string) {
	var numbers [][]int
	var ops []string
	for _, line := range data {
		parts := strings.Fields(line)
		numberLine := []int{}
		for _, part := range parts {
			if part == "+" || part == "*" {
				ops = append(ops, part)
			} else {
				numberLine = append(numberLine, utils.Atoi(part))
			}
		}
		if len(numberLine) > 0 {
			numbers = append(numbers, numberLine)
		}
	}

	start := time.Now()
	ans := 0
	for i := 0; i < len(ops); i++ {
		val := numbers[0][i]
		for j := 1; j < len(numbers); j++ {
			if ops[i] == "+" {
				val += numbers[j][i]
			} else {
				val *= numbers[j][i]
			}
		}
		ans += val
	}

	log.Printf("p1 took %s", time.Since(start))
	log.Printf("p1 ans %v", ans)
}

func main() {
	test := false
	day := "d06"

	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	p1(data)

	start := time.Now()
	ans := 0
	opLineInd := 4
	if test {
		opLineInd = 3
	}
	ops := data[opLineInd]
	ind := 0
	for ; ind < len(ops); ind++ {
		op := string(ops[ind])
		var cols []int
		for ; ind == len(ops)-1 || (ind < len(ops) && string(ops[ind+1]) == " "); ind++ {
			val := 0
			for j := 0; j < opLineInd; j++ {
				digit := string(data[j][ind])
				if digit == " " {
					continue
				}
				val = (val * 10) + utils.Atoi(digit)
			}
			cols = append(cols, val)
		}
		val := cols[0]
		for j := 1; j < len(cols); j++ {
			if op == "+" {
				val += cols[j]
			} else {
				val *= cols[j]
			}
		}
		ans += val
	}

	log.Printf("p2 took %s", time.Since(start))
	log.Printf("p2 %v", ans)
}
