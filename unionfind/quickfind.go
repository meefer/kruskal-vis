package unionfind

// QuickFind is a greedy algorithm for UnionFind problem
type QuickFind []int

// NewQuickFind creates new instance of QuickFind
func NewQuickFind(N int) UnionFind {
	nodes := make(QuickFind, N)
	for i := 0; i < N; i++ {
		nodes[i] = i
	}
	return nodes
}

// Union joins pth and qth sets together
func (qf QuickFind) Union(p, q int) {
	pnode, qnode := qf[p], qf[q]
	for i := 0; i < len(qf); i++ {
		if qf[i] == qnode {
			qf[i] = pnode
		}
	}
}

// Connected checks if p and q are lying in the same set
func (qf QuickFind) Connected(p, q int) bool {
	return qf[p] == qf[q]
}
