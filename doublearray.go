// Package doublearray implementation of trie tree
package doublearray

import "sort"

const endKey rune = 0

// Node is a node of tree
type Node struct {
	Base  int
	Check int
}

// DoubleArray is a trie tree entity
type DoubleArray struct {
	Nodes []Node
}

// New builds trie tree
func New(keywords []string) *DoubleArray {
	d := DoubleArray{Nodes: make([]Node, 1)}
	d.build(keywords)

	return &d
}

// ExactMatchSearch returns that whether the string
// is contained in trie tree.
func (d *DoubleArray) ExactMatchSearch(s string) bool {
	if len(s) == 0 {
		return false
	}

	idx := 0
	for _, key := range s {
		nx := d.Nodes[idx].Base + int(key) - 1

		if len(d.Nodes) < nx+1 {
			return false
		}

		if d.Nodes[nx].Check == idx+1 {
			idx = nx
		} else {
			return false
		}
	}

	node := d.Nodes[d.Nodes[idx].Base+int(endKey)-1]
	return node.Check == idx+1 && node.Base == 0
}

// ContainsMatch returns that whether the string
// has any match part with trie tree.
func (d *DoubleArray) ContainsMatch(s string) bool {
	keys := []rune(s)

	for i := 0; i < len(keys); i++ {
		idx := 0

		for j := i; j < len(keys); j++ {
			k := keys[j]

			nx := d.Nodes[idx].Base + int(k) - 1
			if len(d.Nodes) < nx+1 {
				break
			}

			if d.Nodes[nx].Check == idx+1 {
				idx = nx

				node := d.Nodes[d.Nodes[idx].Base+int(endKey)-1]
				if node.Check == idx+1 && node.Base == 0 {
					return true
				}
			} else {
				break
			}
		}
	}

	return false
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
