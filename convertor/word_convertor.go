package convertor

type WordMap map[string]string

type WordConvertor struct {
	wordMap WordMap
}

func NewWordConvertor(wordMap WordMap) *WordConvertor {
	return &WordConvertor{
		wordMap,
	}
}

func (c *WordConvertor) Convert(words []string) []byte {
	var result []byte

	for _, word := range words {
		w := c.hanToZen(word)
		result = append(result, w...)
	}

	return result
}

func (c *WordConvertor) hanToZen(word string) string {
	if val, ok := c.wordMap[word]; ok {
		return val
	}

	return word
}
