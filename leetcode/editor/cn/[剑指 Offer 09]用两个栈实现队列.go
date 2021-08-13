//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的
//功能。(若队列中没有元素，deleteHead 操作返回 -1 ) 
//
// 
//
// 示例 1： 
//
// 输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
// 
//
// 示例 2： 
//
// 输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
// 
//
// 提示： 
//
// 
// 1 <= values <= 10000 
// 最多会对 appendTail、deleteHead 进行 10000 次调用 
// 
// Related Topics 栈 设计 队列 
// 👍 286 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type CQueue struct {
     stackA []int
     stackB []int
}


func Constructor() CQueue {
     return CQueue {
          stackA: []int{},
          stackB: []int{},
     }
}


func (this *CQueue) AppendTail(value int)  {
     this.stackA = append(this.stackA,value)
}


func (this *CQueue) DeleteHead() int {
     for {
          if len(this.stackA) == 0 {
                 break
          }
         this.stackB = append(this.stackB,this.stackA[len(this.stackA)-1])
         this.stackA = this.stackA[:len(this.stackA)-1]
     }
     if len(this.stackB) == 0 {
       return -1
     }
     res := this.stackB[len(this.stackB)-1]
     this.stackB = this.stackB[:len(this.stackB)-1]
     for  {
       if len(this.stackB) ==  0 {
             break
        }
       this.stackA = append(this.stackA,this.stackB[len(this.stackB)-1])
       this.stackB = this.stackB[:len(this.stackB)-1]
     }
     return res
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)
