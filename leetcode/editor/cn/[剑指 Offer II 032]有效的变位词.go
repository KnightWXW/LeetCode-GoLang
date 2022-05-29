//给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。 
//
// 注意：若 s 和 t 中每个字符出现的次数都相同且字符顺序不完全相同，则称 s 和 t 互为变位词（字母异位词）。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "anagram", t = "nagaram"
//输出: true
// 
//
// 示例 2: 
//
// 
//输入: s = "rat", t = "car"
//输出: false 
//
// 示例 3: 
//
// 
//输入: s = "a", t = "a"
//输出: false 
//
// 
//
// 提示: 
//
// 
// 1 <= s.length, t.length <= 5 * 10⁴ 
// s and t 仅包含小写字母 
// 
//
// 
//
// 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？ 
//
// 
//
// 注意：本题与主站 242 题相似（字母异位词定义不同）：https://leetcode-cn.com/problems/valid-anagram/ 
// Related Topics 哈希表 字符串 排序 👍 17 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	//--------------------------------------------------------------//-------------------------------排序----------------------------------
	/*// Time: O(n)
	// Space: O(n)
	if s == t {
		return false
	}
	sArr, tArr := []byte(s), []byte(t)
	sort.Slice(sArr, func(i , j int) bool {return sArr[i] < sArr[j]})
	sort.Slice(tArr, func(i , j int) bool {return tArr[i] < tArr[j]})
	return string(sArr) == string(tArr)*/
	//-------------------------------------------------------------------

	//-------------------------------排序----------------------------------
	// Time: O(n)
	// Space: O(26)
	if s == t {
		return false
	}
	sArr, tArr := [26]int{}, [26]int{}
	if len(s) != len(t) {
		return false
	}
	for i := 0 ; i < len(s) ; i++ {
		sArr[s[i] - 'a']++
		tArr[t[i] - 'a']++
	}
	for i := 0 ; i < 26 ; i++ {
		if sArr[i] != tArr[i] {
			return false
		}
	}
	return true
	//-------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
