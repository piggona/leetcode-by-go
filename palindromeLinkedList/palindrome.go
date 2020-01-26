package palindromeLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev := slow
	cur := slow.Next
	next := cur.Next
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	start2 := prev
	start1 := head
	for start2 != nil {
		if start1.Val != start2.Val {
			return false
		}
		start1 = start1.Next
		start2 = start2.Next
	}
	return true
}
