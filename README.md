# doublearray

[![Build Status](https://travis-ci.org/iguchi1124/doublearray.svg?branch=master)](https://travis-ci.org/iguchi1124/doublearray)

Provides the `doublearray` package that implements a double-array trie tree.

## Documentation

The full documentation is available on [GoDoc](https://godoc.org/github.com/iguchi1124/doublearray).

## Example

```go
dict := []string{"apple", "orange", "pineapple"}
trie := doublearray.New(dict)
trie.ExactMatchSearch("apple") // true
trie.ExactMatchSearch("banana") // false
```

## License

MIT
