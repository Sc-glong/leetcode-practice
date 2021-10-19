//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。 
//
// 
//
// 示例 1： 
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
// 
//
// 示例 2： 
//
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
// 
//
// 
//
// 限制： 
//
// 
// 0 <= matrix.length <= 100 
// 0 <= matrix[i].length <= 100 
// 
//
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/ 
// Related Topics 数组 矩阵 模拟 👍 297 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
   if len(matrix) == 0 || len(matrix[0]) == 0 {
       return []int{}
   }
   var (
      rows, columns = len(matrix), len(matrix[0])
      order = make([]int, columns * rows)
      index = 0
      left, right, top, bottom = 0, columns-1, 0, rows-1
   )
   for left <= right && top <= bottom {
        for column := left;column <= right;column++{
             order[index] = matrix[top][column]
             index++
        }
        for tmp:=top+1;tmp<=bottom;tmp++{
            order[index] = matrix[tmp][right]
            index++
        }
        if left < right && top < bottom {
           for column:= right-1;column>left;column--{
               order[index] = matrix[bottom][column]
               index++
           }
           for tmp:= bottom;tmp>top;tmp--{
               order[index] = matrix[tmp][left]
               index++
           }
        }
        left++
        right--
        top++
        bottom--
   }
   return order
}
//leetcode submit region end(Prohibit modification and deletion)
