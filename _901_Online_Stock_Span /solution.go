package _901_Online_Stock_Span_

type StockSpanner struct {
	Hm    map[int]int
	Stack []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		Hm:    make(map[int]int),
		Stack: make([]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	cnt := 1
	for len(this.Stack) > 0 && this.Stack[len(this.Stack)-1] <= price {
		top := this.Stack[len(this.Stack)-1]
		cnt += this.Hm[top]
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
	this.Stack = append(this.Stack, price)
	this.Hm[price] = cnt

	return cnt
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
