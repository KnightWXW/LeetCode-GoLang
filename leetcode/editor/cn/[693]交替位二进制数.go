//给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 5
//输出：true
//解释：5 的二进制表示是：101
// 
//
// 示例 2： 
//
// 
//输入：n = 7
//输出：false
//解释：7 的二进制表示是：111. 
//
// 示例 3： 
//
// 
//输入：n = 11
//输出：false
//解释：11 的二进制表示是：1011. 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 2³¹ - 1 
// 
// Related Topics 位运算 👍 157 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func hasAlternatingBits(n int) bool {
	//--------------------------------转换为字符串--------------------------------
	/*// Time: O(n)
	// Space: O(n)
	arr := []int{}
	tem := n
	for tem != 0 {
		a := tem & 1
		tem >>= 1
		arr = append(arr, a)
	}
	for i := 0 ; i < len(arr) - 1 ; i++ {
		if arr[i] == arr[i + 1] {
			return false
		}
	}
	return true*/
	//-------------------------------------------------------------------------

	//---------------------------------模拟-------------------------------------
	/*// Time: O(n)
	// Space: O(1)
	pre := -1
	for n != 0 {
		cur := n & 1
		if pre == cur {
			return false
		}
		pre = cur
		n >>= 1
	}
	return true*/
	//-------------------------------------------------------------------------

	//---------------------------------模拟-------------------------------------
	// Time: O(n)
	// Space: O(1)
	tem := n ^ (n >> 1)
	return tem & (tem + 1) == 0
	//-------------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
