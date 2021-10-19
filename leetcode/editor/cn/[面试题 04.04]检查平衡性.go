//实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。 示例 1: 给定二叉树 [3,9,20,
//null,null,15,7]     3    / \   9  20     /  \    15   7 返回 true 。示例 2: 给定二叉树 [1,2,
//2,3,3,null,null,4,4]       1      / \     2   2    / \   3   3  / \ 4   4 返回 
//false 。 Related Topics 树 深度优先搜索 二叉树 👍 69 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
//      if root == nil {
//         return true
//      }
//      return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
     return height2(root) >= 0
}
func height2(root *TreeNode) int {
        if root == nil {
           return 0
        }
        leftHeight := height2(root.Left)
        rightHeight := height2(root.Right)
        if leftHeight == -1 || rightHeight == -1 || abs(leftHeight - rightHeight) > 1 {
          return -1
        }
        return max(leftHeight,rightHeight) +1
}

func height(root *TreeNode) int {
   if root == nil {
      return 0
   }
   return max(height(root.Left),height(root.Right))+1
}

func abs (x int) int {
   if x <0 {
      x = -1 * x
   }
   return x
}
func max (x, y int) int {
   if x > y {
      return x
   }
   return y
}
//leetcode submit region end(Prohibit modification and deletion)
