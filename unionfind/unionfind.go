// Package unionfind implements an Union-find data structure that tracks a set of elements partitioned into a number of disjoint subsets.
// Details can be found on https://en.wikipedia.org/wiki/Disjoint-set_data_structure
package unionfind

// UnionFind data structure helps to solve a problem of determining if the elements have smth in common by splitting them into disjoint sets
type UnionFind interface {
	Union(p, q int)
	Connected(p, q int) bool
}

// NewUnionFind returns new disjoint-set data structure (union-find) instance
func NewUnionFind(N int) UnionFind {
	return NewWeightedQU(N)
}
