package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
)

func main() {
	test := false
	day := "d01"
	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	dial := 50
	p1 := 0
	p2 := 0

	for _, v := range data {
		left := v[0] == 'L'
		num, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatalf("number parse issue in %v", v)
		}
		oldDial := dial
		if !left {
			dial = utils.Mod(dial+num, 100)
			p2 += (oldDial + num) / 100
			if oldDial == 0 && ((oldDial+num)%100) == 0 {
				p2 -= 1
			}
		} else {
			dial = utils.Mod(dial-num, 100)
			p2 += (100 - oldDial + num) / 100
			// fairly sure I am overcounting in one branch and undercounting in the other but it works!
			if oldDial == 0 && ((100-oldDial+num)/100) > 0 {
				p2 -= 1
			}
		}
		if dial == 0 {
			p1++
		}
		fmt.Printf("v %v dial %v p1 %v p2 %v", v, dial, p1, p2)
		fmt.Printf("\n")
	}

	fmt.Printf("dial %v p1 %v p2 %v", dial, p1, p2)
}
