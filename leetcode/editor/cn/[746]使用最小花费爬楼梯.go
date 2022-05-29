//给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。 
//
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。 
//
// 请你计算并返回达到楼梯顶部的最低花费。 
//
// 
//
// 示例 1： 
//
// 
//输入：cost = [10,15,20]
//输出：15
//解释：你将从下标为 1 的台阶开始。
//- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
//总花费为 15 。
// 
//
// 示例 2： 
//
// 
//输入：cost = [1,100,1,1,1,100,1,1,100,1]
//输出：6
//解释：你将从下标为 0 的台阶开始。
//- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
//- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
//- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
//总花费为 6 。
// 
//
// 
//
// 提示： 
//
// 
// 2 <= cost.length <= 1000 
// 0 <= cost[i] <= 999 
// 
// Related Topics 数组 动态规划 👍 879 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	//----------------------------------递归---------------------------------------
	/*// Time: O(n)
	// Space: O(1)
	var dfs func(cost []int, i int) int
	dfs = func(cost []int, i int) int {
		if i == 0 || i == 1 {
			return cost[i]
		}
		return min(dfs(cost, i - 1), dfs(cost, i - 2)) + cost[i]
	}
	return min(dfs(cost, len(cost) - 1), dfs(cost, len(cost) - 2))*/
	//---------------------------------------------------------------------------

	//---------------------------------记忆化搜索---------------------------------------
	/*// Time: O(n)
	// Space: O(n)
	hashMap := map[int]int{}
	var dfs func(cost []int, i int) int
	dfs = func(cost []int, i int) int {
		if i == 0 || i == 1 {
			return cost[i]
		}
		p1, p2 := 0, 0
		if v1, ok := hashMap[i - 1]; ok {
			p1 = v1
		}else{
			p1 = dfs(cost, i - 1)
		}
		if v2, ok := hashMap[i - 2]; ok {
			p2 = v2
		}else{
			p2 = dfs(cost, i - 2)
		}
		hashMap[i] = min(p1, p2) + cost[i]
		return min(p1, p2) + cost[i]
	}
	return min(dfs(cost, len(cost) - 1), dfs(cost, len(cost) - 2))*/
	//---------------------------------------------------------------------------

	//-------------------------------动态规划------------------------------------
	/*// Time: O(n)
	// Space: O(n)
	n := len(cost)
	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 0
	for i := 2 ; i <= n ; i++ {
		dp[i] = min(dp[i - 2] + cost[i - 2], dp[i - 1] + cost[i - 1])
	}
	return dp[n]*/
	//---------------------------------------------------------------------------

	//----------------------------动态规划(空间优化)--------------------------------
	// Time: O(n)
	// Space: O(1)
	n := len(cost)
	a, b := 0, 0
	for i := 2 ; i <= n ; i++ {
		c := min(a + cost[i - 2], b + cost[i - 1])
		a, b = b, c
	}
	return b
	//---------------------------------------------------------------------------
}

func min(a , b int) int {
	if a > b {
		return b
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
