//给定一个整数数组，找出总和最大的连续数列，并返回总和。 
//
// 示例： 
//
// 输入： [-2,1,-3,4,-1,2,1,-5,4]
//输出： 6
//解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 
//
// 进阶： 
//
// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。 
// Related Topics 数组 分治 动态规划 👍 110 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	//-------------------------贪心算法1-----------------------------
	/*// Time: O(n)
	// Space: O(1)
	ans , tem := math.MinInt64, 0
	for i := 0 ; i < len(nums) ; i++ {
		if nums[i] > tem + nums[i] {
			tem = nums[i]
		}else{
			tem += nums[i]
		}
		if tem > ans {
			ans = tem
		}
	}
	return ans*/
	//-----------------------------------------------------------

	//-------------------------贪心算法2---------------------------
	// Time: O(n)
	// Space: O(1)
	ans, tem := math.MinInt64, 0
	for i := 0 ; i < len(nums) ; i++ {
		if tem < 0 {
			tem = nums[i]
		}else{
			tem += nums[i]
		}
		if tem > ans {
			ans = tem
		}
	}
	return ans
	//-----------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
