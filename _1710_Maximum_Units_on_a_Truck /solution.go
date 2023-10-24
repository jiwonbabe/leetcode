package _1710_Maximum_Units_on_a_Truck_

import "container/heap"

type MaxHeap [][]int

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
	*h = append(*h, i.([]int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	l := old.Len()
	x := old[l-1]
	*h = old[0 : l-1]
	return x
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, b := range boxTypes {
		heap.Push(h, []int{b[0], b[1]})
	}
	units := 0
	for i := 0; i < truckSize; i++ {
		if h.Len() <= 0 {
			break
		}
		max := heap.Pop(h).([]int)
		max[0] -= 1
		units += max[1]
		if max[0] != 0 {
			heap.Push(h, max)
		}
	}

	return units
}
