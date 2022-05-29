//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 
//
// 示例 1： 
//
// 
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
// 
//
// 示例 2： 
//
// 
//输入：l1 = [], l2 = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：l1 = [], l2 = [0]
//输出：[0]
// 
//
// 
//
// 提示： 
//
// 
// 两个链表的节点数目范围是 [0, 50] 
// -100 <= Node.val <= 100 
// l1 和 l2 均按 非递减顺序 排列 
// 
// Related Topics 递归 链表 👍 2288 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//-------------------------------递归-------------------------------
	/*// Time: O(m + n)
	// Space: O(m + n)
	if list1 == nil{
		return list2
	}else if list2 == nil {
		return list1
	}else if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}else{
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}*/
	//--------------------------------------------------------------------

	//--------------------------------迭代---------------------------------
	// Time: O(m + n)
	// Space: O(1)
	ans := &ListNode{}
	pre := ans
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
		}else{
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}

	if list1 != nil {
		pre.Next = list1
	}
	if list2 != nil {
		pre.Next = list2
	}
	return ans.Next
	//--------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
