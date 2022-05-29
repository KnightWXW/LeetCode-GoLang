//给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "Hello"
//输出："hello"
// 
//
// 示例 2： 
//
// 
//输入：s = "here"
//输出："here"
// 
//
// 示例 3： 
//
// 
//输入：s = "LOVELY"
//输出："lovely"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 100 
// s 由 ASCII 字符集中的可打印字符组成 
// 
// Related Topics 字符串 👍 204 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func toLowerCase(s string) string {
	//---------------------------遍历--------------------------------
	/*// Time: O(n)
	// Space: O(1)
	str := []byte(s)
	for i := 0 ; i < len(str) ; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = byte(str[i] - 'A' + 'a')
		}
	}
	return string(str)*/
	//-------------------------------------------------------------

	//---------------------------内置函数---------------------------
	// Time: O(n)
	// Space: O(1)
	return strings.ToLower(s)
	//-------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
