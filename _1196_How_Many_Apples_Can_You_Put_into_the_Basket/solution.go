package _1196_How_Many_Apples_Can_You_Put_into_the_Basket

import "container/heap"

func maxNumberOfApples(weight []int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, w := range weight {
		heap.Push(h, w)
	}

	cnt := 0
	w := 0
	for h.Len() > 0 && w <= 5000 {
		min := heap.Pop(h).(int)
		w += min
		if w > 5000 {
			break
		}
		cnt++
	}

	return cnt
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
