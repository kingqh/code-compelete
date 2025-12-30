// Reverse Linked List
// 反转一个单链表，返回反转后的头结点。
//
// 示例：
// 输入：1 -> 2 -> 3 -> 4 -> 5 -> nil
// 输出：5 -> 4 -> 3 -> 2 -> 1 -> nil
//
// 要求：
// - 实现函数 `reverseList(head *ListNode) *ListNode`
// - 时间复杂度期望 O(n)，空间复杂度 O(1)（迭代法）
// - 也可实现递归版本，递归版本空间复杂度为 O(n)（递归栈）
// - 注意处理空链表和单节点链表
//
// 链表节点定义示例（Go）：
//
// type ListNode struct {
//     Val  int
//     Next *ListNode
// }
//
// 解题思路（迭代）：
// 使用三个指针遍历链表并反转指向：`prev`（已反转链表的头）、`curr`（当前节点）、`next`（保存 curr.Next）。
// 每步操作：
// 1. next = curr.Next
// 2. curr.Next = prev
// 3. prev = curr
// 4. curr = next
// 遍历结束后，`prev` 为新的头结点，返回 `prev`。
//
// 解题思路（递归）：
// 递归反转子链表，直到到达末尾节点作为新头，然后回溯过程中把当前节点接到子链表末尾并断开旧连接。
// 递归伪代码：
// if head == nil || head.Next == nil return head
// newHead := reverseList(head.Next)
// head.Next.Next = head
// head.Next = nil
// return newHead

package main
import (
    "fmt"
    )

type LinkNode struct {
    Val int
    Next *LinkNode
}

func main() {
    head := &LinkNode{Val: 1}
    head.Next = &LinkNode{Val: 2}
    head.Next.Next = &LinkNode{Val: 3}
    head.Next.Next.Next = &LinkNode{Val: 4}
    head.Next.Next.Next.Next = &LinkNode{Val: 5}
    printLinkedList(head)
    head = reverseLinkedList(head)
    printLinkedList(head)
}

func reverseLinkedList(linklist *LinkNode) *LinkNode {
   var prev *LinkNode
   curr := linklist
   for curr != nil {
     next := curr.Next
     curr.Next = prev
     prev = curr 
     curr = next
   }
   return prev
}

func printLinkedList(head *LinkNode) {
    for head != nil {
        fmt.Printf("%v ", head.Val)
        head = head.Next
    }
}
