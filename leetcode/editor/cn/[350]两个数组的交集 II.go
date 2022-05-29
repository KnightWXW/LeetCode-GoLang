//给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现
//次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
// 
//
// 示例 2: 
//
// 
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums1.length, nums2.length <= 1000 
// 0 <= nums1[i], nums2[i] <= 1000 
// 
//
// 
//
// 进阶： 
//
// 
// 如果给定的数组已经排好序呢？你将如何优化你的算法？ 
// 如果 nums1 的大小比 nums2 小，哪种方法更优？ 
// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？ 
// 
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 704 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	//---------------------------哈希表1-----------------------------
	/*// Time: O(m * n * k)
	// Space: O(m + n)
	hashSet1 := map[int]int{}
	hashSet2 := map[int]int{}
	for _, v := range nums1 {
		hashSet1[v]++
	}
	for _, v := range nums2 {
		hashSet2[v]++
	}
	ans := []int{}
	for k1, v1 := range hashSet1 {
		for k2, v2 := range hashSet2 {
			if k1 == k2 {
				tem := min(v1 , v2)
				for tem > 0 {
					ans = append(ans , k1)
					tem--
				}
			}
		}
	}
	return ans*/
	//-------------------------------------------------------------

	//---------------------------哈希表2-----------------------------
	/*// Time: O(m + n)
	// Space: O(min(m,n))
	if len(nums1) < len(nums2) {
		return intersect(nums2, nums1)
	}
	hashSet := map[int]int{}
	for _, v := range nums1{
		hashSet[v]++
	}
	ans := []int{}
	for _, v := range nums2 {
		if hashSet[v] > 0 {
			ans = append(ans , v)
			hashSet[v]--
		}
	}
	return ans*/
	//-------------------------------------------------------------

	//------------------------排序 + 双指针--------------------------
	// Time: O(mlogm + nlogn)
	// Space: O(logm + logn)
	sort.Ints(nums1)
	sort.Ints(nums2)

	index1 , index2 := 0, 0
	ans := []int{}
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] == nums2[index2] {
			ans = append(ans, nums1[index1])
			index1++
			index2++
		}else if nums1[index1] < nums2[index2] {
			index1++
		}else{
			index2++
		}
	}
	return ans
	//-------------------------------------------------------------
}

func min(a , b int) int {
	if a < b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
