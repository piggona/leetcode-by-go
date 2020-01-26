package intersectionOfTwoLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNodeMap(headA, headB *ListNode) *ListNode {
	cur := headA
	bitmap := make(map[*ListNode]bool)
	for cur != nil {
		bitmap[cur] = true
		cur = cur.Next
	}
	curB := headB
	for curB != nil {
		if bitmap[curB] {
			return curB
		}
		curB = curB.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	A, B := false, false
	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}
		curA, curB = curA.Next, curB.Next
		if curA == nil && !A {
			curA = headB
			A = true
		}
		if curB == nil && !B {
			curB = headA
			B = true
		}
	}
	return nil
}
