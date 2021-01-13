package stack

import (
	"log"
)

const MaxSize = 50

type BOOL int

type T int

//Stack 结构
type Stack struct {
	Top, MaxStack int
	Elements      [MaxSize]T
}

//CreateStack 创建一个空栈
func CreateStack(maxSize int) *Stack {
	var stack *Stack = new(Stack)
	stack.Top = -1
	stack.MaxStack = maxSize
	return stack
}

//IsEmpty 栈是否为空
func (st *Stack) IsEmpty() bool {
	return st.Top < 0
}

//IsFull 栈是否满了
func (st *Stack) IsFull() bool {
	return st.Top >= st.MaxStack-1
}

//Push 推到栈顶
func (st *Stack) Push(x T) {
	if st.IsFull() == true {
		log.Fatalf("Overflow")
	} else {
		st.Top++
		st.Elements[st.Top] = x
	}
}

//Pop 删除栈顶元素
func (st *Stack) Pop() {
	if st.IsEmpty() == true {
		log.Fatalf("Underflow")
	} else {
		st.Top--
	}
}

//StackTop 返回栈顶元素
func (st *Stack) StackTop() *T {
	if st.IsEmpty() == true {
		log.Printf("Underflow")
		return nil
	} else {
		var tmpT *T
		*tmpT = st.Elements[st.Top]
		return tmpT
	}
}
