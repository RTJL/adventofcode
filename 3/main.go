package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Return an array of rune array
func getInput(filename string) [][]rune {
	// Read the entire file content
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Convert the file content to a string and split it into lines
	lines := strings.Split(string(data), "\n")

	var runeLines [][]rune
	for _, line := range lines {
		if line == "" {
			continue
		}
		runeLines = append(runeLines, []rune(line))
	}

	return runeLines
}

func checkLine(currLine *[]rune, prevLine *[]rune, nextLine *[]rune) int {
	lineTotal := 0

	for index, char := range *currLine {
		if char == '.' || unicode.IsDigit(char) {
			continue
		}

		// check prevline
		topLeftNum := getNum(prevLine, index-1)
		topNum := getNum(prevLine, index)
		topRightNum := getNum(prevLine, index+1)
		fmt.Println("top", topLeftNum, topNum, topRightNum)
		lineTotal += topLeftNum + topNum + topRightNum

		// check currLine
		currleftNum := getNum(currLine, index-1)
		currRightNum := getNum(currLine, index+1)
		lineTotal += currleftNum + currRightNum
		fmt.Println("current", currleftNum, currRightNum)

		// check nextLine
		bottomLeftNum := getNum(nextLine, index-1)
		bottomNum := getNum(nextLine, index)
		bottomRightNum := getNum(nextLine, index+1)
		lineTotal += bottomLeftNum + bottomNum + bottomRightNum
		fmt.Println("bottom", bottomLeftNum, bottomNum, bottomRightNum)

		fmt.Println()
	}

	fmt.Println(lineTotal)

	return lineTotal
}

// Return the number if it is within the searchIndex
// Updates the line, change the number to `.`
//
// Returns:
//
//	0 if number not found or the actual number
func getNum(line *[]rune, searchIndex int) int {
	if line == nil {
		return 0
	}

	if searchIndex < 0 || searchIndex >= len(*line) {
		return 0
	}

	if !unicode.IsDigit((*line)[searchIndex]) {
		return 0
	}

	// find the start index of the number
	start := searchIndex
	for {
		if start == 0 || !unicode.IsDigit((*line)[start-1]) {
			// current start is the actual start of the number index
			break
		}
		start -= 1
	}

	// find the end index of the number
	end := searchIndex
	for {
		if end == len(*line)-1 || !unicode.IsDigit(rune((*line)[end+1])) {
			// current end is the actual end of the numebr index
			break
		}
		end += 1
	}

	// convert into number
	numberStr := string((*line)[start : end+1])
	number, _ := strconv.Atoi(numberStr)

	// remove number from line
	for i := start; i <= end; i++ {
		(*line)[i] = '.'
	}

	return number
}

func main() {
	res := 0

	lines := getInput("input.txt")
	for index, line := range lines {
		fmt.Println(line)

		var prevLine *[]rune
		if index > 0 {
			prevLine = &lines[index-1]
		}

		var nextLine *[]rune
		if index < len(lines)-1 {
			nextLine = &lines[index+1]
		}
		res += checkLine(&line, prevLine, nextLine)
	}

	fmt.Println("total", res)
}
