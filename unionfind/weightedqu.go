package unionfind

// WeightedQU is a weighted tree implementation of Quick-union algorithm for UnionFind problem
type WeightedQU struct {
	nodes, sizes []int
}

// NewWeightedQU creates new instance of WeightedQU
func NewWeightedQU(N int) UnionFind {
	nodes, sizes := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		nodes[i], sizes[i] = i, 1
	}
	return WeightedQU{nodes, sizes}
}

// Union joins pth and qth sets together
func (wq WeightedQU) Union(p, q int) {
	proot, qroot := wq.root(p), wq.root(q)

	if proot == qroot {
		return
	}

	if wq.sizes[proot] >= wq.sizes[qroot] {
		qroot, proot = proot, qroot
	}

	wq.nodes[proot] = qroot
	wq.sizes[qroot] += wq.sizes[proot]
}

// Connected checks if p and q are lying in the same set
func (wq WeightedQU) Connected(p, q int) bool {
	return wq.root(p) == wq.root(q)
}

func (wq WeightedQU) root(p int) int {
	for p != wq.nodes[p] {
		wq.nodes[p] = wq.nodes[wq.nodes[p]]
		p = wq.nodes[p]
	}
	return p
}
