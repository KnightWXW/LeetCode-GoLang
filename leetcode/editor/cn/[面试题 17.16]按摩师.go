//一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩
//师找到最优的预约集合（总预约时间最长），返回总的分钟数。 
//
// 注意：本题相对原题稍作改动 
//
// 
//
// 示例 1： 
//
// 输入： [1,2,3,1]
//输出： 4
//解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。
// 
//
// 示例 2： 
//
// 输入： [2,7,9,3,1]
//输出： 12
//解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。
// 
//
// 示例 3： 
//
// 输入： [2,1,4,5,3,1,1,3]
//输出： 12
//解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。
// 
// Related Topics 数组 动态规划 👍 237 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func massage(nums []int) int {
	//-----------------------------递归1(超时)----------------------------------
	/*// Time: O(n)
	// Space: O(n)
	var massageDFS func(nums []int, left, right int) int
	massageDFS = func(nums []int, left, right int) int {
		if left > right {
			return 0
		}
		p1, p2 := massageDFS(nums, left + 1, right), massageDFS(nums, left + 2, right) + nums[left]
		return max(p1, p2)
	}

	return massageDFS(nums, 0, len(nums) - 1)*/
	//------------------------------------------------------------------------

	//-----------------------------递归2(超时)----------------------------------
	/*// Time: O(n)
	// Space: O(n)
	var massageDFS func(nums []int, left, right int) int
	massageDFS = func(nums []int, left, right int) int {
		if left > right {
			return 0
		}
		p1, p2 := massageDFS(nums, left, right - 1), massageDFS(nums, left, right - 2) + nums[right]
		return max(p1, p2)
	}

	return massageDFS(nums, 0, len(nums) - 1)*/
	//----------------------------------------------------------------------

	//----------------------------记忆化搜索(超时)-----------------------------
	/*// Time: O(n)
	// Space: O(n)
	n := len(nums)
	hashMap := map[string]int{}

	var massageDFS func(nums []int, left, right int) int
	massageDFS = func(nums []int, left, right int) int {
		if left > right {
			return 0
		}
		key := strconv.Itoa(left) + "#" + strconv.Itoa(right)
		p1, p2 := 0, 0
		if hashMap[strconv.Itoa(left) + "#" + strconv.Itoa(right - 1)] > 0 {
			p1 = hashMap[strconv.Itoa(left) + "#" + strconv.Itoa(right - 1)]
		}else{
			p1 = massageDFS(nums, left, right - 1)
		}

		if hashMap[strconv.Itoa(left) + "#" + strconv.Itoa(right - 2)] > 0 {
			p2 = hashMap[strconv.Itoa(left) + "#" + strconv.Itoa(right - 2)] + nums[right]
		}else{
			p2 = massageDFS(nums, left, right - 2) + nums[right]
		}
		hashMap[key] = max(p1, p2)
		return max(p1, p2)
	}

	return massageDFS(nums, 0, n - 1)*/
	//--------------------------------------------------------------------

	//----------------------------动态规划-----------------------------
	/*// Time: O(n)
	// Space: O(n)
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int , n + 1)
	dp[1] = nums[0]
	dp[2] = max(nums[0], nums[1])

	for i := 3 ; i <= n ; i++ {
		dp[i] = max(dp[i - 1], dp[i - 2] + nums[i - 1])
	}
	return dp[n]*/
	//--------------------------------------------------------------------

	//--------------------------动态规划(空间优化)-----------------------------
	// Time: O(n)
	// Space: O(1)
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	a, b := nums[0], max(nums[0], nums[1])

	for i := 3 ; i <= n ; i++ {
		c := max(b, a + nums[i - 1])
		a, b = b, c
	}
	return b
	//--------------------------------------------------------------------
}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
