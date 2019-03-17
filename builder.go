package doublearray

import "sort"

func (d *DoubleArray) build(keywords []string) {
	sort.Strings(keywords)

	for _, keyword := range keywords {
		if len(keyword) > 0 {
			d.insert(keyword)
		}
	}
}

func (d *DoubleArray) insert(keyword string) {
	index := 0
	for _, key := range keyword {
		d.commit(&index, key)
	}

	d.commit(&index, endKey)
}

func (d *DoubleArray) commit(idx *int, key rune) {
	s := 1
	if d.Nodes[*idx].Base != 0 {
		s = d.Nodes[*idx].Base
	}

	nx := s + int(key) - 1

	if len(d.Nodes) < nx+1 {
		nodes := make([]Node, nx+1)
		copy(nodes, d.Nodes)
		d.Nodes = nodes
	}

	n := d.Nodes[nx]

	if n.Check == *idx+1 {
		*idx = nx
	} else if n.Check == 0 && n.Base == 0 {
		d.Nodes[*idx].Base = s
		d.Nodes[nx].Check = *idx + 1
		*idx = nx
	} else {
		d.rebase(idx, key)
	}
}

func (d *DoubleArray) rebase(idx *int, key rune) {
	b := d.Nodes[*idx].Base

	nodes := make(map[int]Node)
	for i, n := range d.Nodes {
		if n.Check == *idx+1 {
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

	d.Nodes[*idx].Base = base
	for _, k := range keys {
		d.Nodes[int(k)+base-1].Check = *idx + 1
	}

	for ni, n := range nodes {
		k := ni + 1 - b
		nx := base + k - 1

		d.Nodes[nx].Base = n.Base

		for i, nd := range d.Nodes {
			if nd.Check == ni+1 {
				d.Nodes[i].Check = nx + 1

				d.Nodes[ni].Base = 0
				d.Nodes[ni].Check = 0
			}
		}

	}

	*idx = base + int(key) - 1
}
