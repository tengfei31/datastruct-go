package linklist

import (
	"log"
	"testing"
)

func TestCreateList(t *testing.T) {
	node := CreateNode(5)

	//遍历单向链表
	for node != nil {
		log.Printf("%v", node.Element)
		node = node.Link
	}
}

func TestReverseList(t *testing.T) {
	node := CreateNode(5)
	PrintNode(node)

	log.Println("---------反转后-----------")

	newNode := ReverseList(node)
	PrintNode(newNode)
}
