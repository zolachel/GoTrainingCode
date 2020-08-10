package main

import (
	"fmt"
	"regexp"
	"strings"
)

func replaceSpecialCharacters(sentence string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")

	return reg.ReplaceAllString(sentence, "")
}

func wordCount(sentence string) map[string]int {
	sentence = replaceSpecialCharacters(sentence)

	splitedSentence := strings.Split(sentence, " ")

	result := make(map[string]int)

	for _, word := range splitedSentence {
		count, ok := result[word]

		if !ok {
			result[word] = 1
		} else {
			result[word] = count + 1
		}
	}

	return result
}

func main() {
	fmt.Println(wordCount("If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."))
}
