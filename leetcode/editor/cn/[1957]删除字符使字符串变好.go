//一个字符串如果没有 三个连续 相同字符，那么它就是一个 好字符串 。 
//
// 给你一个字符串 s ，请你从 s 删除 最少 的字符，使它变成一个 好字符串 。 
//
// 请你返回删除后的字符串。题目数据保证答案总是 唯一的 。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "leeetcode"
//输出："leetcode"
//解释：
//从第一组 'e' 里面删除一个 'e' ，得到 "leetcode" 。
//没有连续三个相同字符，所以返回 "leetcode" 。
// 
//
// 示例 2： 
//
// 
//输入：s = "aaabaaaa"
//输出："aabaa"
//解释：
//从第一组 'a' 里面删除一个 'a' ，得到 "aabaaaa" 。
//从第二组 'a' 里面删除两个 'a' ，得到 "aabaa" 。
//没有连续三个相同字符，所以返回 "aabaa" 。
// 
//
// 示例 3： 
//
// 
//输入：s = "aab"
//输出："aab"
//解释：没有连续三个相同字符，所以返回 "aab" 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁵ 
// s 只包含小写英文字母。 
// 
// Related Topics 字符串 👍 17 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func makeFancyString(s string) string {
	//----------------------------暴力模拟-------------------------------
	// Time: O(n)
	// Space: O(1)
	cnt := 1
	ans := []rune{}
	for i := 0 ; i < len(s) ; i++ {
		if i + 1 < len(s) && s[i] == s[i + 1] {
			cnt++
		}else{
			cnt = 1
		}
		if cnt <= 2 {
			ans	= append(ans , rune(s[i]))
		}
	}
	return string(ans)
	//---------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
