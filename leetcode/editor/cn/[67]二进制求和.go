//给你两个二进制字符串，返回它们的和（用二进制表示）。 
//
// 输入为 非空 字符串且只包含数字 1 和 0。 
//
// 
//
// 示例 1: 
//
// 输入: a = "11", b = "1"
//输出: "100" 
//
// 示例 2: 
//
// 输入: a = "1010", b = "1011"
//输出: "10101" 
//
// 
//
// 提示： 
//
// 
// 每个字符串仅由字符 '0' 或 '1' 组成。 
// 1 <= a.length, b.length <= 10^4 
// 字符串如果不是 "0" ，就都不含前导零。 
// 
// Related Topics 位运算 数学 字符串 模拟 👍 790 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	//----------------------------模拟---------------------------------
	// Time: O(n)
	// Space: O(1)
	n := max(len(a), len(b))
	add := 0
	ans := ""
	for i := 0 ; i < n ; i++ {
		sum := add
		if i < len(a) {
			sum += int(a[len(a) - i - 1] - '0')
		}
		if i < len(b) {
			sum += int(b[len(b) - i - 1] - '0')
		}
		k := sum % 2
		add = sum / 2
		ans = strconv.Itoa(k) + ans
	}
	if add == 1 {
		ans = "1" + ans
	}
	return ans
	//----------------------------------------------------------------
}

func max(a , b int) int {
	if a < b {
		return b
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
