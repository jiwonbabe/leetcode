package _346_Moving_Average_from_Data_Stream

type MovingAverage struct {
	Size    int
	Nums    []int
	CurrSum int
	Ptr     int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		Size:    size,
		Nums:    make([]int, 0, 10000),
		CurrSum: 0,
		Ptr:     0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.Nums) >= this.Size {
		this.CurrSum -= this.Nums[this.Ptr]
		this.Ptr++
	}
	this.CurrSum += val
	this.Nums = append(this.Nums, val)
	return (float64)(this.CurrSum) / (float64)(len(this.Nums[this.Ptr:]))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
