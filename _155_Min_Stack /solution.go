package _155_Min_Stack_

type MinStack struct {
	Main     []int
	MinTrack [][]int
}

func Constructor() MinStack {
	return MinStack{
		Main:     []int{},
		MinTrack: [][]int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Main = append(this.Main, val)
	n := len(this.MinTrack)
	if n > 0 {
		top := this.MinTrack[n-1][0]
		if top > val {
			this.MinTrack = append(this.MinTrack, []int{val, 1})
		} else {
			this.MinTrack[n-1][1]++
		}
	} else {
		this.MinTrack = append(this.MinTrack, []int{val, 1})
	}
}

func (this *MinStack) Pop() {
	this.Main = this.Main[0 : len(this.Main)-1]
	n := len(this.MinTrack)
	if this.MinTrack[n-1][1] == 1 {
		this.MinTrack = this.MinTrack[0 : n-1]
	} else {
		this.MinTrack[n-1][1]--
	}
}

func (this *MinStack) Top() int {
	return this.Main[len(this.Main)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinTrack[len(this.MinTrack)-1][0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
