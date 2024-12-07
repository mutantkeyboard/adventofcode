// Solution for Advent of Code 2024 - Day 1

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Pair struct {
	First, Second int
}


func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var pairs []Pair

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var number1, number2 int
		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%d %d", &number1, &number2)
		if err != nil {
			log.Fatal(err)
		}
		pairs = append(pairs, Pair{First: number1, Second: number2})
	}

	firstList, secondList := extractAndSortLists(pairs)
	distance := calculateDistance(firstList, secondList)
	fmt.Println("Total distance is: ", distance)

	similarityScore := similarityScore(firstList, secondList)
	fmt.Println("Total similarity score is: ", similarityScore)




	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}	
}

func similarityScore(firstList []int, secondList []int) int {
	count := make(map[int]int)
	var similarityScore int
	for _, num := range secondList {
		count[num]++
	}

	for _, num := range firstList {
		similarityScore += num * count[num]
	}

	return similarityScore
}


func calculateDistance(firstList []int, secondList []int) int {
	var distance int
	for i := 0; i < len(firstList); i++ {
		distance += int(math.Abs(float64(firstList[i] - secondList[i])))
	}
	return distance
}

func extractAndSortLists(pairs []Pair) ([]int, []int) {
	firstList := make([]int, len(pairs))
	secondList := make([]int, len(pairs))
	// extract list
	for i, pair := range pairs {
		firstList[i] = pair.First
		secondList[i] = pair.Second
	}
	// sort list
	sort.Ints(firstList)
	sort.Ints(secondList)

	return firstList, secondList
}
