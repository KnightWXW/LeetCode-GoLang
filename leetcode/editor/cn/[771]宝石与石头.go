//给你一个字符串 jewels 代表石头中宝石的类型，另有一个字符串 stones 代表你拥有的石头。 stones 中每个字符代表了一种你拥有的石头的类型，
//你想知道你拥有的石头中有多少是宝石。 
//
// 字母区分大小写，因此 "a" 和 "A" 是不同类型的石头。 
//
// 
//
// 示例 1： 
//
// 
//输入：jewels = "aA", stones = "aAAbbbb"
//输出：3
// 
//
// 示例 2： 
//
// 
//输入：jewels = "z", stones = "ZZ"
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= jewels.length, stones.length <= 50 
// jewels 和 stones 仅由英文字母组成 
// jewels 中的所有字符都是 唯一的 
// 
// Related Topics 哈希表 字符串 👍 694 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func numJewelsInStones(jewels string, stones string) int {
	//--------------------------------模拟-----------------------------------
	// Time: O(m * n)
	// Space: O(1)
	cnt := 0
	for i := 0 ; i < len(jewels) ; i++ {
		for j := 0 ; j < len(stones) ; j++ {
			if jewels[i] == stones[j] {
				cnt++
			}
		}
	}
	return cnt
	//----------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
