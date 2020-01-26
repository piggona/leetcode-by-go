package minStack

type MinStack struct {
	Data []int
	Min  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	if len(this.Min) == 0 || this.Min[len(this.Min)-1] >= x {
		this.Min = append(this.Min, x)
	}
}

func (this *MinStack) Pop() {
	top := this.Data[len(this.Data)-1]
	this.Data = this.Data[:(len(this.Data) - 1)]
	if top == this.Min[len(this.Min)-1] {
		this.Min = this.Min[:(len(this.Min) - 1)]
	}
}

func (this *MinStack) Top() int {
	if len(this.Data) == 0 {
		return 0
	}
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Min) == 0 {
		return 0
	}
	return this.Min[len(this.Min)-1]
}
