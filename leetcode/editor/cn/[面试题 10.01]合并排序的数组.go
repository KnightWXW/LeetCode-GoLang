//给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。 
//
// 初始化 A 和 B 的元素数量分别为 m 和 n。 
//
// 示例: 
//
// 输入:
//A = [1,2,3,0,0,0], m = 3
//B = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6] 
//
// 说明: 
//
// 
// A.length == n + m 
// 
// Related Topics 数组 双指针 排序 👍 141 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func merge(A []int, m int, B []int, n int)  {
	//-----------------------------自后向前遍历-------------------------------
	// Time: O(n + m)
	// Space: O(1)
	indexA, indexB := m - 1, n - 1
	index := len(A) - 1
	for indexA >= 0 && indexB >= 0 {
		if A[indexA] < B[indexB] {
			A[index] = B[indexB]
			indexB--
		}else{
			A[index] = A[indexA]
			indexA--
		}
		index--
	}
	for indexB >= 0 {
		A[index] = B[indexB]
		indexB--
		index--
	}
	//-------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
