package tokenizer

import "strings"

// Tokenizer is a text tokenizer.
type Tokenizer struct {
	stopWords    []string
	stopWordsMap map[string]struct{}
}

// New creates a new Tokenizer with an option set of stopwords
func New(stopWords ...string) *Tokenizer {
	return &Tokenizer{
		stopWords:    stopWords,
		stopWordsMap: mapWords(stopWords),
	}
}

// Tokenize extracts tokens from a textx, ignoring stop words defined in the Tokenizer
func (t *Tokenizer) Tokenize(text string) []string {
	var tokens []string

	for _, token := range strings.Fields(text) {
		if _, exists := t.stopWordsMap[token]; exists {
			continue
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func mapWords(words []string) map[string]struct{} {
	wordsMap := make(map[string]struct{}, len(words))

	for _, word := range words {
		wordsMap[word] = struct{}{}
	}

	return wordsMap
}
