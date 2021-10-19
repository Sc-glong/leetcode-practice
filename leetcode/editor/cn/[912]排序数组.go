//给你一个整数数组 nums，请你将该数组升序排列。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
// 
//
// 示例 2： 
//
// 输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 50000 
// -50000 <= nums[i] <= 50000 
// 
// Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 
// 👍 312 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
//    sort.Ints(nums)
//    return nums
//  冒泡排序
//      for i:= 0;i<len(nums);i++ {
//         for j:= 0;j<len(nums)-i-1;j++{
//             if nums[j] > nums[j+1] {
//                 tmp := nums[j+1]
//                 nums[j+1] = nums[j]
//                 nums[j] = tmp
//             }
//         }
//      }
//      return nums

 // 快排
//       quickSort(nums)
//       return nums

// 插入排序
//        for i:=1;i<len(nums);i++ {
//            for  j:=i;j>0&&nums[j]<nums[j-1];j-- {
//               tmp := nums[j-1]
//               nums[j-1] = nums[j]
//               nums[j] = tmp
//            }
//        }
//        return nums
}
func quickSort(nums []int){
       if len(nums) < 1 {
          return
       }
       left,right:= 0,len(nums) -1
       flag := nums[0]
       for i:=1;i<=right;{
           if nums[i] > flag {
              nums[i],nums[right] = nums[right] ,nums[i]
              right--
           } else {
              nums[i],nums[left] = nums[left], nums[i]
              i++
              left++
           }
       }
       quickSort(nums[:left])
       quickSort(nums[left+1:])
}

//leetcode submit region end(Prohibit modification and deletion)
