package doublearray

type Iterator struct {
	Index       int
	doubleArray *DoubleArray
}

func (d *DoubleArray) Iterator() *Iterator {
	return &Iterator{0, d}
}

func (i *Iterator) HasNext(key rune) bool {
	idx := i.doubleArray.Nodes[i.Index].Base + int(key) - 1

	if len(i.doubleArray.Nodes) < idx+1 {
		return false
	}

	if i.doubleArray.Nodes[idx].Check == i.Index+1 {
		return true
	}

	return false
}

func (i *Iterator) Next(key rune) {
	i.Index = i.doubleArray.Nodes[i.Index].Base + int(key) - 1
}

func (i *Iterator) Node() *Node {
	return &i.doubleArray.Nodes[i.Index]
}
