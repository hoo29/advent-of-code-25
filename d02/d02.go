package main

import (
	"aoc/utils"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	test := false
	day := "d02"
	start := time.Now()
	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	items := strings.Split(data[0], ",")
	p1 := 0
	for _, v := range items {
		ids := strings.Split(v, "-")
		start, err := strconv.Atoi(ids[0])
		if err != nil {
			log.Fatalf("number parse issue in %v", ids[0])
		}
		end, err := strconv.Atoi(ids[1])
		if err != nil {
			log.Fatalf("number parse issue in %v", ids[1])
		}
		for i := start; i <= end; i++ {
			valid := true
			digits := strings.Split(strconv.Itoa(i), "")
			if digits[0] == "0" {
				valid = false
			} else if len(digits)%2 == 0 {
				if slices.Equal(digits[0:len(digits)/2], digits[len(digits)/2:]) {
					valid = false
				}
			}
			if !valid {
				p1 += i
			}
		}
	}

	log.Printf("%v", p1)

	p2 := 0
	for _, v := range items {
		ids := strings.Split(v, "-")
		start, err := strconv.Atoi(ids[0])
		if err != nil {
			log.Fatalf("number parse issue in %v", ids[0])
		}
		end, err := strconv.Atoi(ids[1])
		if err != nil {
			log.Fatalf("number parse issue in %v", ids[1])
		}
		for number := start; number <= end; number++ {
			valid := true
			digits := strconv.Itoa(number)
			size := len(digits)
			for ind := 1; ind < size; ind++ {
				if size%ind != 0 {
					continue
				}
				str := strings.Repeat(digits[:ind], size/ind)
				if str == digits {
					valid = false
					break
				}
			}

			if !valid {
				p2 += number
			}
		}
	}

	log.Printf("%v", p2)
	log.Printf("took %s", time.Since(start))
}
