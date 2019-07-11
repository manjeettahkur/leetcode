package leetcode

//Node element
type Node struct {
	Val  int
	Next *Node
}

// ListNodes is very good field
type ListNodes struct {
	Head *Node
}

func (ll *ListNodes) AddToHead(data int) {
	ll.Head = &Node{Val: data, Next: ll.Head}
}



func AddTwoNumbers(l1 *Node, l2 *Node) *Node {
	resPre := &Node{}
	cur := resPre
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &Node{Val: sum % 10}
		cur = cur.Next
	}
	return resPre.Next
}
