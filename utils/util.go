package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(day string, test bool) ([]string, error) {
	fileName := day
	if test {
		fileName = fmt.Sprintf("%s_t", day)
	}
	path := filepath.Join("data", fileName)
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
	err = scanner.Err()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
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
