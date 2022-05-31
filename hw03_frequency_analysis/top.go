package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`(\p{Cyrillic}+-*\p{Cyrillic}*)|(\p{Latin}+-*\p{Latin}*)`)

type keyValue struct {
	key   string
	value int
}

func mapToSlice(inputMap map[string]int) []keyValue {
	keyValueSlice := make([]keyValue, len(inputMap))
	i := 0
	for k, v := range inputMap {
		keyValueSlice[i].key = k
		keyValueSlice[i].value = v
		i++
	}
	return keyValueSlice
}

func Top10(text string) []string {
	regexRes := re.FindAllString(text, -1)

	if len(regexRes) == 0 {
		return []string{}
	}

	wordsCounter := map[string]int{}

	for i, word := range regexRes {
		regexRes[i] = strings.ToLower(word)
	}

	for _, word := range regexRes {
		if wordsCounter[word] != 0 {
			wordsCounter[word]++
		} else {
			wordsCounter[word] = 1
		}
	}

	wordsCounterSlice := mapToSlice(wordsCounter)

	sort.Slice(wordsCounterSlice, func(i, j int) bool {
		if wordsCounterSlice[i].value != wordsCounterSlice[j].value {
			return wordsCounterSlice[i].value > wordsCounterSlice[j].value
		}
		return wordsCounterSlice[i].key < wordsCounterSlice[j].key
	})

	var topWordsLen int
	if len(wordsCounterSlice) >= 10 {
		topWordsLen = 10
	} else {
		topWordsLen = len(wordsCounterSlice)
	}

	topWords := make([]string, topWordsLen)

	for i := 0; i < topWordsLen; i++ {
		topWords[i] = wordsCounterSlice[i].key
	}

	return topWords
}
