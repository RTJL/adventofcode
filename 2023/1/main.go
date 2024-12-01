package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var valueMap = map[string]rune{"one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'}

func getInput(filename string) []string {
	// Read the entire file content
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Convert the file content to a string and split it into lines
	lines := strings.Split(string(data), "\n")

	return lines
}

func getCalibrationValue(line string) int {
	runes := []rune(line)

	var firstChar byte
findFirstLoop:
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			firstChar = line[i]
			break findFirstLoop
		} else {
			substring := string(runes[0 : i+1])
			for k, v := range valueMap {
				if strings.Contains(substring, k) {
					firstChar = byte(v)
					// fmt.Println(substring, " ", k)
					break findFirstLoop
				}
			}
		}
	}

	var secondChar byte
findLastLoop:
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			secondChar = line[i]
			break findLastLoop
		} else {
			substring := string(runes[i:len(line)])
			for k, v := range valueMap {
				if strings.Contains(substring, k) {
					secondChar = byte(v)
					// fmt.Println(substring, " ", k)
					break findLastLoop
				}
			}
		}
	}

	calibrationValue := int(firstChar-'0')*10 + int(secondChar-'0')

	return calibrationValue
}

func getCalibrationSum(lines []string) int {
	total := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		calibrationValue := getCalibrationValue(line)
		total += calibrationValue

		fmt.Println(calibrationValue, " ", line)
	}

	return total
}

func main() {
	lines := getInput("input.txt")

	calibrationSum := getCalibrationSum(lines)
	fmt.Println("total: ", calibrationSum)
}
