//数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1
//) 的解决方案。 
//
// 
//
// 示例 1： 
//
// 
//输入：[1,2,5,9,5,9,5,5,5]
//输出：5 
//
// 示例 2： 
//
// 
//输入：[3,2]
//输出：-1 
//
// 示例 3： 
//
// 
//输入：[2,2,1,1,1,2,2]
//输出：2 
// Related Topics 数组 计数 
// 👍 145 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
//
    countMap := make(map[int]int)
    for _,val := range nums {
         if _,ok := countMap[val];ok {
             countMap[val]++
         } else {
             countMap[val] = 1
         }
    }
    count := len(nums)
    max := 0
    index := 0
    for key,value := range countMap {
             if value > count/2 && value > max{
                max = value
                index = key
             }
    }
    if max > 0 {
        return index
    }
    return -1
}
//leetcode submit region end(Prohibit modification and deletion)
