### 解题思路
执行用时：0 ms, 在所有 Java 提交中击败了100.00%
的用户内存消耗：35.9 MB, 在所有 Java 提交中击败了100.00%的用户

### 代码

```java
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
class Solution {
    //设置虚拟节点
    TreeNode preNode = new TreeNode(-1);
    TreeNode currentNode = preNode;
    public TreeNode increasingBST(TreeNode root) {
        inOrder(root);
        return preNode.right;
    }

    //中序遍历，将每次遍历的元素添加到虚拟节点的right指针
    private void inOrder(TreeNode node){
        if (node == null){
            return;
        }
        inOrder(node.left);
        
        currentNode.right = node;
        currentNode = currentNode.right;
        currentNode.left = null;//左节点置空
        
        inOrder(node.right);
    }
}
```