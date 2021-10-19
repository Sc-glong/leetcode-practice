思路和心得：


# （一）dfs_LNR


```python3 []
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def increasingBST(self, root: TreeNode) -> TreeNode:
        dummy = TreeNode(-1)
        mid = dummy     #桥接的点

        def dfs_LNR(x: TreeNode) -> TreeNode:
            nonlocal mid
            if x == None:
                return 
            
            dfs_LNR(x.left)
            
            mid.right = x
            x.left = None
            mid = x

            dfs_LNR(x.right)
        
        dfs_LNR(root)
        return dummy.right
```

```c++ []
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution 
{
public:
    TreeNode * dummy;
    TreeNode * mid;

    TreeNode* increasingBST(TreeNode* root) 
    {
        dummy = new TreeNode(-1);
        mid = dummy;
        dfs_LNR(root);
        return dummy->right;
    }

    void dfs_LNR(TreeNode * x)
    {
        if (x == nullptr)
            return ;
        
        dfs_LNR(x->left);

        mid->right = x;
        x->left = NULL;
        mid = x;
        
        dfs_LNR(x->right);
    }
};
```

```java []
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution 
{
    TreeNode dummy = new TreeNode(0);
    TreeNode mid = dummy;

    public TreeNode increasingBST(TreeNode root) 
    {   
        dfs_LNR(root);
        return dummy.right;
    }

    public void dfs_LNR(TreeNode x)
    {
        if (x == null)
            return ;
        
        dfs_LNR(x.left);

        mid.right = x;
        x.left = null;
        mid = x;

        dfs_LNR(x.right);
    }
}
```