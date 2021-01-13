package linklist

//链表

type Entry struct {
	Key  int
	Data int
}

type T Entry

type Node struct {
	Element T
	Link    *Node
}

//NewNode2 构造新结点
func NewNode2(x T) *Node {
	var p *Node = new(Node)
	p.Element = x
	p.Link = nil
	return p
}
