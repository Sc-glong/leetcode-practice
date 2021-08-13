//编写一个函数，检查输入的链表是否是回文的。 
//
// 
//
// 示例 1： 
//
// 输入： 1->2
//输出： false 
// 
//
// 示例 2： 
//
// 输入： 1->2->2->1
//输出： true 
// 
//
// 
//
// 进阶： 
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？ 
// Related Topics 栈 递归 链表 双指针 
// 👍 76 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
      if head == nil {
        return true
      }
      slow, fast := head, head
      for fast.Next != nil && fast.Next.Next != nil {
          slow = slow.Next
          fast = fast.Next.Next
      }
      tmp := slow.Next
      slow.Next = nil
      tmp = reverse(tmp)
      for tmp != nil && head != nil {
          if head.Val != tmp.Val {
             return false
          }
          tmp = tmp.Next
          head = head.Next
      }
      return true
}
func reverse(head *ListNode) *ListNode {
      var pre *ListNode = nil
      for head != nil {
          tmp := head.Next
          head.Next = pre
          pre = head
          head = tmp
      }
      return pre
}
//leetcode submit region end(Prohibit modification and deletion)
