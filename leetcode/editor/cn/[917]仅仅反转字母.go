//给你一个字符串 s ，根据下述规则反转字符串： 
//
// 
// 所有非英文字母保留在原有位置。 
// 所有英文字母（小写或大写）位置反转。 
// 
//
// 返回反转后的 s 。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 
//输入：s = "ab-cd"
//输出："dc-ba"
// 
//
// 
// 
//
// 示例 2： 
//
// 
//输入：s = "a-bC-dEf-ghIj"
//输出："j-Ih-gfE-dCba"
// 
//
// 
// 
//
// 示例 3： 
//
// 
//输入：s = "Test1ng-Leet=code-Q!"
//输出："Qedo1ct-eeLg=ntse-T!"
// 
//
// 
//
// 提示 
//
// 
// 1 <= s.length <= 100 
// s 仅由 ASCII 值在范围 [33, 122] 的字符组成 
// s 不含 '\"' 或 '\\' 
// 
// Related Topics 双指针 字符串 👍 167 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func reverseOnlyLetters(s string) string {
	//----------------------------------双指针----------------------------------------
	// Time: O(n)
	// Space: O(n)
	arr := []byte(s)
	left, right := 0 , len(arr) - 1
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
	//-------------------------------------------------------------------------------
}

func judge(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
