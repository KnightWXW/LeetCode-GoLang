//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [2,2,3,2]
//输出：3
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1,0,1,0,1,99]
//输出：99
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 3 * 10⁴ 
// -2³¹ <= nums[i] <= 2³¹ - 1 
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 
// 
//
// 
//
// 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？ 
// Related Topics 位运算 数组 👍 848 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	//-----------------------------------hashMap------------------------------------------
	/*// Time: O(n)
	// Space: O(n)
	hashMap := map[int]int{}
	for _, v := range nums {
		hashMap[v]++
	}
	for k, v := range hashMap {
		if v == 1 {
			return k
		}
	}
	return -1*/
	//------------------------------------------------------------------------------------

	//-----------------------------------按位取余------------------------------------------
	// Time: O(nlogC)
	// Space: O(1)
	ans := int32(0)
	for i := 0 ; i < 32 ; i++ {
		tem := int32(0)
		for _, v := range nums {
			tem += int32((v >> i) & 1)
		}
		ans += (tem % 3) << i
	}
	return int(ans)
	//------------------------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
