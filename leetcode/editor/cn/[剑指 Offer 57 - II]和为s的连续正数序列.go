//输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。 
//
// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。 
//
// 
//
// 示例 1： 
//
// 输入：target = 9
//输出：[[2,3,4],[4,5]]
// 
//
// 示例 2： 
//
// 输入：target = 15
//输出：[[1,2,3,4,5],[4,5,6],[7,8]]
// 
//
// 
//
// 限制： 
//
// 
// 1 <= target <= 10^5 
// 
//
// 
// Related Topics 数学 双指针 枚举 👍 322 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findContinuousSequence(target int) [][]int {
   var res [][]int
   i,j,sum := 1,2,3
   for i<=target/2{
      if sum< target {
         j++
         sum += j
      }else {
        if target == sum {
            tmp := make([]int,j-i+1)
            for k:=i;k<=j;k++{
               tmp[k-i] = k
            }
            res = append(res,tmp)
        }
        sum -=i
        i++
      }
   }
   return res
}
//leetcode submit region end(Prohibit modification and deletion)
