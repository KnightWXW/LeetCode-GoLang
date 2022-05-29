//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。 
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：[3,2,3]
//输出：3 
//
// 示例 2： 
//
// 
//输入：[2,2,1,1,1,2,2]
//输出：2
// 
//
// 
//
// 进阶： 
//
// 
// 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。 
// 
// Related Topics 数组 哈希表 分治 计数 排序 👍 1366 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	//---------------------------------排序-------------------------------
	/*// Time: O(nlogn)
	// Space: O(logn)
	sort.Ints(nums)
	return nums[len(nums) / 2]*/
	//--------------------------------------------------------------------

	//-------------------------------摩尔投票-------------------------------
	// Time: O(n)
	// Space: O(1)
	vote, candidate := 0, 0
	for _, v := range nums {
		if vote == 0 {
			candidate = v
		}
		if v == candidate {
			vote++
		}else{
			vote--
		}
	}
	return candidate
	//--------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
