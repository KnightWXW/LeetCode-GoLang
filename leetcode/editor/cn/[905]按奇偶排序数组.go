//给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。 
//
// 你可以返回满足此条件的任何数组作为答案。 
//
// 
//
// 示例： 
//
// 输入：[3,1,2,4]
//输出：[2,4,3,1]
//输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= A.length <= 5000 
// 0 <= A[i] <= 5000 
// 
// Related Topics 数组 双指针 排序 👍 240 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParity(nums []int) []int {
	//-----------------------------新建数组----------------------------------
	/*// Time: O(n)
	// Space: O(n)
	ans := []int{}
	for i := 0 ; i < len(nums) ; i++ {
		if nums[i] & 1 == 0 {
			ans = append(ans , nums[i])
		}
	}
	for i := 0 ; i < len(nums) ; i++ {
		if nums[i] & 1 == 1 {
			ans = append(ans , nums[i])
		}
	}
	return ans*/
	//----------------------------------------------------------------------

	//-----------------------------原始数组-----------------------------------
	// Time: O(n)
	// Space: O(1)
	left, right := 0, len(nums) - 1
	for left < right {
		if nums[left] & 1 > nums[right] & 1 {
			nums[left], nums[right] = nums[right], nums[left]
		}
		if nums[left] & 1 == 0 {
			left++
		}
		if nums[right] & 1 == 1 {
			right--
		}
	}
	return nums
	//----------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
