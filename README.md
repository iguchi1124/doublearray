# doublearray

[![Build Status](https://travis-ci.org/iguchi1124/doublearray.svg?branch=master)](https://travis-ci.org/iguchi1124/doublearray)

An implementation of trie tree for golang.

## Usage

```go
package main

import (
	"fmt"

	"github.com/iguchi1124/doublearray"
)

var dict = []string{"apple", "orange", "pineapple"}

func main() {
	trie := doublearray.New(dict)
	fmt.Println(trie.ExactMatchSearch("apple")) // `true`
	fmt.Println(trie.ExactMatchSearch("banana")) // `false`
}
```
