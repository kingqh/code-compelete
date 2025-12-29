/*

反转链表 (Reverse Linked List)

题目描述：反转一个单链表。

示例：输入: 1->2->3->4->5->NULL，输出: 5->4->3->2->1->NULL。

解题思路：迭代或递归。迭代时，需要维护前一个节点、当前节点和下一个节点，将当前节点的Next指向前一个节点，然后移动指针。


*/

package main

import (
    "fmt"
)

type LinkedList struct {
    Next *LinkedList
    Value int
}

func main() {

    nums := []int{1, 2, 3, 4, 5}
    llist := toLinkedList(nums)
    printLinkedList(llist)

    reverseLlist := reverseLinkedList(llist) 
    printLinkedList(reverseLlist)

}

func reverseLinkedList(llist *LinkedList) *LinkedList {
   head := llist
   tail := llist
   for tail.Next != nil {
        tmp := tail.Next
        tail.Next = tail.Next.Next
        tmp.Next = head
        head = tmp
   }

   return head
}

func toLinkedList(nums []int) *LinkedList {
    if len(nums) <= 0 {
        return nil
    }
    head := &LinkedList{
        Value: nums[0],
    }
    tail := head
    for i := 1; i < len(nums); i++ {
        next := &LinkedList{
            Value: nums[i],
        }
        tail.Next = next
        tail = next
    }
    return head
}

func printLinkedList(llist *LinkedList) {
    str := ""
    for llist.Next != nil {
        str = fmt.Sprintf("%s%v -> ", str, llist.Value)
        llist = llist.Next
    }

    str = fmt.Sprintf("%s%v -> NULL", str, llist.Value)
    fmt.Printf("The linked list is %v", str)
}
