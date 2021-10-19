//三合一。描述如何只用一个数组来实现三个栈。 
//
// 你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。
//stackNum表示栈下标，value表示压入的值。 
//
// 构造函数会传入一个stackSize参数，代表每个栈的大小。 
//
// 示例1: 
//
// 
// 输入：
//["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
//[[1], [0, 1], [0, 2], [0], [0], [0], [0]]
// 输出：
//[null, null, null, 1, -1, -1, true]
//说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。
// 
//
// 示例2: 
//
// 
// 输入：
//["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
//[[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
// 输出：
//[null, null, null, null, 2, 1, -1, -1]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= stackNum <= 2 
// 
// Related Topics 栈 设计 数组 👍 37 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type TripleInOne struct {
     arr []int
     inxs []int
     length int
}


func Constructor(stackSize int) TripleInOne {
     return TripleInOne{
          arr: make([]int, stackSize *3),
          inxs: []int{0, 1, 2},
          length: stackSize*3,
     }
}


func (this *TripleInOne) Push(stackNum int, value int)  {
     if this.inxs[stackNum] >= this.length {
         return
     }
     this.arr[this.inxs[stackNum]] = value
     this.inxs[stackNum] += 3
}


func (this *TripleInOne) Pop(stackNum int) int {
     if this.inxs[stackNum] < 3 {
        return -1
     }
     this.inxs[stackNum] -= 3
     return this.arr[this.inxs[stackNum]]
}


func (this *TripleInOne) Peek(stackNum int) int {
   if this.inxs[stackNum] < 3 {
      return -1
   }
   return this.arr[this.inxs[stackNum]-3]
}


func (this *TripleInOne) IsEmpty(stackNum int) bool {
   return this.inxs[stackNum] < 3
}


/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
//leetcode submit region end(Prohibit modification and deletion)
