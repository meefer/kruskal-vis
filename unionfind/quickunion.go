package unionfind

// QuickUnion is a tree-based algorithm for UnionFind problem
type QuickUnion []int

// NewQuickUnion creates new instance of QuickUnion
func NewQuickUnion(N int) UnionFind {
	nodes := make(QuickUnion, N)
	for i := 0; i < N; i++ {
		nodes[i] = i
	}
	return nodes
}

func (qu QuickUnion) root(p int) int {
	for p != qu[p] {
		p = qu[p]
	}
	return p
}

// Union joins pth and qth sets together
func (qu QuickUnion) Union(p, q int) {
	proot, qroot := qu.root(p), qu.root(q)
	qu[qroot] = proot
}

// Connected checks if p and q are lying in the same set
func (qu QuickUnion) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}
