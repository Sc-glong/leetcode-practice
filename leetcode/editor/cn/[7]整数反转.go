//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。 
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−231, 231 − 1] ，就返回 0。 
//假设环境不允许存储 64 位整数（有符号或无符号）。
//
// 
//
// 示例 1： 
//
// 
//输入：x = 123
//输出：321
// 
//
// 示例 2： 
//
// 
//输入：x = -123
//输出：-321
// 
//
// 示例 3： 
//
// 
//输入：x = 120
//输出：21
// 
//
// 示例 4： 
//
// 
//输入：x = 0
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// -231 <= x <= 231 - 1 
// 
// Related Topics 数学 
// 👍 3007 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
    flag := 0
    if x < 0 {
       flag = 1
       x = 0-x
    }
    tmp,num := 0, 0
    for x > 0 {
       tmp = x % 10
       num = num * 10 + tmp
       x = x/ 10
    }
    if  flag == 1 {
       num  = 0 -num
    }
    if num > math.MaxInt32 || num < math.MinInt32 {
       return 0
    }
    return num

}
//leetcode submit region end(Prohibit modification and deletion)
