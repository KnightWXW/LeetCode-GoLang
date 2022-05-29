//整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。 
//
// 示例1: 
//
// 
// 输入：A = 29 （或者0b11101）, B = 15（或者0b01111）
// 输出：2
// 
//
// 示例2: 
//
// 
// 输入：A = 1，B = 2
// 输出：2
// 
//
// 提示: 
//
// 
// A，B范围在[-2147483648, 2147483647]之间 
// 
// Related Topics 位运算 👍 41 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func convertInteger(A int, B int) int {
	//----------------------------异或------------------------------
	// Time: O(n)
	// Space: O(1)
	// 注意 异或结果 转为 int32 类型
	k, ans := int32(A ^ B), 0
	for k != 0 {
		k &= (k - 1)
		ans++
	}
	return ans
	//--------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
