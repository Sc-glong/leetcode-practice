//给定一个二叉树，返回它的 后序 遍历。 
//
// 示例: 
//
// 输入: [1,null,2,3]  
//   1
//    \
//     2
//    /
//   3 
//
//输出: [3,2,1] 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 深度优先搜索 二叉树 
// 👍 643 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    tmp := []int{}
    traversal(root,&tmp)
    return tmp
}
func traversal(root *TreeNode,tmp *[]int) {
  if root == nil {
     return
  }
  traversal(root.Left,tmp)
  traversal(root.Right,tmp)
  *tmp = append(*tmp,root.Val)
}
//leetcode submit region end(Prohibit modification and deletion)
