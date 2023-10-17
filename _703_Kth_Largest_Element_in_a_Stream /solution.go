package _703_Kth_Largest_Element_in_a_Stream_

import "container/heap"

type KthLargest struct {
	K    int
	Heap *Heap
}

func Constructor(k int, nums []int) KthLargest {
	h := &Heap{}
	heap.Init(h)

	kl := KthLargest{
		K:    k,
		Heap: h,
	}

	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.Heap, val)
	if this.Heap.Len() > this.K {
		heap.Pop(this.Heap)
	}
	return (*this.Heap)[0]
}

type Heap []int

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Heap) Push(i any) {
	*h = append(*h, i.(int))
}
func (h *Heap) Pop() any {
	old := *h
	len := h.Len()
	x := old[len-1]
	*h = old[0 : len-1]
	return x
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
