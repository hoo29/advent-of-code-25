package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ReadFile(day string, test bool) ([]string, error) {
	fileName := day
	if test {
		fileName = fmt.Sprintf("%s_t", day)
	}
	path := filepath.Join("..", "data", fileName)
	var data []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	return data, scanner.Err()
}

func Mod(val int, mod int) int {
	return (val%mod + mod) % mod
}

func Abs(val int) int {
	if val < 0 {
		val = -val
	}
	return val
}

func Atoi(val string) int {
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("number parse issue in %v", val)
	}
	return valInt
}
