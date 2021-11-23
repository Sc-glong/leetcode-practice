//给你一个字符串 s，找到 s 中最长的回文子串。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
// 
//
// 示例 2： 
//
// 
//输入：s = "cbbd"
//输出："bb"
// 
//
// 示例 3： 
//
// 
//输入：s = "a"
//输出："a"
// 
//
// 示例 4： 
//
// 
//输入：s = "ac"
//输出："a"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 1000 
// s 仅由数字和英文字母（大写和/或小写）组成 
// 
// Related Topics 字符串 动态规划 
// 👍 3896 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
  if len(s) < 2 {
     return s
  }
  begin, end := 0, 0
  m := make([][]int, len(s))
  for i:=0;i<len(s);i++ {
     m[i] = make([]int, len(s))
     m[i][i] = 1
  }
  for n:= 2;n<=len(s);n++ {
     for i:=0;i<len(s);i++ {
        j := n+i-1
        if j >=len(s) {
           break
        }
        if s[i] != s[j] {
            m[i][j] = 0
        } else {
              if j-i == 1 {
                 m[i][j] =1
              }else {
                 m[i][j] = m[i+1][j-1]
              }
        }
        if m[i][j] > 0 && j-i>end-begin {
             begin,end = i, j
        }
     }
  }
  return string(s[begin:end+1])
}
//leetcode submit region end(Prohibit modification and deletion)
