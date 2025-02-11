package fastandslowpointers

/*
Given the head of a Singly LinkedList, write a method to return the middle node of the LinkedList.

If the total number of nodes in the LinkedList is even, return the second middle node.
*/

type Solution struct{}

func (s *Solution) findMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head
	slow := head

	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
