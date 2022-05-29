//稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。 
//
// 示例1: 
//
//  输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""],
// s = "ta"
// 输出：-1
// 说明: 不存在返回-1。
// 
//
// 示例2: 
//
//  输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], 
//s = "ball"
// 输出：4
// 
//
// 提示: 
//
// 
// words的长度在[1, 1000000]之间 
// 
// Related Topics 数组 字符串 二分查找 👍 69 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findString(words []string, s string) int {
	//-----------------------------遍历---------------------------------
	/*// Time: O(n)
	// Space: O(1)
	for i := 0 ; i < len(words) ; i++ {
		if s == words[i] {
			return i
		}
	}
	return -1*/
	//------------------------------------------------------------------

	//-----------------------------二分查找--------------------------------
	// Time: O(logn)
	// Space: O(1)
	left, right := 0, len(words) - 1
	for left <= right {
		for left <= right && words[left] == "" {
			left++
		}
		for left <= right && words[right] == "" {
			right--
		}
		mid := left + (right - left) >> 1
		for  left <= mid && words[mid] == "" {
			mid--
		}
		if s < words[mid] {
			right = mid - 1
		}else if s > words[mid] {
			left = mid + 1
		}else{
			return mid
		}
	}
	return -1
	//------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
