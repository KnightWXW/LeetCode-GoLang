//设计一个函数把两个数字相加。不得使用 + 或者其他算术运算符。 
//
// 示例: 
//
// 输入: a = 1, b = 1
//输出: 2 
//
// 
//
// 提示： 
//
// 
// a, b 均可能是负数或 0 
// 结果不会溢出 32 位整数 
// 
// Related Topics 位运算 数学 👍 73 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func add(a int, b int) int {
	//-------------------------------位运算---------------------------------
	// Time: O(n)
	// Space: O(1)
	for b != 0 {
		c := a ^ b
		b = (a & b) << 1
		a = c
	}
	return a
	//--------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
