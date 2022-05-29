//给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。 
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。 
//
// 
//
// 示例1: 
//
// 
//输入: pattern = "abba", str = "dog cat cat dog"
//输出: true 
//
// 示例 2: 
//
// 
//输入:pattern = "abba", str = "dog cat cat fish"
//输出: false 
//
// 示例 3: 
//
// 
//输入: pattern = "aaaa", str = "dog cat cat dog"
//输出: false 
//
// 
//
// 提示: 
//
// 
// 1 <= pattern.length <= 300 
// pattern 只包含小写英文字母 
// 1 <= s.length <= 3000 
// s 只包含小写英文字母和 ' ' 
// s 不包含 任何前导或尾随对空格 
// s 中每个单词都被 单个空格 分隔 
// 
// Related Topics 哈希表 字符串 👍 459 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
	//-----------------------------哈希表(两次哈希)-----------------------------------
	// Time: O(m + n)
	// Space: O(m + n)
	arr := strings.Split(s, " ")
	hashMap1 := map[byte]string{}
	hashMap2 := map[string]byte{}

	if len(arr) != len(pattern) {
		return false
	}

	for i := 0 ; i < len(pattern) ; i++ {
		if v, ok := hashMap1[pattern[i]]; ok && v != arr[i] {
			return false
		}
		if v, ok := hashMap2[arr[i]]; ok && v != pattern[i] {
			return false
		}
		hashMap1[pattern[i]] = arr[i]
		hashMap2[arr[i]] = pattern[i]
	}
	return true
	//-----------------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
