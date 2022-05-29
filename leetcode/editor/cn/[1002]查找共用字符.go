//给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答
//案。
// 
//
// 示例 1： 
//
// 
//输入：words = ["bella","label","roller"]
//输出：["e","l","l"]
// 
//
// 示例 2： 
//
// 
//输入：words = ["cool","lock","cook"]
//输出：["c","o"]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= words.length <= 100 
// 1 <= words[i].length <= 100 
// words[i] 由小写英文字母组成 
// 
// Related Topics 数组 哈希表 字符串 👍 264 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func commonChars(words []string) []string {
	//-----------------------------哈希表---------------------------------
	// Time: 0(n * 26)
	// Space: 0(26)
	arr := [26]int{}
	for i := range arr {
		arr[i] = math.MaxInt64
	}
	for _, v := range words {
		tem := [26]int{}
		for i := 0 ; i < len(v) ; i++ {
			tem[v[i] - 'a']++
		}
		for i := 0 ; i < 26 ; i++ {
			arr[i] = min(arr[i], tem[i])
		}
	}

	ans := []string{}
	for i := 0 ; i < 26 ; i++ {
		for arr[i] > 0 {
			ans = append(ans , string(i + 'a'))
			arr[i]--
		}
	}
	return ans
	//------------------------------------------------------------------
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
