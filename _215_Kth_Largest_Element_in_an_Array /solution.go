package _215_Kth_Largest_Element_in_an_Array_

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(i any) {
	*h = append(*h, i.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	l := h.Len()
	x := old[l-1]
	*h = old[0 : l-1]
	return x
}
