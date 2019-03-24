// Time:  O(n)
// Space: O(1)

type MinStack struct {
	elements_   []int64
	minElement_ int64
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.elements_) == 0 {
		this.minElement_ = int64(x)
		this.elements_ = append(this.elements_, 0)
	} else {
		this.elements_ = append(this.elements_, int64(x)-this.minElement_)
		if this.elements_[len(this.elements_)-1] < 0 {
			this.minElement_ = int64(x)
		}
	}
}

func (this *MinStack) Pop() {
	if this.elements_[len(this.elements_)-1] < 0 {
		this.minElement_ = this.minElement_ - this.elements_[len(this.elements_)-1]
	}
	this.elements_ = this.elements_[:len(this.elements_)-1]
}

func (this *MinStack) Top() int {
	if this.elements_[len(this.elements_)-1] > 0 {
		return int(this.elements_[len(this.elements_)-1] + this.minElement_)
	}
	return int(this.minElement_)
}

func (this *MinStack) GetMin() int {
	return int(this.minElement_)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */