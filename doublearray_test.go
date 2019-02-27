package doublearray

import (
	"math/rand"
	"testing"
)

var (
	keys        = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua", "Ut", "enim", "ad", "minim", "veniam", "quis", "nostrud", "exercitation", "ullamco", "laboris", "nisi", "ut", "aliquip", "ex", "ea", "commodo", "consequat", "Duis", "aute", "irure", "dolor", "in", "reprehenderit", "in", "voluptate", "velit", "esse", "cillum", "dolore", "eu", "fugiat", "nulla", "pariatur", "Excepteur", "sint", "occaecat", "cupidatat", "non", "proident", "sunt", "in", "culpa", "qui", "officia", "deserunt", "mollit", "anim", "id", "est", "laborum"}
	validKeys   = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua"}
	invalidKeys = []string{"Lore", "ipsu", "olor", "it", "met", "conctetur", "adiping"}
)

var d DoubleArray

func init() {
	d = *New(keys)
}

func TestExactMatchSearch(t *testing.T) {
	ok := true
	for _, key := range validKeys {
		ok = ok && d.ExactMatchSearch(key)
	}

	for _, key := range invalidKeys {
		ok = ok && !d.ExactMatchSearch(key)
	}

	if !ok {
		t.Fail()
	}
}

func TestContainsMatch(t *testing.T) {
	var ok bool

	text := "lorem ipsum is simply dummy text"
	ok = d.ContainsMatch(text)

	invalidTexts := []string{"lorem is simply dummy txt", "lorem ipsu is simply dummy txt"}
	for _, text := range invalidTexts {
		ok = ok && !d.ContainsMatch(text)
	}

	if !ok {
		t.Fail()
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
