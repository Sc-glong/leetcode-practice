# [376.摆动序列](https://leetcode-cn.com/problems/wiggle-subsequence/solution/376bai-dong-xu-lie-tan-xin-si-wei-dui-bi-9or7/)
> https://leetcode-cn.com/problems/wiggle-subsequence/solution/376bai-dong-xu-lie-tan-xin-si-wei-dui-bi-9or7/
>
> 难度：中等

## 题目：

如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。

例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。

相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。 子序列
可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。

给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。

提示： -1 <= nums.length <= 1000

- 0 <= nums[i] <= 1000

## 示例：

```
示例 1：
输入：nums = [1,7,4,9,2,5]
输出：6
解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。

示例 2：
输入：nums = [1,17,5,10,13,15,10,5,16,8]
输出：7
解释：这个序列包含几个长度为 7 摆动序列。
其中一个是 [1, 17, 10, 13, 10, 16, 8] ，各元素之间的差值为 (16, -7, 3, -3, 6, -8) 。

示例 3：
输入：nums = [1,2,3,4,5,6,7,8,9]
输出：2

```

## 分析
首先根据题意，当数组为空或者只有一个数组的情况下，直接返回数组长度即可。

所谓摆动序列，可以抛开相等的情况，只考虑当前元素大于或者小于前一个元素的场景。
那么要比较，首先需要有一个常量记录上一次的比较结果。
当上一次的比较结果与本次不相符时，更新比较结果并计数+1，遍历数组重复判断即可。

## 解题：

```python
class Solution:
    def wiggleMaxLength(self, nums):
        ln = len(nums)
        if ln < 2:
            return ln
        diff = nums[1] - nums[0]
        ret = 2 if diff != 0 else 1
        for i in range(2, ln):
            tmp = nums[i] - nums[i - 1]
            if (tmp > 0 and diff <= 0) or (tmp < 0 and diff >= 0):
                ret += 1
                diff = tmp
        return ret
```

欢迎关注我的公众号: **清风Python**，带你每日学习Python算法刷题的同时，了解更多python小知识。

有喜欢力扣刷题的小伙伴可以加我微信（King_Uranus）互相鼓励，共同进步，一起玩转超级码力！

我的个人博客：[https://qingfengpython.cn](https://qingfengpython.cn)

力扣解题合集：[https://github.com/BreezePython/AlgorithmMarkdown](https://github.com/BreezePython/AlgorithmMarkdown)
