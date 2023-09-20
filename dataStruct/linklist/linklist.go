package linklist

import "log"

//链表

type Entry struct {
	Key  int
	Data int
}

type Node struct {
	Element Entry
	Link    *Node
}

// NewNode2 构造新结点
func NewNode2(x Entry) *Node {
	var p *Node = new(Node)
	p.Element = x
	p.Link = nil
	return p
}

// CreateNode 创建一个单向链表
func CreateNode(n int) *Node {
	var node *Node

	for i := n; i >= 1; i-- {
		tmpNode := NewNode2(Entry{Key: i, Data: i})
		if node == nil {
			node = tmpNode
			continue
		}
		tmpNode.Link = node
		node = tmpNode
	}
	return node
}

// ReverseList 反转单向链表
func ReverseList(node *Node) *Node {
	var newNode *Node = nil
	var currNode *Node = node
	for currNode != nil {
		tmpNode := currNode.Link
		currNode.Link = newNode
		newNode = currNode
		currNode = tmpNode
	}
	return newNode
}

// PrintNode 遍历单向链表
func PrintNode(node *Node) {
	//遍历单向链表
	for node != nil {
		log.Printf("%v", node.Element)
		node = node.Link
	}

}
