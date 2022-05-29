//给定一个非负整数数组 nums， nums 中一半整数是 奇数 ，一半整数是 偶数 。 
//
// 对数组进行排序，以便当 nums[i] 为奇数时，i 也是 奇数 ；当 nums[i] 为偶数时， i 也是 偶数 。 
//
// 你可以返回 任何满足上述条件的数组作为答案 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [4,2,5,7]
//输出：[4,5,2,7]
//解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
// 
//
// 示例 2： 
//
// 
//输入：nums = [2,3]
//输出：[2,3]
// 
//
// 
//
// 提示： 
//
// 
// 2 <= nums.length <= 2 * 10⁴ 
// nums.length 是偶数 
// nums 中一半是偶数 
// 0 <= nums[i] <= 1000 
// 
//
// 
//
// 进阶：可以不使用额外空间解决问题吗？ 
// Related Topics 数组 双指针 排序 👍 235 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParityII(nums []int) []int {
	//------------------------------两次遍历----------------------------------
	/*// Time: O(n)
	// Space: O(1)
	n := len(nums)
	ans := make([]int, n)
	index := 0
	for _, v := range nums {
		if v & 1 == 0 {
			ans[index] = v
			index += 2
		}
	}

	index = 1
	for _, v := range nums {
		if v & 1 == 1 {
			ans[index] = v
			index += 2
		}
	}

	return ans*/
	//-----------------------------------------------------------------------

	//--------------------------------双指针----------------------------------
	// Time: O(n)
	// Space: O(1)
	n := len(nums)
	for i, j := 0, 1 ; i < n ; i += 2 {
		if nums[i] & 1 == 1 {
			for nums[j] & 1 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return nums
	//-------------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
