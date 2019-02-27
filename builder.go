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
	cursor := 1

	for _, key := range keyword {
		d.commit(&cursor, key)
	}

	d.commit(&cursor, endKey)
}

func (d *DoubleArray) commit(c *int, key rune) {
	s := 1
	if d.Nodes.At(*c).Base != 0 {
		s = d.Nodes.At(*c).Base
	}

	next := s + int(key)

	if len(d.Nodes) < next {
		nodes := make([]Node, next)
		copy(nodes, d.Nodes)
		d.Nodes = nodes
	}

	n := d.Nodes.At(next)

	if n.Check == *c {
		*c = next
	} else if n.Check == 0 && n.Base == 0 {
		d.Nodes.At(*c).Base = s
		d.Nodes.At(next).Check = *c
		*c = next
	} else {
		d.rebase(c, key)
	}
}

func (d *DoubleArray) rebase(c *int, key rune) {
	b := d.Nodes.At(*c).Base

	nodes := make(map[int]Node)
	for i, n := range d.Nodes {
		if n.Check == *c {
			nodes[i+1] = n
		}
	}

	keys := []rune{key}
	for c := range nodes {
		keys = append(keys, rune(c-b))
	}

	base := b
	for {
		base++
		ok := true

		for _, k := range keys {
			nx := base + int(k)
			if len(d.Nodes) < nx {
				n := make([]Node, nx)
				copy(n, d.Nodes)
				d.Nodes = n
			}

			ok = ok && d.Nodes.At(nx).Check == 0 && d.Nodes.At(nx).Base == 0
		}

		if ok {
			break
		}
	}

	d.Nodes.At(*c).Base = base
	for _, k := range keys {
		d.Nodes.At(int(k) + base).Check = *c
	}

	for cu, n := range nodes {
		k := cu - b
		nx := base + k

		d.Nodes.At(nx).Base = n.Base

		for i, nd := range d.Nodes {
			if nd.Check == cu {
				d.Nodes[i].Check = nx

				d.Nodes.At(cu).Base = 0
				d.Nodes.At(cu).Check = 0
			}
		}

	}

	*c = base + int(key)
}
