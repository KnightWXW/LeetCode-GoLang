//一个 平方和三元组 (a,b,c) 指的是满足 a² + b² = c² 的 整数 三元组 a，b 和 c 。 
//
// 给你一个整数 n ，请你返回满足 1 <= a, b, c <= n 的 平方和三元组 的数目。 
//
// 
//
// 示例 1： 
//
// 输入：n = 5
//输出：2
//解释：平方和三元组为 (3,4,5) 和 (4,3,5) 。
// 
//
// 示例 2： 
//
// 输入：n = 10
//输出：4
//解释：平方和三元组为 (3,4,5)，(4,3,5)，(6,8,10) 和 (8,6,10) 。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 250 
// 
// Related Topics 数学 枚举 👍 11 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func countTriples(n int) int {
	//--------------------------暴力遍历-------------------------------
	/*// Time: O(n ^ 3)
	// Space: O(1)
	ans := 0
	for a := 1; a <= n; a++ {
		for b := 1 ; b <= n ; b++ {
			for c := 1 ; c <= n ; c++ {
				if a * a + b * b ==	c * c {
					ans++
				}
			}
		}
	}
	return ans*/
	//----------------------------------------------------------------

	//------------------------------哈希表-----------------------------
	// Time: O(n ^ 2)
	// Space: O(n)
	hashMap := map[int]bool{}
	for i := 1 ; i <= n ; i++ {
		hashMap[i * i] = true
	}
	ans := 0
	for a := 1 ; a <= n ; a++ {
		for b := 1 ; b <= n ; b++ {
			if hashMap[a * a + b * b] == true {
				ans++
			}
		}
	}
	return ans
	//-----------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
