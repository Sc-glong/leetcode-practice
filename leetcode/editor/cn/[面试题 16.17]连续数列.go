//给定一个整数数组，找出总和最大的连续数列，并返回总和。 
//
// 示例： 
//
// 输入： [-2,1,-3,4,-1,2,1,-5,4]
//输出： 6
//解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 
//
// 进阶： 
//
// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。 
// Related Topics 数组 分治 动态规划 👍 92 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
    max := nums[0]
    for i:=1;i<len(nums);i++{
       if nums[i] + nums[i-1] > nums[i] {
          nums[i] += nums[i-1]
       }
       if max < nums[i] {
          max = nums[i]
       }
    }
    return max
}
//leetcode submit region end(Prohibit modification and deletion)
