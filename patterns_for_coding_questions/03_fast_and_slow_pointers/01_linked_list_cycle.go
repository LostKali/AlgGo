package fastandslowpointers

/*
Given the head of a Singly LinkedList, write a function to determine if the LinkedList
has a cycle in it or not.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct{}

func (s *Solution) hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head.Next
	fast := head.Next.Next

	for fast != nil && fast != slow {
		if fast.Next == nil {
			fast = nil
		} else {
			fast = fast.Next.Next
		}
		slow = slow.Next
	}

	return fast != nil && fast == slow
}

func main() {
	solution := Solution{}
	res := solution.hasCycle(createLinkedList([]int{1, 2, 3, 4, 5, 6}))
	println(res)
}

func createLinkedList(nums []int) *ListNode {
	var head *ListNode
	var node *ListNode

	for _, n := range nums {
		if head == nil {
			head = &ListNode{
				Val: n,
			}
			node = head
		} else {
			node.Next = &ListNode{
				Val: n,
			}
			node = node.Next
		}
	}

	return head
}
