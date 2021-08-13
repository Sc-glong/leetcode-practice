
### 代码

```golang

func missingNumber(nums []int) int {
	n := len(nums)
	ans := n
	for i := 0; i < n; i++ {
		ans ^= i
		ans ^= nums[i]
	}
	return ans
}

```