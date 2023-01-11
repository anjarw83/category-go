package main

import (
	"fmt"
	"strings"
)

const OPEN_TAGS string = "({["
const CLOSED_TAGS string = ")}]"

// 1. Logical Test
func main() {
	correct_input := "123{abcd[123(45)dd]bb}sss"
	invalid_input := "abcd(ex45{uuuu)000]ccc"

	fmt.Println("Input : ", correct_input)
	fmt.Println("Output : ", parse(correct_input))

	fmt.Println("Input : ", invalid_input)
	fmt.Println("Output : ", parse(invalid_input))
}

func parse(param string) bool {
	tags := []string{}

	fmt.Printf("Parsing %s \n", param)
	for _, c := range param {
		if strings.Contains(OPEN_TAGS, string(c)) {
			tags = append(tags, string(c))
		}
		if strings.Contains(CLOSED_TAGS, string(c)) {
			position := strings.Index(OPEN_TAGS, tags[len(tags)-1])
			if string(c) == string(CLOSED_TAGS[position]) {
				tags = tags[:len(tags)-1]
			}
		}
	}
	if len(tags) == 0 {
		return true
	}
	return false
}
