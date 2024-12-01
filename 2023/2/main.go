package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BallSet struct {
}

type Game struct {
	id       int
	ballSets []map[string]int
}

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

func getGame(line string) Game {
	// Game 1: 1 green, 6 red, 4 blue; 2 blue, 6 green, 7 red; 3 red, 4 blue, 6 green; 3 green; 3 blue, 2 green, 1 red
	lineSplit := strings.SplitN(line, ": ", 2)

	idStr, _ := strings.CutPrefix(lineSplit[0], "Game ")
	id, _ := strconv.Atoi(idStr)
	// fmt.Println(id)

	var ballSetsSlice []map[string]int

	sets := strings.Split(lineSplit[1], ";")
	for _, set := range sets {
		trimSet := strings.TrimSpace(set)
		// fmt.Println(trimSet)
		ballSet := strings.Split(trimSet, ", ")

		ballSetsMap := make(map[string]int)

		for _, ball := range ballSet {
			ballSplit := strings.Split(ball, " ")
			num, _ := strconv.Atoi(ballSplit[0])

			// fmt.Println(ballSplit[1], "->", num)
			ballSetsMap[ballSplit[1]] = num
		}

		ballSetsSlice = append(ballSetsSlice, ballSetsMap)
	}

	game := Game{
		id:       id,
		ballSets: ballSetsSlice,
	}
	return game
}

func isValidGame(game Game) bool {
	for _, ballSet := range game.ballSets {
		redCount, ok := ballSet["red"]
		if ok && redCount > 12 {
			return false
		}

		blueCount, ok := ballSet["blue"]
		if ok && blueCount > 14 {
			return false
		}

		greenCount, ok := ballSet["green"]
		if ok && greenCount > 13 {
			return false
		}
	}

	return true
}

func getGamePower(game Game) int {
	minBallCount := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	for _, ballSet := range game.ballSets {
		currRedCount, ok := minBallCount["red"]
		redCount, ok := ballSet["red"]
		if ok && redCount > currRedCount {
			minBallCount["red"] = redCount
		}

		currBlueCount, ok := minBallCount["blue"]
		blueCount, ok := ballSet["blue"]
		if ok && blueCount > currBlueCount {
			minBallCount["blue"] = blueCount
		}

		currGreenCount, ok := minBallCount["green"]
		greenCount, ok := ballSet["green"]
		if ok && greenCount > currGreenCount {
			minBallCount["green"] = greenCount
		}
	}

	power := 1
	for _, v := range minBallCount {
		if v != 0 {
			power *= v
		}
	}

	return power
}

func main() {
	lines := getInput("input.txt")

	power := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		game := getGame(line)
		// isValid := isValidGame(game)
		// if isValid {
		// 	sum += game.id
		// }

		power += getGamePower(game)
	}

	fmt.Println(power)
}
