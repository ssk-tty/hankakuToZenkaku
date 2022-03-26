package convertor

type WordMap map[string]string

type WordConvertor struct {
	words *[]string
}

func NewWordConvertor(words *[]string) *WordConvertor {
	return &WordConvertor{
		words,
	}
}

func (wc *WordConvertor) ToZen() []byte {
	var result []byte

	zenMap := GetZenMap()

	for _, word := range *wc.words {
		result = wc.convertFromHalfToZen(zenMap, word, result)
	}

	return result
}

func (wc *WordConvertor) lenIsGraterThan(word string, base int) bool {
	return len(word) > base
}

func (wc *WordConvertor) convertFromHalfToZen(wordMap WordMap, word string, stack []byte) []byte {
	if wc.lenIsGraterThan(word, 1) {
		return append(stack, word...)
	}

	return append(stack, wc.mapToWordMap(wordMap, word)...)
}

func (wc *WordConvertor) mapToWordMap(wordMap WordMap, word string) string {
	if val, ok := wordMap[word]; ok {
		return val
	}

	return word
}
