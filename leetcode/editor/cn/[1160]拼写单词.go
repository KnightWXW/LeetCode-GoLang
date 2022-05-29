//给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。 
//
// 假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。 
//
// 注意：每次拼写（指拼写词汇表中的一个单词）时，chars 中的每个字母都只能用一次。 
//
// 返回词汇表 words 中你掌握的所有单词的 长度之和。 
//
// 
//
// 示例 1： 
//
// 输入：words = ["cat","bt","hat","tree"], chars = "atach"
//输出：6
//解释： 
//可以形成字符串 "cat" 和 "hat"，所以答案是 3 + 3 = 6。
// 
//
// 示例 2： 
//
// 输入：words = ["hello","world","leetcode"], chars = "welldonehoneyr"
//输出：10
//解释：
//可以形成字符串 "hello" 和 "world"，所以答案是 5 + 5 = 10。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= words.length <= 1000 
// 1 <= words[i].length, chars.length <= 100 
// 所有字符串中都仅包含小写英文字母 
// 
// Related Topics 数组 哈希表 字符串 👍 154 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func countCharacters(words []string, chars string) int {
	//---------------------------------哈希表-------------------------------------
	/*// Time: O(n)
	// Space: O(n)
	hashMap := map[byte]int{}
	ans := 0
	for _, v := range chars {
		hashMap[byte(v)]++
	}
	for i := 0 ; i < len(words) ; i++ {
		mp := map[byte]int{}
		for j := 0 ; j < len(words[i]) ; j++ {
			mp[byte(words[i][j])]++
		}
		flag := true
		for j := 0 ; j < len(words[i]) ; j++ {
			if mp[byte(words[i][j])] > hashMap[byte((words[i][j]))] {
				flag = false
				break
			}
		}
		if flag == true {
			ans += len(words[i])
		}
	}
	return ans*/
	//---------------------------------------------------------------------------

	//---------------------------------数组哈希------------------------------------
	// Time: O(n)
	// Space: O(1)
	hashArr := [26]int{}
	ans := 0
	for _, v := range chars {
		hashArr[int(v - 'a')]++
	}
	for i := 0 ; i < len(words) ; i++ {
		arr := [26]int{}
		for j := 0 ; j < len(words[i]) ; j++ {
			arr[int(words[i][j] - 'a')]++
		}
		flag := true
		for j := 0 ; j < len(words[i]) ; j++ {
			if arr[int(words[i][j] - 'a')] > hashArr[int(words[i][j] - 'a')] {
				flag = false
				break
			}
		}
		if flag == true {
			ans += len(words[i])
		}
	}
	return ans
	//---------------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
