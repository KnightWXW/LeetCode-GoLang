//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。 
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
//序列。 
// 
//
// 示例 1： 
//
// 
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1,0,3,2,3]
//输出：4
// 
//
// 示例 3： 
//
// 
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 2500 
// -10⁴ <= nums[i] <= 10⁴ 
// 
//
// 
//
// 进阶： 
//
// 
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗? 
// 
// Related Topics 数组 二分查找 动态规划 👍 2467 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	//------------------------------递归(超时)-------------------------------
	/*// Time: O(n ^ 2)
	// Space: O(n)
	var dfslengthOfLIS func(nums []int, index int) int
	dfslengthOfLIS = func(nums []int, index int) int {
		if index < 0 {
			return 0
		}
		if index == 0 {
			return 1
		}

		cur := 0
		for i := 0 ; i < index ; i++ {
			if nums[index] > nums[i] {
				cur = max(dfslengthOfLIS(nums, i), cur)
			}
		}
		return cur + 1
	}

	ans := 1
	for i := 0 ; i < len(nums) ; i++ {
		ans = max(ans, dfslengthOfLIS(nums, i))
	}
	return ans*/
	//--------------------------------------------------------------------

	//------------------------------记忆化搜索-------------------------------
	/*// Time: O(n ^ 2)
	// Space: O(n)
	n := len(nums)
	dp := make([]int, n)
	var dfslengthOfLIS func(nums []int, index int) int
	dfslengthOfLIS = func(nums []int, index int) int {
		if index < 0 {
			return 0
		}
		if index == 0 {
			return 1
		}
		if dp[index] != 0 {
			return dp[index]
		}

		cur := 0
		for i := 0 ; i < index ; i++ {
			if nums[index] > nums[i] {
				cur = max(dfslengthOfLIS(nums, i), cur)
			}
		}
		dp[index] = cur + 1
		return dp[index]
	}

	ans := 1
	for i := 0 ; i < len(nums) ; i++ {
		ans = max(ans, dfslengthOfLIS(nums, i))
	}
	return ans*/
	//--------------------------------------------------------------------

	//----------------------------贪心 + 动态规划----------------------------
	// Time: O(n ^ 2)
	// Space: O(n)
	n := len(nums)
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	ans := 0
	for i := 1 ; i < n ; i++ {
		dp[i] = 1
		for j := 0 ; j < i ; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
			ans = max(ans, dp[i])
		}
	}
	return ans
	//-------------------------------------------------------------------
}

func max(a , b int) int {
	if a < b {
		return b
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
