package cmd

import "fmt"

type WordMap map[string]string

func Convert(wordMap WordMap, words []string) []byte {
	var result []byte

	for _, word := range words {
		w := hanToZen(wordMap, word)
		fmt.Printf("old: %s, new: %s\n", word, w)
		result = append(result, w...)
	}

	return result
}

func hanToZen(wordMap WordMap, word string) string {
	if val, ok := wordMap[word]; ok {
		return val
	}

	return word
}
