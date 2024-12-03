package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) []int {
	parts := strings.Split(line, " ")
	report := make([]int, len(parts))
	for index, part := range parts {
		level, _ := strconv.Atoi(part)
		report[index] = level
	}

	return report
}

func getInput(filename string) [][]int {
	// Read the entire file content
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Convert the file content to a string and split it into lines
	lines := strings.Split(string(data), "\n")

	reports := make([][]int, len(lines))

	for index, line := range lines {
		reports[index] = parseLine(line)
	}

	return reports
}

func checkReport(report []int) bool {
	// true if increasing, false if decreasing
	isIncrease := report[len(report)-1] < report[0]

	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]

		// increasing, so diff shoud be positive
		if isIncrease && diff < 0 {
			return false
		}

		// decreasing, so diff should be negative
		if !isIncrease && diff > 0 {
			return false
		}

		absDiff := diff
		if absDiff < 0 {
			absDiff *= -1
		}

		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	return true
}

// brute force
func checkReportWithFix(report []int) bool {
	isValid := checkReport(report)
	if isValid {
		return true
	}

	fmt.Println(report)

	for i := 0; i < len(report); i++ {
		newReport := make([]int, 0, len(report)-1)
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		isValid = checkReport(newReport)
		fmt.Println(isValid, newReport, report, i)
		if isValid {
			return true
		}
	}

	return false
}

func main() {
	reports := getInput("input.txt")
	fmt.Println(reports)

	validCount := 0

	for _, report := range reports {
		isValid := checkReportWithFix(report)
		if isValid {
			validCount += 1
		}
	}

	fmt.Println(validCount)
}
