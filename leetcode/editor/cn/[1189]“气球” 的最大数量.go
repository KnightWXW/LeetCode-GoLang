//给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。 
//
// 字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。 
//
// 
//
// 示例 1： 
//
// 
//
// 输入：text = "nlaebolko"
//输出：1
// 
//
// 示例 2： 
//
// 
//
// 输入：text = "loonbalxballpoon"
//输出：2
// 
//
// 示例 3： 
//
// 输入：text = "leetcode"
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= text.length <= 10^4 
// text 全部由小写英文字母组成 
// 
// Related Topics 哈希表 字符串 计数 👍 109 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxNumberOfBalloons(text string) int {
	//-------------------------------模拟--------------------------------------
	// Time: O(n)
	// Space: O(1)
	b, a, l, o, n := 0, 0, 0, 0, 0
	for _, v := range text {
		switch v {
			case 'b':
				b++
			case 'a':
				a++
			case 'l':
				l++
			case 'o':
				o++
			case 'n':
				n++
		}
	}
	return min(b, min(a, min(l / 2, min(o / 2, n))))
	//-------------------------------------------------------------------------
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
