package main

import (
	"fmt"
	"os"
	"strings"
)

const tie1 = "AX"
const tie2 = "BY"
const tie3 = "CZ"

const xWin = "CX"
const yWin = "AY"
const zWin = "BZ"

func CheckForTie(roundSubmission string) bool {
	return roundSubmission == tie1 || roundSubmission == tie2 || roundSubmission == tie3
}

func CheckForWin(roundSubmission string) bool {
	return roundSubmission == xWin || roundSubmission == yWin || roundSubmission == zWin
}

func GetToolVal(submission string) int8 {
	if submission == "X" || submission == "A" {
		return 1
	}
	if submission == "Y" || submission == "B" {
		return 2
	}
	return 3
}

func getAdjustedScore(entryArr []string) int {
	adjustedScore := 0
	for _, round := range entryArr {
		roundSubmission := strings.ReplaceAll(round, " ", "")
		if len(round) == 0 {
			continue
		}
		suggestion := roundSubmission[len(roundSubmission)-1:]
		opponentsSubmission := roundSubmission[0:1]

		opponentToolVal := GetToolVal(opponentsSubmission)
		if suggestion == "Y" {
			adjustedScore = adjustedScore + int(opponentToolVal) + 3
			continue
		}

		associatedWins := map[string]string{
			"C": "X",
			"A": "Y",
			"B": "Z",
		}
		if suggestion == "Z" {
			yourSuggestion := associatedWins[opponentsSubmission]
			toolVal := GetToolVal(yourSuggestion)
			adjustedScore = adjustedScore + int(toolVal) + 6
			continue
		}
		fmt.Println("they put ", opponentsSubmission, "so ")
		associatedLosses := map[string]string{
			"A": "Z",
			"B": "X",
			"C": "Y",
		}
		yourSuggestion := associatedLosses[opponentsSubmission]
		toolValue := GetToolVal(yourSuggestion)
		adjustedScore = adjustedScore + int(toolValue)
	}
	return adjustedScore
}

func main() {
	content, err := os.ReadFile("input-day2.txt")
	if err != nil {
		fmt.Println("megatron 1 error: ")
	}
	str := string(content)
	entryArr := strings.Split(str, "\n")
	fmt.Println(entryArr[0])
	totalScore := 0
	for i, round := range entryArr {
		roundSubmission := strings.ReplaceAll(round, " ", "")
		fmt.Println("ROUND #", i, ": ", round)
		if len(round) == 0 {
			continue
		}
		yourSubmission := roundSubmission[len(roundSubmission)-1:]
		fmt.Println("your tool submission ", yourSubmission)
		youTie := CheckForTie(roundSubmission)
		yourToolVal := GetToolVal(yourSubmission)

		if youTie {
			totalScore = totalScore + int(3+yourToolVal)
			continue
		}
		fmt.Println(roundSubmission, " round sub")
		youWin := CheckForWin(roundSubmission)

		if youWin {
			totalScore = totalScore + int(6+yourToolVal)
			continue
		}

		totalScore = totalScore + int(yourToolVal)
	}
	fmt.Println("PT1 Answer:  ", totalScore)

	adjustedScore := getAdjustedScore(entryArr)
	fmt.Println("PT 2 Answer: ", adjustedScore)
}
