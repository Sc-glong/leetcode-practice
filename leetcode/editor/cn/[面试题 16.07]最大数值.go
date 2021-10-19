//编写一个方法，找出两个数字a和b中最大的那一个。不得使用if-else或其他比较运算符。 
// 示例： 
// 输入： a = 1, b = 2
//输出： 2
// 
// Related Topics 位运算 脑筋急转弯 数学 👍 91 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maximum(a int, b int) int {
     var diff float64
     diff = float64(a-b)
     return int((int(math.Abs(diff)) + a +b)/2)
}
//leetcode submit region end(Prohibit modification and deletion)
