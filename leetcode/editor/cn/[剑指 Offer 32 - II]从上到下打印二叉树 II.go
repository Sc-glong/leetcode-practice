//从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。 
//
// 
//
// 例如: 
//给定二叉树: [3,9,20,null,null,15,7], 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
// 返回其层次遍历结果： 
//
// [
//  [3],
//  [9,20],
//  [15,7]
//]
// 
//
// 
//
// 提示： 
//
// 
// 节点总数 <= 1000 
// 
//
// 注意：本题与主站 102 题相同：https://leetcode-cn.com/problems/binary-tree-level-order-
//traversal/ 
// Related Topics 树 广度优先搜索 二叉树 👍 142 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    var bfs func(node *TreeNode, level int)
    bfs = func(node *TreeNode, level int) {
        if node == nil {
           return
        }
        if level >= len(res) {
           res = append(res,[]int{node.Val})
        } else {
           res[level] = append(res[level], node.Val)
        }
        level++
        bfs(node.Left,level)
        bfs(node.Right, level)
    }
    bfs(root,0)
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
