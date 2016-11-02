// unionfind implements one path variant of the path compression union find algorithm.
// This is inspired from the Algorithms, Part I course @ http://www.coursera.org.
// This is given by Princeton university by Kevin Wayna and Robert Sedgewick
package unionfind

import "fmt"
//
//	Unionfind has many applications. The two main operations of use are -
//	find - determine which subset a particular element belongs to.
//	union - join two subsets

type UnionFind struct {
	arr   []int
	count int
}

// New initializes a new UnionFind struct and returns a pointer.
func New(count int) *UnionFind {

	id := make([]int, count)
	for i := 0; i < count; i++ {
		id[i] = i
	}
	return &UnionFind{arr: id, count: count}
}

func (uf *UnionFind) Find(p int) int {
	return uf.arr[p]
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.arr[p] == uf.arr[q]
}

func (uf *UnionFind) Union(p, q int) {
	if uf.Connected(p, q) {
		return
	}
	p_id := uf.arr[p]
	for i := 0; i < len(uf.arr); i++ {
		if uf.arr[i] == p_id {
			uf.arr[i] = uf.arr[q]
		}
	}
	uf.count--
}

func (uf *UnionFind) GetNumClusters() int {
	return uf.count
}

func (uf *UnionFind) PrintClusters() {
	for i := 0; i < len(uf.arr); i++ {
		fmt.Println(uf.arr[i])
	} 
}
