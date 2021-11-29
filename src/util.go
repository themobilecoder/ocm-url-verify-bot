package main

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(path string) (map[string]bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urlMap = map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urlMap[strings.TrimSpace(scanner.Text())] = true
	}
	return urlMap, scanner.Err()
}
