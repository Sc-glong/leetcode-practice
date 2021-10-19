//从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，
//可以看成任意数字。A 不能视为 14。 
//
// 
//
// 示例 1: 
//
// 
//输入: [1,2,3,4,5]
//输出: True 
//
// 
//
// 示例 2: 
//
// 
//输入: [0,0,1,2,5]
//输出: True 
//
// 
//
// 限制： 
//
// 数组长度为 5 
//
// 数组的数取值为 [0, 13] . 
// Related Topics 数组 排序 👍 161 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isStraight(nums []int) bool {
    length := len(nums)
    // 选择
//     for i:=0;i<length;i++{
//       for j:=i+1;j<length;j++{
//          if nums[j] < nums[i] {
//             temp := nums[i]
//             nums[i] = nums[j]
//             nums[j] = temp
//          }
//       }
//     }

      //冒泡
//     for i:=0;i<length;i++{
//       for j:=0;j<length-i-1;j++{
//          if nums[j] > nums[j+1] {
//              temp := nums[j]
//              nums[j] = nums[j+1]
//              nums[j+1] = temp
//          }
//       }
//     }

     // 插入
//     for i:=0;i<length;i++{
//       for j:=i;j>0&& nums[j] < nums[j-1];j--{
//            temp := nums[j]
//            nums[j] = nums[j+1]
//            nums[j+1] = temp
//       }
//     }

    joker := 0
    for i:=0;i<length-1;i++{
       if nums[i] == 0 {
          joker++
       }else if nums[i] == nums[i+1] {
          return false
       }
    }
    return nums[4] - nums[joker] < 5
}
func quickSort(nums []int) {
   length := len(nums)
   if length <=1 {
      return
   }
   flag := nums[0]
   low := 0
   high := length-1
   for i:=1;i<length;i++{
      if nums[i] > flag {
         nums[i],nums[high] = nums[high],nums[i]
         high --
      } else {
         nums[i], nums[low] = nums[low],nums[i]
         i++
         low++
      }
   }
   quickSort(nums[:low])
   quickSort(nums[low+1:])
}
//leetcode submit region end(Prohibit modification and deletion)
