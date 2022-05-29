//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。 
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。 
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。 
//
// 
//
// 示例 1： 
//
// 
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
// 
//
// 示例 2： 
//
// 
//输入：l1 = [0], l2 = [0]
//输出：[0]
// 
//
// 示例 3： 
//
// 
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
// 
//
// 
//
// 提示： 
//
// 
// 每个链表中的节点数在范围 [1, 100] 内 
// 0 <= Node.val <= 9 
// 题目数据保证列表表示的数字不含前导零 
// 
// Related Topics 递归 链表 数学 👍 7993 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//------------------------------遍历-------------------------------------
	// Time: O(max(n + m))
	// Space: O(max(n + m))
	len1, len2 := digitLen(l1), digitLen(l2)
	tail := l1
	curS, curL, ans := l1, l2, l2
	add, sum := 0, 0
	if len1 > len2 {
		curS, curL, ans = l2, l1, l1
	}
	for curS != nil {
		sum = curS.Val + curL.Val + add
		add = sum / 10
		curL.Val = sum % 10
		tail = curL
		curS, curL = curS.Next, curL.Next
	}
	for curL != nil {
		sum = curL.Val + add
		add = sum / 10
		curL.Val = sum % 10
		tail = curL
		curL = curL.Next
	}
	if add != 0 {
		tail.Next = &ListNode{Val : add}
	}
	return ans
	//--------------------------------------------------------------------
}

func digitLen(l *ListNode) int {
	cnt := 1
	for l != nil {
		l = l.Next
		cnt++
	}
	return cnt
}
//leetcode submit region end(Prohibit modification and deletion)
