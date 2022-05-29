//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。 
//
// 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100] 
//
// 示例 2： 
//
// 
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁴ 
// -10⁴ <= nums[i] <= 10⁴ 
// nums 已按 非递减顺序 排序 
// 
//
// 
//
// 进阶： 
//
// 
// 请你设计时间复杂度为 O(n) 的算法解决本问题 
// 
// Related Topics 数组 双指针 排序 👍 497 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func sortedSquares(nums []int) []int {
	//---------------------------------排序-------------------------------------
	/*// Time: O(nlogn)
	// Space: O(logn)
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})

	for i := 0 ; i < len(nums) ; i++ {
		nums[i] *= nums[i]
	}

	return nums*/
	//-------------------------------------------------------------------------

	//---------------------------------双指针-------------------------------------
	// Time: O(n)
	// Space: O(n)
	ans := make([]int, len(nums))
	left, right := 0, len(nums) - 1
	index := len(nums) - 1
	for left <= right {
		if nums[left] * nums[left] < nums[right] * nums[right] {
			ans[index] = nums[right] * nums[right]
			right--
		}else{
			ans[index] = nums[left] * nums[left]
			left++
		}
		index--
	}
	return ans
	//-------------------------------------------------------------------------
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
