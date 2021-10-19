//输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。 
//
// 
//
// 参考以下这颗二叉搜索树： 
//
//      5
//    / \
//   2   6
//  / \
// 1   3 
//
// 示例 1： 
//
// 输入: [1,6,3,2,5]
//输出: false 
//
// 示例 2： 
//
// 输入: [1,3,2,6,5]
//输出: true 
//
// 
//
// 提示： 
//
// 
// 数组长度 <= 1000 
// 
// Related Topics 栈 树 二叉搜索树 递归 二叉树 单调栈 
// 👍 324 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func verifyPostorder(postorder []int) bool {
    return recur(postorder, 0, len(postorder)-1)
}
func recur(postorder []int, i,j int) bool {
  if i >= j {
     return true
  }
  m := i
  for postorder[m] < postorder[j] {
      m++
  }
  p := m
  for postorder[p] > postorder[j] {
       p++
  }
  return p == j && recur(postorder,i,m-1) && recur(postorder,m,j-1)
}
//leetcode submit region end(Prohibit modification and deletion)
