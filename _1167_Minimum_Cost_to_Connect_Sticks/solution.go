package _1167_Minimum_Cost_to_Connect_Sticks

import "container/heap"

func connectSticks(sticks []int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, s := range sticks {
		heap.Push(h, s)
	}

	cost := 0
	for h.Len() > 1 {
		s := 0
		for i := 0; i < 2; i++ {
			s += heap.Pop(h).(int)
		}
		heap.Push(h, s)
		cost += s
	}

	return cost
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
	len := h.Len()
	x := old[len-1]
	*h = old[0 : len-1]
	return x
}
