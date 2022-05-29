//给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <= i < len(arr) / 2，都有 arr[2 * i 
//+ 1] = 2 * arr[2 * i]” 时，返回 true；否则，返回 false。 
//
// 
//
// 示例 1： 
//
// 
//输入：arr = [3,1,3,6]
//输出：false
// 
//
// 示例 2： 
//
// 
//输入：arr = [2,1,2,6]
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：arr = [4,-2,2,-4]
//输出：true
//解释：可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= arr.length <= 3 * 10⁴ 
// arr.length 是偶数 
// -10⁵ <= arr[i] <= 10⁵ 
// 
// Related Topics 贪心 数组 哈希表 排序 👍 99 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func canReorderDoubled(arr []int) bool {
	//----------------------------------哈希表---------------------------------
	// Time: O(nlogn)
	// Space: O(n)
	hashMap := map[int]int{}
	for _, v := range arr {
		hashMap[v]++
	}

	sort.Slice(arr, func(a , b int) bool {return abs(arr[a]) < abs(arr[b])})

	for i := 0 ; i < len(arr) ; i++ {
		if hashMap[2 * arr[i]] < hashMap[arr[i]] {
			return false
		}
		hashMap[2 * arr[i]] -= hashMap[arr[i]]
		hashMap[arr[i]] = 0
	}
	return true
	//------------------------------------------------------------------------
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
