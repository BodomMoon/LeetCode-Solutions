// Time:  O(m + n)
// Space: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var curA, curB, tailA, tailB *ListNode
	// default intilize to nil

	for curA, curB = headA, headB; curA != nil && curB != nil; {
		if curA == curB {
			return curA
		}

		if curA.Next != nil {
			curA = curA.Next
		} else if tailA == nil {
			tailA = curA
			curA = headB
		} else {
			break
		}

		if curB.Next != nil {
			curB = curB.Next
		} else if tailB == nil {
			tailB = curB
			curB = headA
		} else {
			break
		}
	}
	return nil
}