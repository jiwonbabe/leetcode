package _973_K_Closest_Points_to_Origin_

import "container/heap"

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)
	for _, p := range points {
		p = append(p, findDist(p))
		heap.Push(h, p)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := [][]int{}
	for h.Len() > 0 {
		point := heap.Pop(h).([]int)
		result = append(result, point[:2])
	}
	return result
}

type MaxHeap [][]int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i][2] > h[j][2]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(i any) {
	*h = append(*h, i.([]int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	l := h.Len()
	x := old[l-1]
	*h = old[0 : l-1]
	return x
}

func findDist(point []int) int {
	return (point[0] * point[0]) + (point[1] * point[1])
}
