package link

import "fmt"

// 连接节点
type DoubleLinkNode struct {
	Value interface{}     // 数据
	Pre   *DoubleLinkNode // 上一个节点
	Next  *DoubleLinkNode // 下一个节点
}

func ReverseDoubleLinkNode(head *DoubleLinkNode) *DoubleLinkNode {
	var pre *DoubleLinkNode
	var next *DoubleLinkNode

	for head != nil {
		// 保留下次要执行的节点
		next = head.Next

		head.Next = pre
		head.Pre = next
		// 将head加入到pre中
		pre = head

		// 下次执行的节点赋值给head
		head = next
	}
	return pre
}

func RangeDouble(max int) *DoubleLinkNode {
	head := &DoubleLinkNode{}
	cur := head
	var pre *DoubleLinkNode
	for i := 0; i < max; i++ {
		if cur == nil {
			cur = &DoubleLinkNode{}
		}
		cur.Value = i
		if pre != nil {
			pre.Next = cur
		}
		cur.Pre = pre
		pre = cur
		cur = cur.Next
	}
	return head
}

// 打印链表
func PrintDoubleLink(head *DoubleLinkNode) {
	if head != nil {
		fmt.Print(head.Value, " ")
		PrintDoubleLink(head.Next)
	}
}
