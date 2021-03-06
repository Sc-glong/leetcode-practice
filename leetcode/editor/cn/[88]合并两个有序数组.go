//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。 
//
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nu
//ms2 的元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
// 
//
// 示例 2： 
//
// 
//输入：nums1 = [1], m = 1, nums2 = [], n = 0
//输出：[1]
// 
//
// 
//
// 提示： 
//
// 
// nums1.length == m + n 
// nums2.length == n 
// 0 <= m, n <= 200 
// 1 <= m + n <= 200 
// -109 <= nums1[i], nums2[i] <= 109 
// 
// Related Topics 数组 双指针 
// 👍 972 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int)  {
//    copy(nums1[m:],nums2)
//    sort.Ints(nums1)

      //双指针
//       sorted := make([]int,0,m+n)
//       p1,p2 := 0,0
//       for {
//          if p1 >= m {
//             sorted = append(sorted,nums2[p2:]...)
//             break
//          }
//          if p2>=n {
//             sorted = append(sorted,nums1[p1:]...)
//             break
//          }
//          if nums1[p1] > nums2[p2] {
//               sorted = append(sorted,nums2[p2])
//               p2++
//          } else {
//               sorted = append(sorted,nums1[p1])
//               p1++
//          }
//       }
//       copy(nums1,sorted)

       //逆向双指针
       for p1,p2,tail := m-1,n-1,m+n-1;p1>=0 || p2>= 0;tail-- {
             var cur int
             if p1<0 {
               cur = nums2[p2]
               p2--
             }else if p2<0 {
               cur = nums1[p1]
               p1--
             }else if nums1[p1] > nums2[p2] {
                cur = nums1[p1]
                p1--
             } else {
                cur = nums2[p2]
                p2--
             }
             nums1[tail] = cur
       }
}
//leetcode submit region end(Prohibit modification and deletion)
