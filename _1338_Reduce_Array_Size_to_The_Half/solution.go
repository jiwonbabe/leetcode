package _1338_Reduce_Array_Size_to_The_Half

import "container/heap"

func minSetSize(arr []int) int {
	h := &MaxHeap{}
	heap.Init(h)
	occurs := make(map[int]int)
	for _, e := range arr {
		occurs[e]++
	}
	for k, v := range occurs {
		heap.Push(h, [2]int{k, v})
	}
	length := len(arr)
	removed := 0
	size := 0
	for length-removed > length/2 {
		max := heap.Pop(h).([2]int)
		removed += max[1]
		size++
	}

	return size
}

type MaxHeap [][2]int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(i any) {
	*h = append(*h, i.([2]int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	l := h.Len()
	x := old[l-1]
	*h = old[0 : l-1]
	return x
}
