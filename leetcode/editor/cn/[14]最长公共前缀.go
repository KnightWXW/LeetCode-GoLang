//编写一个函数来查找字符串数组中的最长公共前缀。 
//
// 如果不存在公共前缀，返回空字符串 ""。 
//
// 
//
// 示例 1： 
//
// 
//输入：strs = ["flower","flow","flight"]
//输出："fl"
// 
//
// 示例 2： 
//
// 
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。 
//
// 
//
// 提示： 
//
// 
// 1 <= strs.length <= 200 
// 0 <= strs[i].length <= 200 
// strs[i] 仅由小写英文字母组成 
// 
// Related Topics 字符串 👍 2178 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	//----------------------------贪心 + 纵向比较-----------------------------------
	// Time: O(n * m)
	// Space: O(1)
	for j := 0 ; j < len(strs[0]) ; j++ {
		k := strs[0][j]
		for i := 1 ; i < len(strs) ; i++ {
			if j >= len(strs[i]) || strs[i][j] != k {
				return string(strs[0][ : j])
			}
		}
	}
	return strs[0]
	//----------------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
