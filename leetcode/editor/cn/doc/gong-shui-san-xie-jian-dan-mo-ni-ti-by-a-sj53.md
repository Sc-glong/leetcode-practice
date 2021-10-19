## 模拟

按照题意模拟一遍，得到数列 *nums*，再从 *nums* 中找出最大值即可。

代码：
```Java []
class Solution {
    public int getMaximumGenerated(int n) {
        if (n == 0) return 0;
        int[] nums = new int[n + 1];
        nums[0] = 0;
        nums[1] = 1;
        for (int i = 0; i < n; i++) {
            if (2 * i <= n) nums[2 * i] = nums[i];
            if (2 * i + 1 <= n) nums[2 * i + 1] = nums[i] + nums[i + 1];
        }
        int ans = 0;
        for (int i : nums) ans = Math.max(ans, i);
        return ans;
    }
}
```
* 时间复杂度：*O(n)*
* 空间复杂度：*O(n)*

---

## 打表 

利用数据范围，可以直接使用 `static` 进行打表构造。

代码：
```Java []
class Solution {
    static int N = 110;
    static int[] nums = new int[N];
    static {
        nums[0] = 0;
        nums[1] = 1;
        for (int i = 0; i < N; i++) {
            if (2 * i < N) nums[2 * i] = nums[i];
            if (2 * i + 1 < N) nums[2 * i + 1] = nums[i] + nums[i + 1];
        }
        for (int i = 0, max = 0; i < N; i++) {
            nums[i] = max = Math.max(max, nums[i]);
        }
    }
    public int getMaximumGenerated(int n) {
        return nums[n];       
    }
}
```
* 时间复杂度：将打表逻辑放到本地进行，复杂度为 *O(1)*
* 空间复杂度：*O(n)*