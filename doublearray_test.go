package doublearray

import (
	"math/rand"
	"testing"
)

func TestExactMatchSearch(t *testing.T) {
	keys := []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua", "Ut", "enim", "ad", "minim", "veniam", "quis", "nostrud", "exercitation", "ullamco", "laboris", "nisi", "ut", "aliquip", "ex", "ea", "commodo", "consequat", "Duis", "aute", "irure", "dolor", "in", "reprehenderit", "in", "voluptate", "velit", "esse", "cillum", "dolore", "eu", "fugiat", "nulla", "pariatur", "Excepteur", "sint", "occaecat", "cupidatat", "non", "proident", "sunt", "in", "culpa", "qui", "officia", "deserunt", "mollit", "anim", "id", "est", "laborum"}
	tests := []struct {
		in  string
		out bool
	}{
		{"Lorem", true},
		{"ipsum", true},
		{"dolor", true},
		{"sit", true},
		{"amet", true},
		{"consectetur", true},
		{"adipiscing", true},
		{"elit", true},
		{"sed", true},
		{"do", true},
		{"eiusmod", true},
		{"tempor", true},
		{"incididunt", true},
		{"ut", true},
		{"labore", true},
		{"et", true},
		{"dolore", true},
		{"magna", true},
		{"aliqua", true},
		{"Lore", false},
		{"lorem", false},
		{"ipsu", false},
		{"olor", false},
		{"i", false},
	}

	doubleArray := *New(keys)

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			result := doubleArray.ExactMatchSearch(test.in)

			if result != test.out {
				t.Errorf("got %t, want %t", result, test.out)
			}
		})

	}
}

func TestCommonPrefixSearch(t *testing.T) {
	tests := []struct {
		in  string
		out CommonPrefixSearchResult
	}{
		{
			"I have an pineapple.",
			[]struct {
				Index, Len int
			}{
				{10, 9},
				{14, 5},
			},
		},
	}

	doubleArray := *New([]string{"apple", "orange", "pineapple"})

	for _, test := range tests {

		t.Run(test.in, func(t *testing.T) {
			ok := true
			results := doubleArray.CommonPrefixSearch(test.in)

			for i, result := range results {
				if test.out[i] != result {
					ok = false
				}
			}

			if !ok {
				t.Errorf("got %v, want %v", results, test.out)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	keys := []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua", "Ut", "enim", "ad", "minim", "veniam", "quis", "nostrud", "exercitation", "ullamco", "laboris", "nisi", "ut", "aliquip", "ex", "ea", "commodo", "consequat", "Duis", "aute", "irure", "dolor", "in", "reprehenderit", "in", "voluptate", "velit", "esse", "cillum", "dolore", "eu", "fugiat", "nulla", "pariatur", "Excepteur", "sint", "occaecat", "cupidatat", "non", "proident", "sunt", "in", "culpa", "qui", "officia", "deserunt", "mollit", "anim", "id", "est", "laborum"}
	for i := 0; i < b.N; i++ {
		New(keys)
	}
}

func BenchmarkExactMatchSearch(b *testing.B) {
	keys := []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua", "Ut", "enim", "ad", "minim", "veniam", "quis", "nostrud", "exercitation", "ullamco", "laboris", "nisi", "ut", "aliquip", "ex", "ea", "commodo", "consequat", "Duis", "aute", "irure", "dolor", "in", "reprehenderit", "in", "voluptate", "velit", "esse", "cillum", "dolore", "eu", "fugiat", "nulla", "pariatur", "Excepteur", "sint", "occaecat", "cupidatat", "non", "proident", "sunt", "in", "culpa", "qui", "officia", "deserunt", "mollit", "anim", "id", "est", "laborum"}
	doubleArray := *New(keys)

	var s []string
	for i := 0; i < b.N; i++ {
		s = append(s, keys[rand.Intn(len(keys))])
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		doubleArray.ExactMatchSearch(s[i])
	}
}
