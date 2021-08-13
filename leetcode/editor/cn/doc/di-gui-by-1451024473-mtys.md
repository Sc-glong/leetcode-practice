### 解题思路
递归思路

### 代码

```python3
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        ans =[]
        def order(root):
            if not root:return 
            order(root.left)
            order(root.right)
            ans.append(root.val)
        order(root)
        return ans 
```