// Package doublearray implementation of trie tree
package doublearray

const endKey rune = 0

// DoubleArray is a trie tree entity
type DoubleArray struct {
	Nodes
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

	cursor := 1
	for _, key := range s {
		next := d.Nodes.At(cursor).Base + int(key)

		if len(d.Nodes) < next {
			return false
		}

		if d.Nodes.At(next).Check == cursor {
			cursor = next
		} else {
			return false
		}
	}

	lastNode := d.Nodes.At(d.Nodes.At(cursor).Base + int(endKey))
	return lastNode.Check == cursor && lastNode.Base == 0
}

// ContainsMatch returns that whether the string
// has any match part with trie tree.
func (d *DoubleArray) ContainsMatch(s string) bool {
	keys := []rune(s)

	for i := 0; i < len(keys); i++ {
		cursor := 1

		for j := i; j < len(keys); j++ {
			k := keys[j]

			next := d.Nodes.At(cursor).Base + int(k)
			if len(d.Nodes) < next {
				break
			}

			if d.Nodes.At(next).Check == cursor {
				cursor = next

				node := d.Nodes.At(d.Nodes.At(cursor).Base + int(endKey))
				if node.Check == cursor && node.Base == 0 {
					return true
				}
			} else {
				break
			}
		}
	}

	return false
}
