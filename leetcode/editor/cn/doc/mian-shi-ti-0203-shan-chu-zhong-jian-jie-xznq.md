### 解题思路
将待删节点的下一节点的值复制入待删节点中，再删掉待删节点的下一节点。
此时的结果等价于将待删节点删除。

### 代码

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public void deleteNode(ListNode node) {
        //杀不掉我，我就变成你，然后再干掉你，就等于杀死了自己。
        node.val = node.next.val;
        node.next = node.next.next; 
    }
}



```