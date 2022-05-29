//给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。 
//
// 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那
//么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,3,5,4,7]
//输出：3
//解释：最长连续递增序列是 [1,3,5], 长度为3。
//尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。 
// 
//
// 示例 2： 
//
// 
//输入：nums = [2,2,2,2,2]
//输出：1
//解释：最长连续递增序列是 [2], 长度为1。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁴ 
// -10⁹ <= nums[i] <= 10⁹ 
// 
// Related Topics 数组 👍 255 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findLengthOfLCIS(nums []int) int {
	//-------------------------------贪心算法-----------------------------------
	// Time: O(n)
	// Space: O(1)
	if len(nums) == 1 {
		return 1
	}
	cur, ans := 1 , 0
	for i := 0 ; i < len(nums) - 1 ; i++ {
		if nums[i] < nums[i + 1] {
			cur++
		}else{
			cur = 1
		}
		ans = max(ans, cur)
	}
	return ans
	//------------------------------------------------------------------------
}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
