package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	MIN_LEVEL = 1
	MAX_LEVEL = 3
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safeCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var levels []int
		for _, numStr := range strings.Fields(line) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Printf("Error parsing number: %v", err)
				continue
			}
			levels = append(levels, num)
		}

		if problemDampener(levels) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true 
	}

	increasing := levels[0] < levels[1]
	decreasing := levels[0] > levels[1]

	if !increasing && !decreasing {
		return false
	}

	for i := 1; i < len(levels); i++ {
		difference := levels[i] - levels[i-1]

		if !(MIN_LEVEL <= int(math.Abs(float64(difference))) && int(math.Abs(float64(difference))) <= MAX_LEVEL) {
			return false
		}

		if increasing && difference <= 0 {
			return false
		}
		if decreasing && difference >= 0 {
			return false
		}
	}

	return true
}

func problemDampener(levels []int) bool {
    if isSafe(levels) {
        return true
    }

    for i := range levels {
        toleredLevels := make([]int, len(levels)-1)
        copy(toleredLevels, levels[:i])
        copy(toleredLevels[i:], levels[i+1:])

        if isSafe(toleredLevels) {
            return true
        }
    }
    return false
}