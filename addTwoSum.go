/**
- Task: You are given two non-empty linked lists representing two non-negative integers.
- The digits are stored in reverse order, and each of their nodes contains a single digit.
- Add the two numbers and return the sum as a linked list.
- You may assume the two numbers do not contain any leading zero, except the number 0 itself.

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

dummyHead := new(ListNode) // nil struct -- will point to .Next as the first element of the asnwer
curr := dummyHead // diverting the curr pointer to a nil Head
carryAm := 0 // carry is zero

for l1 != nil || l2 != nil { // checking if we reached the end of both lists
sum := carryAm // adding carry amount to a sum

if l1 != nil {
sum += l1.Val
l1 = l1.Next
}

if l2 != nil {
sum += l2.Val
l2 = l2.Next
}

carryAm = sum / 10 // counting a carry amount for the next step
curr.Next = &ListNode{Val: sum % 10} // counting the value of a digit
curr = curr.Next // moving pointer to the next digit
}

if carryAm > 0 {
curr.Next = &ListNode{Val: carryAm} // if by the time we reached the end of both numbers the carry amount is still not zero, we put it in the next node by itself -- overflow with 1
}

return dummyHead.Next
}
