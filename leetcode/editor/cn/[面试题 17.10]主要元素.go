//数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1
//) 的解决方案。 
//
// 
//
// 示例 1： 
//
// 
//输入：[1,2,5,9,5,9,5,5,5]
//输出：5 
//
// 示例 2： 
//
// 
//输入：[3,2]
//输出：-1 
//
// 示例 3： 
//
// 
//输入：[2,2,1,1,1,2,2]
//输出：2 
// Related Topics 数组 计数 👍 199 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	//------------------------------摩尔投票---------------------------------
	// Time: O(n)
	// Space: O(1)
	vote, candidate := 0, 0
	for _, v := range nums {
		if vote == 0 {
			candidate = v
			vote = 1
		}else{
			if v == candidate {
				vote++
			}else{
				vote--
			}
		}
	}
	// 判断你是否真的超过一半：
	if vote > 0 {
		cnt := 0
		for _, v := range nums {
			if v == candidate {
				cnt++
			}
		}
		if cnt * 2 > len(nums) {
			return candidate
		}
	}

	return -1
	//----------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
