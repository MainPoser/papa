package main

import (
	"fmt"
	"strings"
)

type funcType func(string) bool //代表筛选逻辑函数，可以按需实现

func filter(a []string, f funcType) []string {
	for i := range a {
		if f(a[i]) {
			return filter(append(a[:i], a[i+1:]...), f)
		}
	}
	return a
}

// 连接节点
type LinkNode struct {
	Value interface{} // 数据
	Next  *LinkNode   // 下一个节点
}

// 创建链表
func (node *LinkNode) Range(max int) {
	cur := node
	for i := 0; i < max; i++ {
		cur.Next = &LinkNode{}
		cur.Next.Value = i
		cur = cur.Next
	}
}

// 打印链表
func (node *LinkNode) Print() {
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Value, " ")
	}
	fmt.Println()
}

// 打印链表
func (node *LinkNode) PrintReverse() {

	// 定义递归函数
	var reversePrint func(*LinkNode)

	// 实现递归函数
	reversePrint = func(linkNode *LinkNode) {
		if linkNode == nil {
			return
		}
		reversePrint(linkNode.Next)
		fmt.Print(linkNode.Value, " ")
	}

	// 调用递归函数
	reversePrint(node)
	fmt.Println()
}

// 逆序
func (node *LinkNode) Reverse() {
	if node == nil || node.Next == nil {
		return
	}
	var reverseNode *LinkNode // 反转后的节点
	var runNode *LinkNode     // 负责往前跑的遍历节点
	curNode := node.Next      // 当前节点

	for curNode != nil { // 判断是不是跑到最后了
		// 负责跑的节点一次往前跑一步
		runNode = curNode.Next

		// 让当前节点的下一个节点为逆序节点，将当前节点加入到逆序链表的头部
		curNode.Next = reverseNode

		// 让逆序节点为当前节点，记住整个逆序后的链表
		reverseNode = curNode

		// 更新当前节点为奔跑节点，方便奔跑节点继续奔跑
		curNode = runNode
	}

	// pre 记录了逆序后的节点
	node.Next = reverseNode
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	// 字符串倒序
	str := "192.168.19.100,2022::19:100"
	fmt.Println(strings.Join(reverse(strings.Split(str, ",")), ","))
	// 链表倒序
	ln := &LinkNode{}
	ln.Range(10)
	ln.Reverse()
	ln.Print()
	// 过滤集合
	newStr := filter([]string{"8", "2", "1", "4", "5"}, func(s string) bool {
		return s == "1"
	})
	fmt.Printf("newStr: %v\n", newStr)

}
