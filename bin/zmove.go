package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func frecent(rank float64, time int, now int) float64 {
	dx := now - time
	if dx < 3600 {
		return rank * 4
	} else if dx < 86400 {
		return rank * 2
	} else if dx < 604800 {
		return rank / 2
	}
	return rank / 4.0
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func main() {
	keyword := os.Args[1]
	now, _ := strconv.Atoi(os.Args[2])

	candidates := make(map[string]float64)

	data, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		fields := strings.Split(line, "|")
		if len(fields) != 3 {
			continue
		}

		path := fields[0]
		rank, _ := strconv.ParseFloat(fields[1], 64)
		time, _ := strconv.Atoi(fields[2])

		if !IsDir(path) {
			continue
		}

		if strings.ToLower(path) == strings.ToLower(keyword) || strings.Contains(strings.ToLower(path), strings.ToLower(keyword)) {
			candidates[path] = frecent(rank, time, now)
		}
	}

	if len(candidates) == 0 {
		os.Exit(0)
	}

	var bestPath string
	var bestScore float64
	for path, score := range candidates {
		if score > bestScore {
			bestPath = path
			bestScore = score
		}
	}

	fmt.Println(bestPath)
}
