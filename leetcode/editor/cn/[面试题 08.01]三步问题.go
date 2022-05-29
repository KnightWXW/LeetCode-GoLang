//三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模100
//0000007。 
//
// 示例1: 
//
// 
// 输入：n = 3 
// 输出：4
// 说明: 有四种走法
// 
//
// 示例2: 
//
// 
// 输入：n = 5
// 输出：13
// 
//
// 提示: 
//
// 
// n范围在[1, 1000000]之间 
// 
// Related Topics 记忆化搜索 数学 动态规划 👍 79 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func waysToStep(n int) int {
	//---------------------------递归(超时)---------------------------------
	/*// Time: O(n)
	// Space: O(n)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 	4
	}
	return (waysToStep(n - 1) + waysToStep(n - 2) + waysToStep(n - 3)) % 1000000007*/
	//---------------------------------------------------------------

	//-------------------------记忆化搜索------------------------------
	/*// Time: O(n)
	// Space: O(n)
	mapArr := make([]int, n + 1)

	var waysToStepDFS func(n int) int

	waysToStepDFS = func(n int) int {
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}
		if n == 3 {
			return 	4
		}

		p1, p2, p3 := 0, 0, 0

		if mapArr[n - 1] != 0 {
			p1 = mapArr[n - 1]
		}else{
			p1 = waysToStepDFS(n - 1)
		}

		if mapArr[n - 2] != 0 {
			p2 = mapArr[n - 2]
		}else{
			p2 = waysToStepDFS(n - 2)
		}

		if mapArr[n - 3] != 0 {
			p3 = mapArr[n - 3]
		}else{
			p3 = waysToStepDFS(n - 3)
		}
		mapArr[n] = (p1 + p2 + p3) % 1000000007
		return (p1 + p2 + p3) % 1000000007
	}

	return waysToStepDFS(n)*/
	//---------------------------------------------------------------

	//-------------------------动态规划------------------------------
	/*// Time: O(n)
	// Space: O(1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 	4
	}
	dp := make([]int, n + 1)
	dp[1], dp[2], dp[3] = 1, 2, 4
	for i := 4 ; i <= n ; i++ {
		dp[i] = (dp[i - 1] + dp[i - 2] + dp[i - 3]) % 1000000007
	}
	return dp[n]*/
	//---------------------------------------------------------------

	//-------------------------动态规划------------------------------
	// Time: O(n)
	// Space: O(1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 	4
	}
	a, b, c := 1, 2, 4
	for i := 4 ; i <= n ; i++ {
		d := (a + b + c) % 1000000007
		a, b, c = b, c, d
	}
	return c
	//---------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
