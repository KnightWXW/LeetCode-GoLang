//给你一个 仅 由字符 'a' 和 'b' 组成的字符串 s 。如果字符串中 每个 'a' 都出现在 每个 'b' 之前，返回 true ；否则，返回 
//false 。 
//
// 
//
// 示例 1： 
//
// 输入：s = "aaabbb"
//输出：true
//解释：
//'a' 位于下标 0、1 和 2 ；而 'b' 位于下标 3、4 和 5 。
//因此，每个 'a' 都出现在每个 'b' 之前，所以返回 true 。
// 
//
// 示例 2： 
//
// 输入：s = "abab"
//输出：false
//解释：
//存在一个 'a' 位于下标 2 ，而一个 'b' 位于下标 1 。
//因此，不能满足每个 'a' 都出现在每个 'b' 之前，所以返回 false 。
// 
//
// 示例 3： 
//
// 输入：s = "bbb"
//输出：true
//解释：
//不存在 'a' ，因此可以视作每个 'a' 都出现在每个 'b' 之前，所以返回 true 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 100 
// s[i] 为 'a' 或 'b' 
// 
// Related Topics 字符串 👍 5 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func checkString(s string) bool {
	//-------------------------------模拟------------------------------------
	/*// Time: O(n)
	// Space: O(1)
	a, b := -1, len(s)
	for i := len(s) - 1 ; i >= 0 ; i-- {
		if s[i] == 'a' {
			a = i
			break
		}
	}
	for i := 0 ; i < len(s) ; i++ {
		if s[i] == 'b' {
			b = i
			break
		}
	}
	return a < b*/
	//----------------------------------------------------------------------

	//------------------------------内置函数----------------------------------
	// Time: O(n)
	// Space: O(1)
	return  !strings.Contains(s, "ba")
	//----------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
