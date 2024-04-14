package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var nums = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getStrInput(filePath string) (str []string, err error) {
	b, err := os.ReadFile("../assets/inputs.txt")

	if err != nil {
		fmt.Println(err)
	}

	str = strings.Split(string(b), "\n")

	return
}
func combineNumbersInToInt(x int, y int) int {
	strNum := fmt.Sprint(x) + fmt.Sprint(y)
	n, _ := strconv.Atoi(strNum)
	return n
}

func getNumber(s string) int {
	firstNum := -1
	lastNum := -1

	for i, r := range s {
		if firstNum == -1 && unicode.IsDigit(r) {
			firstNum = int(r - '0')
			lastNum = firstNum
			continue

		}
		if firstNum > -1 && unicode.IsDigit(r) {
			lastNum = int(r - '0')
			continue
		}
		for key, num := range nums {
			if i+len(key) <= len(s) && s[i:i+len(key)] == key {
				if firstNum == -1 {
					firstNum = num
				}
				lastNum = num
				continue
			}
		}

	}
	if firstNum == -1 {
		return firstNum
	}
	return combineNumbersInToInt(firstNum, lastNum)
}
func getNumbersFromStr(str []string) []int {
	var numAr []int

	for _, s := range str {
		n := getNumber(s)
		if n != -1 {
			numAr = append(numAr, n)
		}

	}
	return numAr
}
func getSum(numArr []int) int {
	sum := 0
	for _, n := range numArr {
		sum += n
	}
	return sum
}

func main() {
	strArr, _ := getStrInput("../assets/inputs.txt")
	numAr := getNumbersFromStr(strArr)

	leSum := getSum(numAr)

	fmt.Println(leSum)
	fmt.Println(nums)
}
