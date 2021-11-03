// https://play.golang.org/p/MtZCXGcESeF
// https://stackoverflow.com/questions/49065781/limit-size-of-the-priority-queue-for-gos-heap-interface-implementation

package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

const (
	n = 100
	m = 5
)

func main() {
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		x := i // n- i
		heap.Push(h, x)
		if h.Len() > m {
			heap.Pop(h)
		}
	}

	r := make([]int, h.Len())
	for i := len(r) - 1; i >= 0; i-- {
		r[i] = heap.Pop(h).(int)
	}
	fmt.Printf("%v\n", r)
}
