package unionfind

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	N := 10
	ufs := []struct {
		name string
		UnionFind
	}{
		{"QuickFind", NewQuickFind(N)},
		{"QuickUnion", NewQuickUnion(N)},
		{"WeightedQU", NewWeightedQU(N)},
	}
	for _, uf := range ufs {
		for _, s := range []struct {
			q, p int
		}{
			{0, 2},
			{2, 3},
			{7, 9},
			{6, 8},
			{4, 6},
			{8, 9},
		} {
			uf.Union(s.p, s.q)
		}
	}

	testCases := []struct {
		q, p int
		want bool
	}{
		{6, 8, true},
		{4, 9, true},
		{3, 0, true},
		{2, 6, false},
		{1, 4, false},
		{7, 5, false},
	}

	for _, uf := range ufs {
		t.Run(uf.name, func(t *testing.T) {
			for _, tC := range testCases {
				got := uf.Connected(tC.p, tC.q)
				if got != tC.want {
					t.Errorf("TestUnionFind(%d, %d) == %t, want %t", tC.p, tC.q, got, tC.want)
				}
			}
		})
	}
}
