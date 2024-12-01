package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Read lines from file
func getInput(filename string) ([]int, []int) {
	// Read the entire file content
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Convert the file content to a string and split it into lines
	lines := strings.Split(string(data), "\n")

	leftList := []int{}
	rightList := []int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		left, right := parseLine(line)
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return leftList, rightList
}

func parseLine(line string) (left int, right int) {
	parts := strings.SplitN(line, "   ", 2)

	left, _ = strconv.Atoi(parts[0])
	right, _ = strconv.Atoi(parts[1])

	return left, right
}

func getDistance(leftList []int, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	distance := 0

	for index, left := range leftList {
		diff := left - rightList[index]
		if diff < 0 {
			diff *= -1
		}
		distance += diff
	}

	return distance
}

func getSimilarity(leftList []int, rightList []int) int {
	similarity := 0

	rightMap := make(map[int]int)

	for _, right := range rightList {
		rightMap[right] += 1
	}

	for _, left := range leftList {
		similarity += left * rightMap[left]
	}

	return similarity
}

func main() {
	leftList, rightList := getInput("input.txt")
	fmt.Println(leftList, rightList)

	dist := getDistance(leftList, rightList)
	fmt.Println(dist)

	simi := getSimilarity(leftList, rightList)
	fmt.Println(simi)
}
