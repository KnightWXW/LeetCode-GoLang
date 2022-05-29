//找出数组中重复的数字。 
//
// 
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请
//找出数组中任意一个重复的数字。 
//
// 示例 1： 
//
// 输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3 
// 
//
// 
//
// 限制： 
//
// 2 <= n <= 100000 
// Related Topics 数组 哈希表 排序 👍 829 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findRepeatNumber(nums []int) int {
	//-----------------------------排序----------------------------------
	/*// Time: O(nlogn)
	// Space: O(logn)
	sort.Ints(nums)
	for i := 1 ; i < len(nums) ; i++ {
		if nums[i - 1] == nums[i] {
			return nums[i]
		}
	}
	return -1*/
	//-------------------------------------------------------------------

	//-----------------------------哈希表---------------------------------
	/*// Time: O(n)
	// Space: O(n)
	hashSet := map[int]bool{}
	for _, v := range nums {
		if hashSet[v] {
			return v
		}
		hashSet[v] = true
	}
	return -1*/
	//-------------------------------------------------------------------

	//-------------------------------位图---------------------------------
	// Time: O(n)
	// Space: O(n >> 6)
	n := len(nums)
	bitArr := make([]int, n >> 6 + 1)
	for _, v := range nums {
		if bitArr[v >> 6] & (1 << (v & 63)) != 0 {
			return v
		}
		bitArr[v >> 6] |= (1 << (v & 63))
	}
	return -1
	//-------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
