//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。 
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
// 返回： 
//
// [3,9,20,15,7]
// 
//
// 
//
// 提示： 
//
// 
// 节点总数 <= 1000 
// 
// Related Topics 树 广度优先搜索 二叉树 👍 129 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
     // bfs
     res := []int{}
     if root == nil {
         return res
     }
     queue := make([]*TreeNode,0)
     queue = append(queue,root)
     for len(queue) > 0 {
         node := queue[0]
         res = append(res,node.Val)
         queue = queue[1:]
         if node.Left != nil {
             queue = append(queue,node.Left)
         }
         if node.Right != nil {
             queue = append(queue,node.Right)
         }
     }
     return res
}
//leetcode submit region end(Prohibit modification and deletion)
