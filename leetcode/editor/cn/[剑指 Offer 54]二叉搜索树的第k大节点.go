//给定一棵二叉搜索树，请找出其中第k大的节点。 
//
// 
//
// 示例 1: 
//
// 输入: root = [3,1,4,null,2], k = 1
//   3
//  / \
// 1   4
//  \
//   2
//输出: 4 
//
// 示例 2: 
//
// 输入: root = [5,3,6,2,4,null,null,1], k = 3
//       5
//      / \
//     3   6
//    / \
//   2   4
//  /
// 1
//输出: 4 
//
// 
//
// 限制： 
//
// 1 ≤ k ≤ 二叉搜索树元素个数 
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 202 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
   var dfs func(node *TreeNode)
   var res int
   tmp := k
   dfs = func(root *TreeNode) {
      if root == nil {
         return
      }
      dfs(root.Right)
      if tmp == 0 {
         return
      }
      tmp--
      if tmp == 0 {
        res = root.Val
      }
      dfs(root.Left)
   }
   dfs(root)
   return res
}
//leetcode submit region end(Prohibit modification and deletion)
