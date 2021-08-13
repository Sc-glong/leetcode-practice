# [260. 只出现一次的数字 III](https://leetcode-cn.com/problems/single-number-iii/)
***
### 思路一：哈希表
>最容易想到的思路，只是不符合题目空间复杂度O(1)的进阶要求，想要O(1)还是要看位运算
1. 用哈希表记录每个数出现的次数
2. 遍历哈希表，将次数为1的加入到结果中

### 代码
```Python3 []
class Solution:
    def singleNumber(self, nums: List[int]) -> List[int]:
        cnt = collections.Counter(nums)
        res = []
        for key, val in cnt.items():
            if val == 1:
                res.append(key)
        return res
```

```Java []
class Solution {
    public int[] singleNumber(int[] nums) {
        HashMap<Integer, Integer> hashmap = new HashMap<>();
        for(int num : nums){
            hashmap.put(num, hashmap.getOrDefault(num, 0) + 1);
        }
        int[] res = new int[2];
        int idx = 0;
        for(Map.Entry<Integer, Integer> entry : hashmap.entrySet()){
            if(entry.getValue() == 1){
                res[idx++] = entry.getKey();
            }
        }
        return res;
    }
}
```

**复杂度分析**
- 时间复杂度：*O(n)*
- 空间复杂度：*O(n)*
***
### 思路二：位运算
1. 因为a^a = 0，先遍历数组对所有数异或，得到两个只出现一次的数的异或结果eor
2. 通过eor & - eor取得eor的最右边的1，令其为rightOne。因为两个数对应位不同异或后该位才会为1，所以rightOne对应位在两个数中一定是一个为0，一个为1，这是在为后面分组做准备
3. 再遍历数组，每个数和rightOne与，将所有数分成与运算后为0和为1两组，要找的两个数分别在两组
4. 把其中一组的所有数异或，得到其中一个要找的数res，另一个数即为res ^ eor

### 代码
```Python3 []
class Solution:
    def singleNumber(self, nums: List[int]) -> List[int]:
        eor = 0
        for num in nums:
            eor ^= num
        rightOne = eor & - eor
        res = 0
        for num in nums:
            if num & rightOne != 0:
                res ^= num
        return [res, res ^ eor]
```

```Java []
class Solution {
    public int[] singleNumber(int[] nums) {
        int eor = 0;
        for(int num : nums){
            eor ^= num;
        }
        int rightOne = eor & (- eor);
        int res = 0;
        for(int num : nums){
            if((num & rightOne) != 0){
                res ^= num;
            }
        }
        return new int[]{res, res ^ eor};
    }
}
```

**复杂度分析**
- 时间复杂度：*O(n)*
- 空间复杂度：*O(1)*
***
### 数字出现次数相关问题（位运算）
|  题目   | 我的题解  |
|  ----  | ----  |
| [136. 只出现一次的数字](https://leetcode-cn.com/problems/single-number/)  | [我的题解](https://leetcode-cn.com/problems/single-number/solution/136-zhi-chu-xian-yi-ci-de-shu-zi-wei-yun-igeb/) |
| [137. 只出现一次的数字 II](https://leetcode-cn.com/problems/single-number-ii/)  | [我的题解](https://leetcode-cn.com/problems/single-number-ii/solution/137-zhi-chu-xian-yi-ci-de-shu-zi-iiwei-y-oa1w/)|
| [645. 错误的集合](https://leetcode-cn.com/problems/set-mismatch/)  | [我的题解](https://leetcode-cn.com/problems/set-mismatch/solution/645-cuo-wu-de-ji-he-wei-yun-suan-fen-zu-k5i5n/) |
| [260. 只出现一次的数字 III](https://leetcode-cn.com/problems/single-number-iii/)  | [我的题解](https://leetcode-cn.com/problems/single-number-iii/solution/260-zhi-chu-xian-yi-ci-de-shu-zi-iiiha-x-l37m/) |