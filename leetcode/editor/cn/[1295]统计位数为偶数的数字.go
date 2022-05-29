//给你一个整数数组 nums，请你返回其中位数为 偶数 的数字的个数。 
//
// 
//
// 示例 1： 
//
// 输入：nums = [12,345,2,6,7896]
//输出：2
//解释：
//12 是 2 位数字（位数为偶数） 
//345 是 3 位数字（位数为奇数）  
//2 是 1 位数字（位数为奇数） 
//6 是 1 位数字 位数为奇数） 
//7896 是 4 位数字（位数为偶数）  
//因此只有 12 和 7896 是位数为偶数的数字
// 
//
// 示例 2： 
//
// 输入：nums = [555,901,482,1771]
//输出：1 
//解释： 
//只有 1771 是位数为偶数的数字。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 500 
// 1 <= nums[i] <= 10^5 
// 
// Related Topics 数组 👍 70 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findNumbers(nums []int) int {
	//-----------------------------模拟----------------------------------
	/*// Time: O(n)
	// Space: O(1)
	var digit func(a int) int
	digit = func(a int) int {
		ans := 0
		for a != 0 {
			a /= 10
			ans++
		}
		return ans
	}

	ans := 0
	for _, v := range nums {
		if digit(v) & 1 == 0 {
			ans++
		}
	}
	return ans*/
	//---------------------------------------------------------------

	//---------------------------转换成string--------------------------
	// Time: O(n)
	// Space: O(1)
	ans := 0
	for _, v := range nums {
		tem := strconv.Itoa(v)
		if len(tem) & 1 == 0 {
			ans++
		}
	}
	return ans
	//-----------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
