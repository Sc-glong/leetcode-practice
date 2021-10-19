//在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。 
//
// 示例 1: 
//
// 
//输入：s = "abaccdeff"
//输出：'b'
// 
//
// 示例 2: 
//
// 
//输入：s = "" 
//输出：' '
// 
//
// 
//
// 限制： 
//
// 0 <= s 的长度 <= 50000 
// Related Topics 队列 哈希表 字符串 计数 👍 140 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) byte {
//       pos := &[26]int{}
//       for _,v := range s {
//           pos[v-'a']++
//       }
//       for i,v := range s {
//          if pos[v-'a'] == 1 {
//             return s[i]
//          }
//       }
//       return ' '

         //方法2
         n := len(s)
         pos := [26]int{}
         for i := range s {
            pos[i] = n
         }
         for i,v := range s {
            if pos[v - 'a'] == n{
               pos[v - 'a'] = i
            }else{
               pos[v - 'a'] = n + 1
            }
         }
         ans := n
         for _,v := range pos[:]{
             if v < ans{
                ans = v
             }
         }
         if ans >= n {
           return ' '
         }
         return s[ans]
}
//leetcode submit region end(Prohibit modification and deletion)
