//给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。 
//
// 示例 1： 
//
// 输入: num = 1775(110111011112)
//输出: 8
// 
//
// 示例 2： 
//
// 输入: num = 7(01112)
//输出: 4
// 
// Related Topics 位运算 动态规划 👍 49 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func reverseBits(num int) int {
    cur := 0
    insert := 0
    res := 1
    for i:=0;i<32;i++{
        if (num & (1<<i)) != 0  {
           cur++
           insert++
        } else {
           insert = cur +1
           cur =0
        }
        res = max(res,insert)
    }
    return res
}
func max(x,y int) int{
   if x>y{
     return x
   }
   return y
}
//leetcode submit region end(Prohibit modification and deletion)
