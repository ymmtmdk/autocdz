package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pwd := os.Args[1]
	now, _ := strconv.Atoi(os.Args[2])
	result := make(map[string]map[string]int)

	file, err := os.Open(os.Args[3])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		path := fields[0]
		rank, _ := strconv.ParseFloat(fields[1], 64)
		time, _ := strconv.Atoi(fields[2])

		if rank >= 1 {
			result[path] = map[string]int{"rank": int(rank), "time": time}
		}
	}

	if _, ok := result[pwd]; ok {
		result[pwd]["rank"]++
		result[pwd]["time"] = now
	} else {
		result[pwd] = map[string]int{"rank": 1, "time": now}
	}

	rankTotal := 0
	for _, v := range result {
		rankTotal += v["rank"]
	}

	for path, h := range result {
		rank := h["rank"]
		if rankTotal > 1000 {
			rank = int(float64(rank) * 0.9)
		}
		fmt.Printf("%s|%d|%d\n", path, rank, h["time"])
	}
}
