package cmd

import "fmt"

func Convert(wordMap map[string]string, words []string) []byte {
	var result []byte

	for _, word := range words {
		w := hanToZen(wordMap, word)
		fmt.Printf("old: %s, new: %s\n", word, w)
		result = append(result, w...)
	}

	return result
}

func hanToZen(wordMap map[string]string, word string) string {
	if val, ok := wordMap[word]; ok {
		return val
	}

	return word
}
