//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。 
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "hello"
//输出："holle"
// 
//
// 示例 2： 
//
// 
//输入：s = "leetcode"
//输出："leotcede" 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 3 * 10⁵ 
// s 由 可打印的 ASCII 字符组成 
// 
// Related Topics 双指针 字符串 👍 244 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	//--------------------------------双指针---------------------------------------
	// Time: O(n)
	// Space: O(1)
	n := len(s)
	arr := []byte(s)
	left, right := 0, n - 1
	for left < right {
		for left < right && judge(arr[left]) == false {
			left++
		}
		for left < right && judge(arr[right]) == false {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return string(arr)
	//----------------------------------------------------------------------------
}

func judge(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
