/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// an additional func gcd
func gcd(a, b int) int {
	// base case
	if b == 0 {
		return a
	}

	// recursive case
	return gcd(b, a%b)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		var divisor int = gcd(current.Val, current.Next.Val)
		newNode := &ListNode{
			Val:  divisor,
			Next: nil,
		}

		newNode.Next = current.Next
		current.Next = newNode
        current = current.Next.Next
	}
	return head
}