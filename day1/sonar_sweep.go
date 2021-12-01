package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func loadInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, i)
	}
	return numbers, scanner.Err()
}

func countIncreases(depths []int) int {
	var increasesCount int
	var previous int
	for i, n := range depths {
		if i > 0 && n > previous {
			increasesCount++
		}
		previous = n
	}
	return increasesCount
}

func countIncreasesBySlidingWindow(depths []int) int {
	var increasesCount int
	for i := range depths {
		if i < len(depths)-3 {
			if depths[i+3] > depths[i] {
				increasesCount++
			}
		}
	}
	return increasesCount
}

func runForInput(path string) {
	log.Printf("-- %s", path)
	depths, err := loadInput(path)
	if err != nil {
		log.Fatalf("failed to load input: %s", err)
	}

	log.Printf("number of depths increases: %d", countIncreases(depths))
	log.Printf("number of depths increases (sliding window): %d", countIncreasesBySlidingWindow(depths))
}

func main() {
	runForInput("test.txt")
	runForInput("input.txt")
}
