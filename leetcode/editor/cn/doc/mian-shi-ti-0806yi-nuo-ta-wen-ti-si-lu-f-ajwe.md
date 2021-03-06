![汉诺塔.png](https://pic.leetcode-cn.com/1627285467-wriBGy-%E6%B1%89%E8%AF%BA%E5%A1%94.png)
### 解题思路
采用递归的思路
三要素如下：
递归结束条件：只剩下最后一个盘子需要移动
递归函数主功能：
1.首先将 n-1 个盘子，从第一个柱子移动到第二个柱子
2.然后将最后一个盘子移动到第三个柱子上
3.最后将第二个柱子上的 n-1 个盘子，移动到第三个柱子上
函数的等价关系式：
f(n,A,B,C) 表示将n个盘子从A移动到C
f(n,A,B,C)=f(n-1,A,C,B)+f(1,A,B,C)+f(n-1,B,A,C)
### 代码

```java
class Solution {
    public void hanota(List<Integer> A, List<Integer> B, List<Integer> C) {
        movePlant(A.size(),A,B,C);
    }
    //size 需要移动的盘子的数量
    //start 起始的柱子
    //auxiliary 辅助柱子
    //target 目标柱子
    public void movePlant(int size,List<Integer> start,List<Integer> auxiliary,List<Integer> target){
        //当只剩一个盘子时，直接将它从第一个柱子移动到第三个柱子
        if(size == 1){
            target.add(start.remove(start.size()-1));
            return;
        }
        //首先将 n-1 个盘子，从第一个柱子移动到第二个柱子
        movePlant(size - 1,start,target,auxiliary);
        //然后将最后一个盘子移动到第三个柱子上
        target.add(start.remove(start.size()-1));
        //最后将第二个柱子上的 n-1 个盘子，移动到第三个柱子上
        movePlant(size - 1,auxiliary,start,target);
       
    }
}
```
递归对于刚接触的朋友还是有一定的难度，大家多分析三要素和思路，多刷几道题，可能就会茅塞顿开，不必因为刚开始搞不明白就自我否定，畏难放弃呀！

冲冲冲,大家一起进步！