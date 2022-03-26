package convertor

type WordMap map[string]string

type WordConvertor struct {
	words []string
}

func NewWordConvertor(words []string) *WordConvertor {
	return &WordConvertor{
		words,
	}
}

func lenIsGraterThan(word string, base int) bool {
	return len(word) > base
}

func (wc *WordConvertor) ToZen() []byte {
	var result []byte

	zenMap := GetZenMap()

	for _, word := range wc.words {
		if lenIsGraterThan(word, 1) {
			result = append(result, word...)
			continue
		}

		result = append(result, wc.mapping(zenMap, word)...)
	}

	return result
}

func (wc *WordConvertor) mapping(wordMap WordMap, word string) string {
	if val, ok := wordMap[word]; ok {
		return val
	}

	return word
}
