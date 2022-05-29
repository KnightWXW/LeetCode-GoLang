//给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
// 
//
// 示例 2： 
//
// 
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//解释：[4,9] 也是可通过的
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums1.length, nums2.length <= 1000 
// 0 <= nums1[i], nums2[i] <= 1000 
// 
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 514 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	//---------------------------hashMap---------------------------------
	/*// Time: O(max(n, m))
	// Space: O(n + m)
	hashMap1, hashMap2 := map[int]int{}, map[int]int{}
	for _, v := range nums1 {
		hashMap1[v]++
	}
	for _, v := range nums2 {
		hashMap2[v]++
	}
	ans := []int{}
	if len(hashMap1) < len(hashMap2) {
		hashMap1,hashMap2 = hashMap2, hashMap1
	}
	for k1, _ := range hashMap1 {
		if _, ok := hashMap2[k1]; ok {
			ans = append(ans, k1)
		}
	}
	return ans*/
	//------------------------------------------------------------------

	//--------------------------排序 + 双指针------------------------------
	// Time: O(mlogm + nlogn)
	// Sapce: O(logm + logn)
	sort.Ints(nums1)
	sort.Ints(nums2)
	index1, index2 := 0, 0
	ans := []int{}
	for index1 < len(nums1) && index2 < len(nums2) {
		tem1, tem2 := nums1[index1], nums2[index2]
		if tem1 == tem2 {
			if len(ans) == 0 || tem1 != ans[len(ans) - 1]{
				ans	= append(ans, tem1)
			}
			index1++
			index2++
		}else if tem1 < tem2 {
			index1++
		}else{
			index2++
		}
	}
	return ans
	//------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
