package main

import (
	"aoc/utils"
	"log/slog"
	"math"
	"strconv"
	"strings"
	"time"
)

type machine struct {
	state   []bool
	buttons [][]int
	reqs    []int
}

func parseMachines(data []string) []machine {
	machines := make([]machine, len(data))
	for i, d := range data {
		stateEnd := utils.FindCharIndex(']', d)
		state := make([]bool, stateEnd-1)
		for s := 1; s < stateEnd; s++ {
			if utils.GetRuneFromString(d, s) == '#' {
				state[s-1] = true
			} else {
				state[s-1] = false
			}
		}
		buttonsEnd := utils.FindCharIndex('{', d)
		buttonsRaw := strings.Split(d[stateEnd+2:buttonsEnd-1], " ")
		buttons := [][]int{}
		for _, buttonRaw := range buttonsRaw {
			sub := buttonRaw[1 : len(buttonRaw)-1]
			l := []int{}
			for b := range strings.SplitSeq(sub, ",") {
				l = append(l, utils.Atoi(b))
			}
			buttons = append(buttons, l)
		}

		reqs := []int{}
		for r := range strings.SplitSeq(d[buttonsEnd+1:len(d)-1], ",") {
			reqs = append(reqs, utils.Atoi(r))
		}

		machines[i] = machine{
			state:   state,
			buttons: buttons,
			reqs:    reqs,
		}
	}

	return machines
}

func checkMatch[T comparable](desiredState []T, currenState []T) bool {
	for i := range len(desiredState) {
		if currenState[i] != desiredState[i] {
			return false
		}
	}
	return true
}

func alterState(state []bool, button []int) []bool {
	newState := make([]bool, len(state))
	copy(newState, state)
	for _, i := range button {
		newState[i] = !state[i]
	}

	return newState
}

func historyKey(state []bool) uint32 {
	var key uint32
	key = 0
	for _, b := range state {
		key <<= 1
		if b {
			key |= 1
		}
	}
	return key
}

func buttonToString(button []int) string {
	var sb strings.Builder
	sb.WriteRune('(')
	for i, b := range button {
		sb.WriteString(strconv.Itoa(b))
		if i < len(button)-1 {
			sb.WriteRune(',')
		}
	}
	sb.WriteRune(')')
	sb.WriteRune(' ')
	return sb.String()
}

func solveP1(desiredState []bool, currenState []bool, buttons [][]int, depth int, maxDepth int, history map[uint32]int, seq string, test bool) int {
	key := historyKey(currenState)
	if v, ok := history[key]; ok && v <= depth {
		return math.MaxInt
	}
	history[key] = depth
	if depth > maxDepth {
		return math.MaxInt
	}

	if checkMatch(desiredState, currenState) {
		if test {
			slog.Info("matched", "pressed", seq)
		}
		return depth
	}

	minVal := math.MaxInt
	for _, b := range buttons {
		newSeq := seq
		if test {
			newSeq += buttonToString(b)
		}
		newState := alterState(currenState, b)
		ans := solveP1(desiredState, newState, buttons, depth+1, maxDepth, history, newSeq, test)
		if ans < minVal {
			minVal = ans
		}
	}

	return minVal
}

func p1(data []string, test bool) {
	start := time.Now()
	ans := 0

	machines := parseMachines(data)
	for _, m := range machines {
		history := map[uint32]int{}
		a := solveP1(m.state, make([]bool, len(m.state)), m.buttons, 0, 20, history, "", test)
		ans += a
	}

	slog.Info("p1 ans", "value", ans)
	slog.Info("p1 took", "value", time.Since(start))
}

func main() {
	test := false
	day := "d10"
	if test {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	data, err := utils.ReadFile(day, test)
	if err != nil {
		slog.Error("failed to read file", "err", err)
	}
	p1(data, test)
}
