package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var max = map[string]int{
	"blue":  14,
	"green": 13,
	"red":   12,
}

func getText(filepath string) []string {
	b, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(string(b), "\n")
}

func possibleIds(str []string) []int {
	result := []int{}
	for i, s := range str {
		id := strings.Split(strings.Split(s, ":")[0], " ")[1]
		game := strings.ReplaceAll(s, fmt.Sprintf("Game %d: ", i+1), "")
		for _, sets := range strings.Split(game, ";") {
			gc := 0
			rc := 0
			bc := 0
			for _, set := range strings.Split(sets, ",") {
				count := strings.Split(strings.Trim(set, " "), " ")[0]
				color := strings.Split(strings.Trim(set, " "), " ")[1]
				if color == "red" {
					r, _ := strconv.Atoi(count)
					rc += r
				}
				if color == "green" {
					g, _ := strconv.Atoi(count)
					gc += g
				}
				if color == "blue" {
					b, _ := strconv.Atoi(count)
					bc += b
				}
			}
			r, _ := strconv.Atoi(id)
			add := true
			for c, m := range max {
				if c == "green" && gc > m {
					add = false
				}
				if c == "blue" && bc > m {
					add = false
				}
				if c == "red" && rc > m {
					add = false
				}
			}
			if add {
				result = append(result, r)
			}
		}

	}
	return result
}

func main() {
	str := getText("input.txt")
	sum := 0
	for _, id := range possibleIds(str) {
		sum += id
	}
	fmt.Println(sum)
	fmt.Println(possibleIds(str))
}
