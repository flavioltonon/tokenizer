package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		stopWords []string
	}

	tests := []struct {
		name string
		args args
		want *Tokenizer
	}{
		{
			name: "Given a set of stop words, a new Tokenizer should be created",
			args: args{
				stopWords: []string{"foo", "bar", "baz"},
			},
			want: &Tokenizer{
				stopWords: []string{"foo", "bar", "baz"},
				stopWordsMap: map[string]struct{}{
					"foo": {},
					"bar": {},
					"baz": {},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ObjectsAreEqual(tt.want, New(tt.args.stopWords...))
		})
	}
}

func TestTokenizer_Tokenize(t *testing.T) {
	type args struct {
		text string
	}

	tests := []struct {
		name      string
		tokenizer *Tokenizer
		args      args
		want      []string
	}{
		{
			name:      "Given a text Tokenizer.Tokenize should return a set of string tokens, ignoring any stopwords",
			tokenizer: New("foo", "bar", "baz"),
			args: args{
				text: "123 foo 456 bar 789 baz qux",
			},
			want: []string{"123", "456", "789", "qux"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, tt.tokenizer.Tokenize(tt.args.text))
		})
	}
}

func Test_mapWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			name: "mapWords should create a map where the keys are each one of the input words",
			args: args{
				words: []string{"foo", "bar", "baz"},
			},
			want: map[string]struct{}{
				"foo": {},
				"bar": {},
				"baz": {},
			},
		},
		{
			name: "If no words are provided, an empty map should be returned",
			args: args{
				words: []string{},
			},
			want: map[string]struct{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mapWords(tt.args.words))
		})
	}
}
