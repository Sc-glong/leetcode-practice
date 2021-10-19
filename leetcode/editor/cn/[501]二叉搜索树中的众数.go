//给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。 
//
// 假定 BST 有如下定义： 
//
// 
// 结点左子树中所含结点的值小于等于当前结点的值 
// 结点右子树中所含结点的值大于等于当前结点的值 
// 左子树和右子树都是二叉搜索树 
// 
//
// 例如： 
//给定 BST [1,null,2,2], 
//
//    1
//    \
//     2
//    /
//   2
// 
//
// 返回[2]. 
//
// 提示：如果众数超过1个，不需考虑输出顺序 
//
// 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内） 
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 343 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    var base, count, maxCount int
    var answer []int
    update := func(x int) {
        if x == base {
           count++
        } else {
           base, count = x, 1
        }
        if count == maxCount {
           answer = append(answer, base)
        } else if count > maxCount {
              maxCount = count
              answer = []int{base}
        }
    }
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
          return
        }
        dfs(node.Left)
        update(node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return answer
}
//leetcode submit region end(Prohibit modification and deletion)
