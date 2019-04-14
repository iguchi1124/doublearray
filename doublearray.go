// Package doublearray implementation of trie tree
package doublearray

import "sort"

const endKey rune = 0

// Node is a node of dobule-array trie tree.
type Node struct {
	Base  int
	Check int
}

// DoubleArray is a dobule-array trie tree.
type DoubleArray struct {
	Nodes []Node
}

// New builds a dobule-array trie tree from keywords.
func New(keywords []string) *DoubleArray {
	d := DoubleArray{Nodes: make([]Node, 1)}
	d.build(keywords)

	return &d
}

// BreadthFirstSearchFn is used to BreadthFirstSearch function argument.
// Takes matched data leading index and its length, returning if iteration should be continue.
type BreadthFirstSearchFn func(index int, length int) bool

// BreadthFirstSearch is used to breadth-first-search for trie tree from string.
func (d *DoubleArray) BreadthFirstSearch(s string, fn BreadthFirstSearchFn) {
	for i := 0; i < len(s); i++ {
		iterator := d.iterator()
		for j, key := range s[i:] {
			if iterator.hasNext(key) {
				iterator = iterator.next(key)

				if iterator.isLeaf() && !fn(i, j+1) {
					break
				}
			} else {
				break
			}
		}
	}
}

// ExactMatchSearch returning if string contains exactly matched keyword in trie tree.
func (d *DoubleArray) ExactMatchSearch(s string) bool {
	if len(s) == 0 {
		return false
	}

	i := d.iterator()
	for _, key := range s {
		if i.hasNext(key) {
			i = i.next(key)
		} else {
			return false
		}
	}

	return i.isLeaf()
}

// CommonPrefixSearchResult is used to CommonPrefixSearch function return value.
type CommonPrefixSearchResult []struct {
	Index int
	Len   int
}

// CommonPrefixSearch returning all matched keywords with trie tree.
func (d *DoubleArray) CommonPrefixSearch(s string) CommonPrefixSearchResult {
	var results CommonPrefixSearchResult

	d.BreadthFirstSearch(s, func(index int, length int) bool {
		results = append(results, struct{ Index, Len int }{index, length})

		return true
	})

	return results
}

// ContainsMatch returning if string contains matched keyword with trie tree at least one.
func (d *DoubleArray) ContainsMatch(s string) (ok bool) {
	d.BreadthFirstSearch(s, func(index int, length int) bool {
		ok = true
		return false
	})

	return
}

func (d *DoubleArray) build(keywords []string) {
	sort.Strings(keywords)

	for _, keyword := range keywords {
		if len(keyword) > 0 {
			d.insert(keyword)
		}
	}
}

func (d *DoubleArray) insert(keyword string) {
	keys := []rune(keyword)
	keys = append(keys, endKey)

	i := 0
	for _, key := range keys {
		nx := d.setBase(i, key)
		d.Nodes[nx].Check = i + 1
		i = nx
	}
}

func (d *DoubleArray) setBase(idx int, key rune) int {
	base := 1
	if d.Nodes[idx].Base != 0 {
		base = d.Nodes[idx].Base
	}

	nx := base + int(key) - 1

	if len(d.Nodes) < nx+1 {
		nodes := make([]Node, nx+1)
		copy(nodes, d.Nodes)
		d.Nodes = nodes
	}

	n := d.Nodes[nx]
	if n.Check == 0 && n.Base == 0 {
		d.Nodes[idx].Base = base
	} else if n.Check != idx+1 {
		nx = d.resetBase(idx, key)
	}

	return nx
}

func (d *DoubleArray) resetBase(idx int, key rune) int {
	b := d.Nodes[idx].Base

	nodes := make(map[int]Node)
	for i, n := range d.Nodes {
		if n.Check == idx+1 {
			nodes[i] = n
		}
	}

	keys := []rune{key}
	for i := range nodes {
		keys = append(keys, rune(i+1-b))
	}

	base := b
	for {
		base++
		ok := true

		for _, k := range keys {
			nx := base + int(k) - 1
			if len(d.Nodes) < nx+1 {
				n := make([]Node, nx+1)
				copy(n, d.Nodes)
				d.Nodes = n
			}

			ok = ok && d.Nodes[nx].Check == 0 && d.Nodes[nx].Base == 0
		}

		if ok {
			break
		}
	}

	d.Nodes[idx].Base = base

	for ni, n := range nodes {
		nx := base + ni - b

		d.Nodes[nx].Base = n.Base
		d.Nodes[nx].Check = idx + 1

		for i, nd := range d.Nodes {
			if nd.Check == ni+1 {
				d.Nodes[i].Check = nx + 1

				d.Nodes[ni].Base = 0
				d.Nodes[ni].Check = 0
			}
		}

	}

	return base + int(key) - 1
}
