//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：nums = [0]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 3000 
// -10⁵ <= nums[i] <= 10⁵ 
// 
// Related Topics 数组 双指针 排序 👍 3931 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
      n := len(nums)
      sort.Ints(nums)
      ans := make([][]int,0)
      for first:=0;first<n;first++ {
            if first > 0 && nums[first] == nums[first-1] {
               continue
            }
            third := n-1
            target := -1 * nums[first]
            for second:= first+1;second<n;second++ {
                   if second > first +1 && nums[second] == nums[second-1] {
                        continue
                   }
                   for second < third && nums[second] + nums[third] >  target {
                      third--
                   }
                   if second == third {
                        break
                   }
                   if nums[second] + nums[third] == target {
                      ans = append(ans,[]int{nums[first], nums[second], nums[third]})
                   }
            }
      }
      return ans
}
//leetcode submit region end(Prohibit modification and deletion)
