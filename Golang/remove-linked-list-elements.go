// Time:  O(n)
// Space: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeElements(head *ListNode, val int) *ListNode {
    for ptrToNext := &head; *ptrToNext != nil; {
        if (*ptrToNext).Val == val {
            *ptrToNext = (*ptrToNext).Next
            continue
        } else (*ptrToNext) != nil{
            ptrToNext = &((*ptrToNext).Next)  
        }
    }
    return head
}