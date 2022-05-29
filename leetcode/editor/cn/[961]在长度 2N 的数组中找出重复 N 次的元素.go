//给你一个整数数组 nums ，该数组具有以下属性： 
//
// 
// 
// 
// nums.length == 2 * n. 
// nums 包含 n + 1 个 不同的 元素 
// nums 中恰有一个元素重复 n 次 
// 
//
// 找出并返回重复了 n 次的那个元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3,3]
//输出：3
// 
//
// 示例 2： 
//
// 
//输入：nums = [2,1,2,5,3,2]
//输出：2
// 
//
// 示例 3： 
//
// 
//输入：nums = [5,1,5,2,5,3,5,4]
//输出：5
// 
// 
// 
//
// 
//
// 提示： 
//
// 
// 2 <= n <= 5000 
// nums.length == 2 * n 
// 0 <= nums[i] <= 10⁴ 
// nums 由 n + 1 个 不同的 元素组成，且其中一个元素恰好重复 n 次 
// 
// Related Topics 数组 哈希表 👍 108 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func repeatedNTimes(nums []int) int {
	//--------------------------------排序--------------------------------------
	/*// Time: O(nlogn)
	// Space: O(logn)
	sort.Ints(nums)
	n := len(nums)
	if nums[n / 2] == nums[n / 2 + 1] {
		return nums[n / 2]
	}else{
		return nums[n / 2 - 1]
	}*/
	//-------------------------------------------------------------------------

	//--------------------------------hashMap----------------------------------
	/*// Time: O(n)
	// Space: O(n)
	hashMap := map[int]int{}
	for i := 0 ; i < len(nums) ; i++ {
		hashMap[nums[i]]++
		if hashMap[nums[i]] > 1 {
			return nums[i]
		}
	}
	return -1*/
	//-------------------------------------------------------------------------

	//----------------------------数学推导 + 位运算-------------------------------
	// Time: O(n)
	// Space: O(1)
	for i := 1 ; i < len(nums) ; i++ {
		if nums[i] ^ nums[i - 1] == 0 {
			return nums[i]
		}
	}
	if nums[0] == nums[2] || nums[0] == nums[3] {
		return nums[0]
	}else{
		return nums[1]
	}
	//-------------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
