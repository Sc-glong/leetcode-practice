//请判断一个链表是否为回文链表。 
//
// 示例 1: 
//
// 输入: 1->2
//输出: false 
//
// 示例 2: 
//
// 输入: 1->2->2->1
//输出: true
// 
//
// 进阶： 
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？ 
// Related Topics 链表 双指针 
// 👍 1025 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
   // 双指针
//    ars := []int{}
//    for ; head != nil;head = head.Next {
//       ars = append(ars,head.Val)
//    }
//    count := len(ars)
//    for i:=0;i<count;i++ {
//        if ars[i] != ars[count-i-1] {
//           return false
//        }
//    }
//    return true

      //递归
      //尾递归的形式，在归的规程中，用另一个指针从开始遍历，比较两者的值

      //先用快慢指针找到中间节点
      //再反转中间节点后的链表
      //比较
      //还原
      if head == nil {
         return true
      }
      first := endOfHalfList(head)
      secondList := reverseList(first.Next)
      p1 := head
      p2 := secondList
      result := true
      for ;result && p2 != nil;p1=p1.Next,p2=p2.Next {
         if p1.Val != p2.Val {
            result = false
         }
      }
      first.Next = reverseList(secondList)
      return result
}
//反转链表
func reverseList(head *ListNode) *ListNode {
     var pre *ListNode
     cur := head
     for ;cur.Next != nil; cur = cur.Next {
          next := cur.Next
          cur.Next = pre
          pre = cur
          cur = next
     }
     return pre
}

//快慢指针得到中间节点
func endOfHalfList(head *ListNode) *ListNode {
     low := head
     fast:= head
     for fast.Next!=nil &&fast.Next.Next!= nil {
         low = low.Next
         fast = fast.Next.Next
     }
     return low
}
//leetcode submit region end(Prohibit modification and deletion)
