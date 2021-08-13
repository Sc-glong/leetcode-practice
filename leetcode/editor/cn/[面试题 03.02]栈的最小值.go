//请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(
//1)。 示例： MinStack minStack = new MinStack(); minStack.push(-2); minStack.push(0);
// minStack.push(-3); minStack.getMin();   --> 返回 -3. minStack.pop(); minStack.top
//();      --> 返回 0. minStack.getMin();   --> 返回 -2. Related Topics 栈 设计 
// 👍 53 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
     stack []int
     minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
   return MinStack{
       stack: []int{},
       minStack: []int{math.MaxInt64},
   }
}


func (this *MinStack) Push(x int)  {
       this.stack = append(this.stack,x)
       this.minStack = append(this.minStack,min(this.minStack[len(this.minStack)-1],x))
}


func (this *MinStack) Pop()  {
  this.stack = this.stack[:len(this.stack)-1]
  this.minStack = this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
   return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
   return this.minStack[len(this.minStack)-1]
}
func min(x,y int) int {
  if x < y {
    return x
  }
  return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
