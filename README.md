## Tokenizer

Tokenizer is a simple text tokenizer written in Go.

### Instalation

> go get github.com/flavioltonon/tokenizer

### Usage

```
// Create a tokenizer with a set of stopwords
t := tokenizer.New("foo", "bar", "baz")

// Tokenize a text
tokens := t.Tokenize("123 foo 456 bar 789 baz qux") // []string{"123", "456", "789", "qux"}
```