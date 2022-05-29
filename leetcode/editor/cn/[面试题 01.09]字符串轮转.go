//字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。 
//
// 示例1: 
//
//  输入：s1 = "waterbottle", s2 = "erbottlewat"
// 输出：True
// 
//
// 示例2: 
//
//  输入：s1 = "aa", s2 = "aba"
// 输出：False
// 
//
// 
// 
//
// 提示： 
//
// 
// 字符串长度在[0, 100000]范围内。 
// 
//
// 说明: 
//
// 
// 你能只调用一次检查子串的方法吗？ 
// 
// Related Topics 字符串 字符串匹配 👍 116 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isFlipedString(s1 string, s2 string) bool {
	//-----------------------------搜索子字符串-------------------------------------
	// Time: O(n)
	// Space: O(n)
	return len(s1) == len(s2) && strings.Contains(s2 + s2,s1)
	//---------------------------------------------------------------------------

	//-----------------------------模拟--------------------------------------
	/*// Time: O(n ^ 2)
	// Space: O(1)
	if len(s1) != len(s2) {
		return false
	}
	n := len(s1)
	if n == 0 {
		return true
	}
	for i := 0 ; i < n ; i++ {
		j := 0
		for ; j < n ; j++ {
			if s1[j] != s2[(j + i) % n] {
				break
			}
		}
		if j == n {
			return true
		}
	}
	return false*/
	//---------------------------------------------------------------------------

	//-----------------------------模拟--------------------------------------
	/*// Time: O(n)
	// Space: O(n)
	if len(s1) != len(s2) {
		return false
	}
	n := len(s1)
	if n == 0 {
		return true
	}
	for i := 0 ; i < n ; i++ {
		tem := s1[i : ] + s1[ : i]
		if tem == s2 {
			return true
		}
	}
	return false*/
	//---------------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
