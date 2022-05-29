//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。 
//
// 示例 1： 
//
// 输入: s1 = "abc", s2 = "bca"
//输出: true 
// 
//
// 示例 2： 
//
// 输入: s1 = "abc", s2 = "bad"
//输出: false
// 
//
// 说明： 
//
// 
// 0 <= len(s1) <= 100 
// 0 <= len(s2) <= 100 
// 
// Related Topics 哈希表 字符串 排序 👍 67 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func CheckPermutation(s1 string, s2 string) bool {
	//----------------------------排序---------------------------------
	/*// Time: O(nlogn)
	// Space: O(logn)
	arr1, arr2 := []byte(s1), []byte(s2)
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	return string(arr1) == string(arr2)*/
	//--------------------------------------------------------------

	//--------------------------数组哈希------------------------------
	// Time: O(n)
	// Space: O(1)
	arr := [26]int{}
	for _, v := range s1 {
		arr[v - 'a']++
	}
	for _, v := range s2 {
		arr[v - 'a']--
	}

	for i := 0 ; i < 26 ; i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
	//--------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
