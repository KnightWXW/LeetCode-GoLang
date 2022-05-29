//给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i 
//- j) <= k 。如果存在，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3,1], k = 3
//输出：true 
//
// 示例 2： 
//
// 
//输入：nums = [1,0,1,1], k = 1
//输出：true 
//
// 示例 3： 
//
// 
//输入：nums = [1,2,3,1,2,3], k = 2
//输出：false 
//
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁵ 
// -10⁹ <= nums[i] <= 10⁹ 
// 0 <= k <= 10⁵ 
// 
// Related Topics 数组 哈希表 滑动窗口 👍 474 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	//-----------------------------双层遍历------------------------------------
	/*// Time: O(n ^ 2)
	// Space: O(1)
	for i := 0; i < len(nums) ; i++ {
		for j := i + 1 ; j <= i + k && j < len(nums) ; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false*/
	//----------------------------------------------------------------------

	//-----------------------------哈希表-----------------------------------
	// Time: O(n)
	// Space: O(n)
	hashMap := make(map[int]int)
	for i, v := range nums {
		if v, ok := hashMap[v]; ok && i - v <= k{
			return true
		}
		hashMap[v] = i
	}
	return false
	//----------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
