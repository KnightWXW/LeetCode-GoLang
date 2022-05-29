//给你一个整数数组 nums 和一个整数 target 。 
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ： 
//
// 
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。 
// 
//
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3
// 
//
// 示例 2： 
//
// 
//输入：nums = [1], target = 1
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 20 
// 0 <= nums[i] <= 1000 
// 0 <= sum(nums[i]) <= 1000 
// -1000 <= target <= 1000 
// 
// Related Topics 数组 动态规划 回溯 👍 1167 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, target int) int {
	//--------------------------------递归----------------------------------------
	/*// Time: O(2^n)
	// Space: O(n)
	var dfs func(nums []int, index int, sum int) int
	dfs = func(nums []int, index int, sum int) int {
		if index == len(nums) {
			if target == sum {
				return 1
			}else{
				return 0
			}
		}
		return dfs(nums, index + 1, sum + nums[index]) + dfs(nums, index + 1, sum - nums[index])
	}

	return dfs(nums, 0, 0)*/
	//---------------------------------------------------------------------------

	//----------------------------递归(记忆化搜索)----------------------------------
	/*// Time: O(2^n)
	// Space: O(n)
	hashMap := map[string]int{}
	var dfs func(nums []int, index int, sum int) int
	dfs = func(nums []int, index int, sum int) int {
		if index == len(nums) {
			if target == sum {
				return 1
			}else{
				return 0
			}
		}
		tem := string(sum) + "&" + string(index)
		k, ok := hashMap[tem]
		if ok == true {
			return k
		}
		ans := dfs(nums, index + 1, sum + nums[index]) + dfs(nums, index + 1, sum - nums[index])
		hashMap[tem] = ans
		return ans
	}

	return dfs(nums, 0, 0)*/
	//---------------------------------------------------------------------------

	//--------------------------------动态规划-------------------------------------
	/*// Time: O(n ∗ 2 * sum + 1)
	// Space: O(n ∗ 2 * sum + 1)
	sum := 0
	for i := 0 ; i < len(nums) ; i++ {
		sum += nums[i]
	}
	if abs(target) > abs(sum) {
		return 0
	}
	n := len(nums)
	dp := make([][]int, n + 1)
	for i := range dp {
		dp[i] = make([]int, 2 * sum + 1)
	}
	dp[0][sum] = 1
	for i := 1 ; i <= n ; i++ {
		tem := nums[i - 1]
		for j := -sum ; j <= sum ; j++ {
			if j - tem + sum >= 0 {
				dp[i][j + sum] += dp[i - 1][j - tem + sum]
			}
			if j + tem + sum <= 2 * sum {
				dp[i][j + sum] += dp[i - 1][j + tem + sum]
			}
		}
	}
	return dp[n][target + sum]*/
	//----------------------------------------------------------------------------

	//--------------------------------动态规划(背包问题)-------------------------------------
	/*// Time: O(n ∗ k)
	// Space: O(n ∗ k)
	sum := 0
	for i := 0 ; i < len(nums) ; i++ {
		sum += nums[i]
	}
	if abs(target) > abs(sum) {
		return 0
	}
	dif := sum - target
	if dif < 0 || dif & 1 == 1 {
		return 0
	}
	n , k := len(nums), dif / 2
	dp := make([][]int, n + 1)
	for i := 0 ; i < len(dp) ; i++ {
		dp[i] = make([]int, k + 1)
	}
	dp[0][0] = 1
	for i := 1 ; i <= n ; i++ {
		for j := 0 ; j <= k ; j++ {
			if j >= nums[i - 1] {
				dp[i][j] = dp[i - 1][j] + dp[i - 1][j - nums[i - 1]]
			}else{
				dp[i][j] = dp[i - 1][j]
			}
		}
	}
	return dp[n][k]*/
	//----------------------------------------------------------------------------

	//------------------------动态规划(背包问题)(空间优化)------------------------------
	// Time: O(k)
	// Space: O(k)
	sum := 0
	for i := 0 ; i < len(nums) ; i++ {
		sum += nums[i]
	}
	if abs(target) > abs(sum) {
		return 0
	}
	dif := sum - target
	if dif < 0 || dif & 1 == 1 {
		return 0
	}

	n , k := len(nums), dif / 2
	dp := make([]int, k + 1)

	dp[0] = 1
	for i := 1 ; i <= n ; i++ {
		for j := k ; j >= nums[i - 1] ; j-- {
			dp[j] += dp[j - nums[i - 1]]
		}
	}
	return dp[k]
	//----------------------------------------------------------------------------
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
