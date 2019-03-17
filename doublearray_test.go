package doublearray

import (
	"math/rand"
	"testing"
)

var keys = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua", "Ut", "enim", "ad", "minim", "veniam", "quis", "nostrud", "exercitation", "ullamco", "laboris", "nisi", "ut", "aliquip", "ex", "ea", "commodo", "consequat", "Duis", "aute", "irure", "dolor", "in", "reprehenderit", "in", "voluptate", "velit", "esse", "cillum", "dolore", "eu", "fugiat", "nulla", "pariatur", "Excepteur", "sint", "occaecat", "cupidatat", "non", "proident", "sunt", "in", "culpa", "qui", "officia", "deserunt", "mollit", "anim", "id", "est", "laborum"}

var exactMatchSearchTests = []struct {
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

var containsMatchTests = []struct {
	in  string
	out bool
}{
	{"lorem ipsum is simply dummy text", true},
	{"Lorem is simply dummy", true},
	{"lorem ipsu is simply dummy txt", false},
}

var d DoubleArray

func init() {
	d = *New(keys)
}

func TestExactMatchSearch(t *testing.T) {
	for _, test := range exactMatchSearchTests {
		t.Run(test.in, func(t *testing.T) {
			result := d.ExactMatchSearch(test.in)

			if result != test.out {
				t.Errorf("got %t, want %t", result, test.out)
			}
		})

	}
}

func TestContainsMatch(t *testing.T) {
	for _, test := range containsMatchTests {
		t.Run(test.in, func(t *testing.T) {
			result := d.ContainsMatch(test.in)

			if result != test.out {
				t.Errorf("got %t, want %t", result, test.out)
			}
		})
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(keys)
	}
}

func BenchmarkExactMatchSearch(b *testing.B) {
	var s []string
	for i := 0; i < b.N; i++ {
		s = append(s, keys[rand.Intn(len(keys))])
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		d.ExactMatchSearch(s[i])
	}
}

func BenchmarkContainsMatch(b *testing.B) {
	text := "lorem ipsum is simply dummy text"
	for i := 0; i < b.N; i++ {
		d.ContainsMatch(text)
	}
}
