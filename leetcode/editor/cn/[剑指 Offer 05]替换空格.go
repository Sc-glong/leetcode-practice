//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。 
//
// 
//
// 示例 1： 
//
// 输入：s = "We are happy."
//输出："We%20are%20happy." 
//
// 
//
// 限制： 
//
// 0 <= s 的长度 <= 10000 
// Related Topics 字符串 👍 159 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
     var str string = ""
     for _,i := range s {
         if i == ' ' {
            str += "%20"
         }else {
         str += string(i)
         }
     }
     return str
}
//leetcode submit region end(Prohibit modification and deletion)
