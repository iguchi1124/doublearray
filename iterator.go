package doublearray

type iterator struct {
	index       int
	doubleArray *DoubleArray
}

func (d *DoubleArray) iterator() *iterator {
	return &iterator{0, d}
}

func (i *iterator) hasNext(key rune) bool {
	idx := i.doubleArray.Nodes[i.index].Base + int(key) - 1

	if len(i.doubleArray.Nodes) < idx+1 {
		return false
	}

	if i.doubleArray.Nodes[idx].Check == i.index+1 {
		return true
	}

	return false
}

func (i *iterator) next(key rune) {
	i.index = i.doubleArray.Nodes[i.index].Base + int(key) - 1
}

func (i *iterator) isLeaf() bool {
	if i.hasNext(endKey) {
		return i.doubleArray.Nodes[i.node().Base+int(endKey)-1].Base == 0
	}

	return false
}

func (i *iterator) node() *Node {
	return &i.doubleArray.Nodes[i.index]
}
