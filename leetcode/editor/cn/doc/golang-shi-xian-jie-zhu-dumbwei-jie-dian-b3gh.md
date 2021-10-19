```
func increasingBST(root *TreeNode) *TreeNode {
    dumb := TreeNode{}
    pre := &dumb
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode){
        if node == nil {
            return
        }
        dfs(node.Left)
        pre.Right = node
        node.Left = nil
        pre = node
        dfs(node.Right)
    }
    dfs(root)
    return dumb.Right
}
```
