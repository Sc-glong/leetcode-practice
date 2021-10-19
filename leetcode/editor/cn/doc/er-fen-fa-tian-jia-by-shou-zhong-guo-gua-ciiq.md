### 解题思路
把问题简化成:
把一个数字加到一个有序数组里,得到一个新的有序数组.
再简化成:
找到这个数应该加到哪个位置
最后简化成一个二分搜索问题...
### 代码

```golang
type KthLargest struct {
	k    int
	list []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{
		k:    k,
		list: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	this.list = binaryAdd(this.list, val)
	l := len(this.list)
	return this.list[l-this.k]
}
func binaryAdd(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return append(nums[:l], append([]int{target}, nums[l:]...)...)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
```