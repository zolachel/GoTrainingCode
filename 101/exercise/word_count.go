package main

import (
	"fmt"
	"strings"
)

func wordCount(sentence string) map[string]int {
	sentence = strings.ReplaceAll(strings.ReplaceAll(sentence, ",", ""), ".", "")

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
