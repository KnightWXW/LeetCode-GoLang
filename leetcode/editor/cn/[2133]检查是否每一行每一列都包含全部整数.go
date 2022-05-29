//对一个大小为 n x n 的矩阵而言，如果其每一行和每一列都包含从 1 到 n 的 全部 整数（含 1 和 n），则认为该矩阵是一个 有效 矩阵。 
//
// 给你一个大小为 n x n 的整数矩阵 matrix ，请你判断矩阵是否为一个有效矩阵：如果是，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：matrix = [[1,2,3],[3,1,2],[2,3,1]]
//输出：true
//解释：在此例中，n = 3 ，每一行和每一列都包含数字 1、2、3 。
//因此，返回 true 。
// 
//
// 示例 2： 
//
// 
//
// 
//输入：matrix = [[1,1,1],[1,2,3],[1,2,3]]
//输出：false
//解释：在此例中，n = 3 ，但第一行和第一列不包含数字 2 和 3 。
//因此，返回 false 。
// 
//
// 
//
// 提示： 
//
// 
// n == matrix.length == matrix[i].length 
// 1 <= n <= 100 
// 1 <= matrix[i][j] <= n 
// 
// Related Topics 数组 哈希表 矩阵 👍 8 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func checkValid(matrix [][]int) bool {
	//-------------------------------hashMap-------------------------------------
	// Time: O(n2)
	// Space: O(n)
	row, col := len(matrix), len(matrix[0])

	for i := 0 ; i < row ; i++ {
		hashMap := map[int]bool{}
		for j := 0 ; j < col ; j++ {
			if hashMap[matrix[i][j]] == true {
				return false
			}
			hashMap[matrix[i][j]] = true
		}
	}

	for j := 0 ; j < col ; j++ {
		hashMap := map[int]bool{}
		for i := 0 ; i < row ; i++ {
			if hashMap[matrix[i][j]] == true {
				return false
			}
			hashMap[matrix[i][j]] = true
		}
	}

	return true
	//---------------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
