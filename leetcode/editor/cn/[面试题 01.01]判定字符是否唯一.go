//实现一个算法，确定一个字符串 s 的所有字符是否全都不同。 
//
// 示例 1： 
//
// 输入: s = "leetcode"
//输出: false 
// 
//
// 示例 2： 
//
// 输入: s = "abc"
//输出: true
// 
//
// 限制： 
// 
// 0 <= len(s) <= 100 
// 如果你不使用额外的数据结构，会很加分。 
// 
// Related Topics 位运算 哈希表 字符串 排序 👍 208 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isUnique(astr string) bool {
	//---------------------------位运算------------------------------
	// Time: O(n)
	// Space: O(1)
	flag := 0
	for i := 0 ; i < len(astr) ; i++ {
		tem := astr[i] - 'a'
		if flag & (1 << tem) != 0 {
			return false
		}
		flag |= 1 << tem
	}
	return true
	//-------------------------------------------------------------

	//---------------------------数组哈希------------------------------
	/*// Time: O(n)
	// Space: O(1)
	arr := [26]int{}
	for i := 0 ; i < len(astr) ; i++ {
		tem := astr[i] - 'a'
		arr[tem]++
		if arr[tem] > 1 {
			return false
		}
	}
	return true*/
	//-------------------------------------------------------------

	//---------------------------哈希表----------------------------
	/*// Time: O(n)
	// Space: O(n)
	hashMap := map[byte]int{}
	for i := 0 ; i < len(astr) ; i++ {
		hashMap[astr[i]]++
		if hashMap[astr[i]] > 1 {
			return false
		}
	}
	return true*/
	//-------------------------------------------------------------

	//---------------------------排序---------------------------
	/*// Time: O(nlogn)
	// Space: O(nlogn)
	s := []byte(astr)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	for i := 1 ; i < len(s) ; i++ {
		if s[i] == s[i - 1] {
			return false
		}
	}
	return true*/
	//-------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
