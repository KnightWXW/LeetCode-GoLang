//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。 
//
// 
//
// 示例 1： 
//
// 输入：s = "We are happy."
//输出："We%20are%20happy." 
//
// 
//
// 限制： 
//
// 0 <= s 的长度 <= 10000 
// Related Topics 字符串 👍 291 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	//------------------------------内置函数--------------------------------
	// Time: O(n)
	// Space: O(1)
	return strings.ReplaceAll(s, " ", "%20")
	//-------------------------------------------------------------------

	//------------------------------遍历模拟--------------------------------
	// Time: O(n)
	// Space: O(n)
	ans := []byte{}
	for i := 0 ; i < len(s) ; i++ {
		if s[i] == ' ' {
			ans = append(ans, '%','2','0')
		}else{
			ans = append(ans, s[i])
		}
	}
	return string(ans)
	//-------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
