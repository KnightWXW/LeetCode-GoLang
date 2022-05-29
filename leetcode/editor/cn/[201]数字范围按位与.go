//给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）
//。 
//
// 
//
// 示例 1： 
//
// 
//输入：left = 5, right = 7
//输出：4
// 
//
// 示例 2： 
//
// 
//输入：left = 0, right = 0
//输出：0
// 
//
// 示例 3： 
//
// 
//输入：left = 1, right = 2147483647
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= left <= right <= 2³¹ - 1 
// 
// Related Topics 位运算 👍 377 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func rangeBitwiseAnd(left int, right int) int {
	//--------------------位运算模拟(超时)--------------------------
	/*// Time: O()
	// Space: O()
	ans := right
	for i := left ; i < right ; i++ {
		ans &= i
	}
	return ans*/
	//------------------------------------------------------------

	//-----------------------------移位---------------------------
	// Time: O(logn)
	// Space: O(1)
	flag := 0
	for left < right {
		left >>= 1
		right >>= 1
		flag++
	}
	return right << flag
	//------------------------------------------------------------

	//----------------------Brian Kernighan算法------------------------
	/*// Time: O(logn)
	// Space: O(1)
	// 除去 最后一位的 1
	for left < right {
		right &= (right - 1)
	}
	return right*/
	//------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
