//给你一棵二叉搜索树，请你返回一棵 平衡后 的二叉搜索树，新生成的树应该与原来的树有着相同的节点值。 
//
// 如果一棵二叉搜索树中，每个节点的两棵子树高度差不超过 1 ，我们就称这棵二叉搜索树是 平衡的 。 
//
// 如果有多种构造方法，请你返回任意一种。 
//
// 
//
// 示例： 
//
// 
//
// 输入：root = [1,null,2,null,3,null,4,null,null]
//输出：[2,1,3,null,null,null,4]
//解释：这不是唯一的正确答案，[3,1,4,null,2,null,null] 也是一个可行的构造方案。
// 
//
// 
//
// 提示： 
//
// 
// 树节点的数目在 1 到 10^4 之间。 
// 树节点的值互不相同，且在 1 到 10^5 之间。 
// 
// Related Topics 贪心 树 深度优先搜索 二叉搜索树 分治 二叉树 
// 👍 70 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    //先中序遍历，再利用贪心
    if root == nil {
           return root
    }
    tmp := []int{}
    midOrder(root,&tmp)
    return buildTree(tmp,0,len(tmp)-1)
}
func midOrder(root *TreeNode,tmp *[]int) {
   if root ==nil {
      return
   }
   midOrder(root.Left,tmp)
   *tmp = append(*tmp,root.Val)
   midOrder(root.Right,tmp)
}

func buildTree(tmp []int,left,right int) *TreeNode {
    if left > right {
      return nil
    }
    mid := (left + right) / 2
    root := &TreeNode{Val: tmp[mid]}
    root.Left = buildTree(tmp,left, mid-1)
    root.Right = buildTree(tmp,mid+1 ,right)
    return root
}
//leetcode submit region end(Prohibit modification and deletion)
