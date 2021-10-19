//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。 
//
// 
//
// 示例: 
//
// MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.
// 
//
// 
//
// 提示： 
//
// 
// 各函数的调用总次数不超过 20000 次 
// 
//
// 
//
// 注意：本题与主站 155 题相同：https://leetcode-cn.com/problems/min-stack/ 
// Related Topics 栈 设计 👍 197 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
    arrMin []int
    arrB []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack {
           arrMin: []int{},
           arrB: []int{},
     }
}


func (this *MinStack) Push(x int)  {
   this.arrB = append(this.arrB,x)
   if len(this.arrMin) == 0 {
     this.arrMin = append(this.arrMin,x)
   } else {
       if this.arrMin[len(this.arrMin)-1] > x {
           this.arrMin = append(this.arrMin,x)
       } else {
           this.arrMin = append(this.arrMin,this.arrMin[len(this.arrMin)-1])
       }
   }
}


func (this *MinStack) Pop()  {
   this.arrB = this.arrB[:len(this.arrB)-1]
   this.arrMin = this.arrMin[:len(this.arrMin)-1]
}


func (this *MinStack) Top() int {
   return this.arrB[len(this.arrB)-1]
}


func (this *MinStack) Min() int {
   return this.arrMin[len(this.arrMin)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
//leetcode submit region end(Prohibit modification and deletion)
