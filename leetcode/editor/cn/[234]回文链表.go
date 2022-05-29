//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。 
//
// 
//
// 示例 1： 
//
// 
//输入：head = [1,2,2,1]
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：head = [1,2]
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 链表中节点数目在范围[1, 10⁵] 内 
// 0 <= Node.val <= 9 
// 
//
// 
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？ 
// Related Topics 栈 递归 链表 双指针 👍 1318 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	//-------------------------------数组-----------------------------------
	/*// Time: O(n)
	// Space: O(n)
	arr := []int{}
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	for i := 0 ; i < len(arr) / 2 ; i++ {
		if arr[i] != arr[len(arr) - i - 1] {
			return false
		}
	}
	return true*/
	//----------------------------------------------------------------------

	//------------------------------反转链表----------------------------------
	// Time: O(n)
	// Space: O(1)
	mid := searchMid(head)
	reverseNode := reverse(mid.Next)
	node1, node2 := head, reverseNode
	for node1 != nil && node2 != nil {
		if node1.Val != node2.Val {
			return false
		}
		node1, node2 = node1.Next, node2.Next
	}
	return true
	//----------------------------------------------------------------------
}

// 反转链表：
func reverse(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		tem := cur.Next
		cur.Next = pre
		pre = cur
		cur = tem
	}
	return pre
}

// 寻找中间节点：
func searchMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
//leetcode submit region end(Prohibit modification and deletion)
