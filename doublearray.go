// Package doublearray implementation of trie tree
package doublearray

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
