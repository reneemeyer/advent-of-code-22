package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func getInputArr() ([]string, error) {
	content, err := os.ReadFile("input")
	if err != nil {
		return nil, err
	}
	str := string(content)
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1), nil
}

func main() {
	inputArr, err := getInputArr()
	if err != nil {
		fmt.Println("ERROR getting input data")
	}

	highestCalorieCount := 0
	var calorieCounts []int
	for _, elfCalorieSet := range inputArr {
		calorieArr := strings.Split(elfCalorieSet, "\n")

		sumOfCurrentSet := 0
		for _, v := range calorieArr {
			numericVal, err := strconv.Atoi(v)
			if err != nil {
				// skip the line
				numericVal = 0
			}
			sumOfCurrentSet = sumOfCurrentSet + numericVal
		}
		calorieCounts = append(calorieCounts, sumOfCurrentSet)
		if sumOfCurrentSet > highestCalorieCount {
			highestCalorieCount = sumOfCurrentSet
		}
	}
	fmt.Println("Top Elf has : ", highestCalorieCount, " Calories")

	sort.Ints(calorieCounts)

	secondPlace := calorieCounts[len(calorieCounts)-2]
	thirdPlace := calorieCounts[len(calorieCounts)-3]
	topThreeSum := highestCalorieCount + secondPlace + thirdPlace
	fmt.Println("The Top 3 elves have: ", topThreeSum)
}
