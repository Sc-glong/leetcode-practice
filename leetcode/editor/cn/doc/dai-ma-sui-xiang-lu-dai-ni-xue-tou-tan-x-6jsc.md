## 思路 

**不少同学对贪心算法还处于朦胧状态，我特意录了一期视频，讲一讲[贪心算法理论基础](https://www.bilibili.com/video/BV1WK4y1R71x)**，这里详细介绍了我们做贪心算法的时候常遇到的问题，相信结合本篇题解，会对你学习贪心有所帮助。


为了了满足更多的小孩，就不要造成饼干尺寸的浪费。

大尺寸的饼干既可以满足胃口大的孩子也可以满足胃口小的孩子，那么就应该优先满足胃口大的。

**这里的局部最优就是大饼干喂给胃口大的，充分利用饼干尺寸喂饱一个，全局最优就是喂饱尽可能多的小孩**。

可以尝试使用贪心策略，先将饼干数组和小孩数组排序。

然后从后向前遍历小孩数组，用大饼干优先满足胃口大的，并统计满足小孩数量。

如图：

![455.分发饼干](https://pic.leetcode-cn.com/1625388787-TQCyYh-file_1625388787553)

这个例子可以看出饼干9只有喂给胃口为7的小孩，这样才是整体最优解，并想不出反例，那么就可以撸代码了。


C++代码整体如下：

```C++
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
class Solution {
public:
    int findContentChildren(vector<int>& g, vector<int>& s) {
        sort(g.begin(), g.end());
        sort(s.begin(), s.end());
        int index = s.size() - 1; // 饼干数组的下表
        int result = 0;
        for (int i = g.size() - 1; i >= 0; i--) {
            if (index >= 0 && s[index] >= g[i]) {
                result++;
                index--;
            }
        }
        return result;
    }
};
```

从代码中可以看出我用了一个index来控制饼干数组的遍历，遍历饼干并没有再起一个for循环，而是采用自减的方式，这也是常用的技巧。

有的同学看到要遍历两个数组，就想到用两个for循环，那样逻辑其实就复杂了。

**也可以换一个思路，小饼干先喂饱小胃口**

代码如下：

```C++
class Solution {
public:
    int findContentChildren(vector<int>& g, vector<int>& s) {
        sort(g.begin(),g.end());
        sort(s.begin(),s.end());
        int index = 0;
        for(int i = 0;i < s.size();++i){
            if(index < g.size() && g[index] <= s[i]){
                index++;
            }
        }
        return index;
    }
};
```

## 总结

这道题是贪心很好的一道入门题目，思路还是比较容易想到的。

文中详细介绍了思考的过程，**想清楚局部最优，想清楚全局最优，感觉局部最优是可以推出全局最优，并想不出反例，那么就试一试贪心**。

## 其他语言版本


Java：
```java
class Solution {
    public int findContentChildren(int[] g, int[] s) {
        Arrays.sort(g);
        Arrays.sort(s);
        int start = 0;
        int count = 0;
        for (int i = 0; i < s.length && start < g.length; i++) {
            if (s[i] >= g[start]) {
                start++;
                count++;
            }
        }
        return count;
    }
}
```

Python：
```python3
class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        g.sort()
        s.sort()
        res = 0
        for i in range(len(s)):
            if res <len(g) and s[i] >= g[res]:  #小饼干先喂饱小胃口
                res += 1
        return res
```
Go：

Javascript:
```Javascript

var findContentChildren = function(g, s) {
    g = g.sort((a, b) => a - b)
    s = s.sort((a, b) => a - b)
    let result = 0
    let index = s.length - 1
    for(let i = g.length - 1; i >= 0; i--) {
        if(index >= 0 && s[index] >= g[i]) {
            result++
            index--
        }
    } 
    return result
};

```

# 贪心算法力扣题目总结

按照如下顺序刷力扣上的题目，相信会帮你在学习贪心算法的路上少走很多弯路。

1. [贪心算法理论基础](https://mp.weixin.qq.com/s/O935TaoHE9Eexwe_vSbRAg)
2. [455.分发饼干](https://mp.weixin.qq.com/s/YSuLIAYyRGlyxbp9BNC1uw)
3. [376.摆动序列](https://mp.weixin.qq.com/s/Xytl05kX8LZZ1iWWqjMoHA)
4. [53.最大子序和](https://mp.weixin.qq.com/s/DrjIQy6ouKbpletQr0g1Fg)
5. [122买卖股票的最佳时机II](https://mp.weixin.qq.com/s/VsTFA6U96l18Wntjcg3fcg)
6. [55.跳跃游戏](https://mp.weixin.qq.com/s/606_N9j8ACKCODoCbV1lSA)
7. [45.跳跃游戏II](https://mp.weixin.qq.com/s/kJBcsJ46DKCSjT19pxrNYg)
8. [1005.K次取反后最大化的数组和](https://mp.weixin.qq.com/s/dMTzBBVllRm_Z0aaWvYazA)
9. [134.加油站](https://mp.weixin.qq.com/s/aDbiNuEZIhy6YKgQXvKELw)
10. [135.分发糖果](https://mp.weixin.qq.com/s/8MwlgFfvaNYmjGwjuMlETQ)
11. [860柠檬水找零](https://mp.weixin.qq.com/s/0kT4P-hzY7H6Ae0kjQqnZg)
12. [406.根据身高重建队列](https://mp.weixin.qq.com/s/-2TgZVdOwS-DvtbjjDEbfw)
13. [406.根据身高重建队列（续集）](https://mp.weixin.qq.com/s/K-pRN0lzR-iZhoi-1FgbSQ)
14. [452.用最少数量的箭引爆气球](https://mp.weixin.qq.com/s/HxVAJ6INMfNKiGwI88-RFw)
15. [435.无重叠区间](https://mp.weixin.qq.com/s/oFOEoW-13Bm4mik-aqAOmw)
16. [763.划分字母区间](https://mp.weixin.qq.com/s/pdX4JwV1AOpc_m90EcO2Hw)
17. [56.合并区间](https://mp.weixin.qq.com/s/royhzEM5tOkUFwUGrNStpw)
18. [738.单调递增的数字](https://mp.weixin.qq.com/s/TAKO9qPYiv6KdMlqNq_ncg)
19. [714.买卖股票的最佳时机含手续费](https://mp.weixin.qq.com/s/olWrUuDEYw2Jx5rMeG7XAg)
20. [968.我要监控二叉树！](https://mp.weixin.qq.com/s/kCxlLLjWKaE6nifHC3UL2Q)
21. [贪心算法总结篇！](https://mp.weixin.qq.com/s/ItyoYNr0moGEYeRtcjZL3Q)


------------

**大家好，我是程序员Carl，点击[我的头像](https://mp.weixin.qq.com/s/_DzddsMeQW5JPI6qoC7ARQ)**，查看力扣详细刷题攻略，你会发现相见恨晚！

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**


