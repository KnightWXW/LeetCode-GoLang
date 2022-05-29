//给你一个整数数组 arr，请你检查是否存在两个整数 N 和 M，满足 N 是 M 的两倍（即，N = 2 * M）。 
//
// 更正式地，检查是否存在两个下标 i 和 j 满足： 
//
// 
// i != j 
// 0 <= i, j < arr.length 
// arr[i] == 2 * arr[j] 
// 
//
// 
//
// 示例 1： 
//
// 输入：arr = [10,2,5,3]
//输出：true
//解释：N = 10 是 M = 5 的两倍，即 10 = 2 * 5 。
// 
//
// 示例 2： 
//
// 输入：arr = [7,1,14,11]
//输出：true
//解释：N = 14 是 M = 7 的两倍，即 14 = 2 * 7 。
// 
//
// 示例 3： 
//
// 输入：arr = [3,1,7,11]
//输出：false
//解释：在该情况下不存在 N 和 M 满足 N = 2 * M 。
// 
//
// 
//
// 提示： 
//
// 
// 2 <= arr.length <= 500 
// -10^3 <= arr[i] <= 10^3 
// 
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 48 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func checkIfExist(arr []int) bool {
	//--------------------------暴力循环-----------------------------
	/*// Time: O(n2)
	// Space: O(1)
	for i := 0 ; i < len(arr) ; i++ {
		for j := 0 ; j < len(arr) ; j++ {
			if arr[i] == 2 * arr[j] && i != j {
				return true
			}
		}
	}
	return false*/
	//-------------------------------------------------------------

	//-----------------------------哈希表---------------------------
	// Time: O(n)
	// Space: O(n)
	hashSet := map[int]bool{}
	for _, v := range arr {
		if hashSet[2 * v] == true {
			return true
		}
		if v & 1 == 0 && hashSet[v / 2] == true {
			return true
		}
		hashSet[v] = true
	}
	return false
	//------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
