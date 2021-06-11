### 解题思路
列表的递归反转， 对撞指针解题变为 递归

### 代码

```python3
class Solution:
    def reverseString(self, s: List[str]) -> None:
        """
        Do not return anything, modify s in-place instead.
        """
        """ 列表的递归反转 """

        n = len(s)
        low = 0
        high = n - 1

        def reverse(low, high, s):
            if low < high:
                s[low], s[high] = s[high], s[low]
                reverse(low+1, high-1, s)

        reverse(low, high, s)
    
        return s

```