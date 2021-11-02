package SingleLinkedList

type Item interface {
}

type LinkNode struct {
	Payload Item
	Next    *LinkNode
}

type LinkNoder interface {
	Add(payload Item)
	Delete(index int) (data Item)
	GetLength() (length int)
	Insert(index int, payload Item)
	Search(payload Item) (index int)
	GetAll() (datas []Item)
	Reverse() *LinkNode
	RecursionReverse() *LinkNode
}

func (head *LinkNode) Add(payload Item) {
	iterator := head

	// 尾插法
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = &LinkNode{
		Payload: payload,
		Next:    nil,
	}

	// 头插法
	//iterator.Next = &LinkNode{
	//	Payload: payload,
	//	Next:    iterator.Next,
	//}
}

func (head *LinkNode) Delete(index int) (data Item) {
	iterator := head

	if index < 0 || index > iterator.GetLength() {
		return nil
	}

	for i := 0; i < index; i++ {
		iterator = iterator.Next
	}

	iterator.Next = iterator.Next.Next
	data = iterator.Next.Payload
	return
}

func (head *LinkNode) GetLength() (length int) {
	iterator := head

	for iterator.Next != nil {
		length++
		iterator = iterator.Next
	}
	return
}

func (head *LinkNode) Insert(index int, payload Item) {
	iterator := head
	length := head.GetLength()
	if index < 0 || index > length {
		return
	}

	for i := 0; i < index; i++ {
		iterator = iterator.Next
	}

	iterator.Next = &LinkNode{
		Payload: payload,
		Next:    iterator.Next,
	}
}

func (head *LinkNode) Search(payload Item) (index int) {
	iterator := head
	length := head.GetLength()
	for iterator.Next != nil {
		if iterator.Payload == payload {
			return index
		} else {
			index++
			iterator = iterator.Next
			if index > length-1 {
				break
			}
			continue
		}
	}

	if iterator.Payload == payload {
		return index
	}
	return -1
}

func (head *LinkNode) GetAll() (datas []Item) {
	datas = make([]Item, 0, head.GetLength())
	iterator := head
	for iterator.Next != nil {
		datas = append(datas, iterator.Payload)
		iterator = iterator.Next
	}

	datas = append(datas, iterator.Payload)
	return datas
}

func (head *LinkNode) Reverse() *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}

	reverseHead := head
	head = head.Next
	reverseHead.Next = nil
	p := head.Next
	for head != nil {
		head.Next = reverseHead
		reverseHead = head
		head = p
		if p != nil {
			p = p.Next
		}
	}
	return reverseHead
}

func (head *LinkNode) RecursionReverse() *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}

	second := head.Next
	newHead := second.RecursionReverse()
	second.Next = head
	head.Next = nil
	return newHead
}

func NewLinkNode(length int) *LinkNode {
	if length <= 0 {
		return nil
	}

	head := &LinkNode{
		Payload: nil,
		Next:    nil,
	}

	iterator := head
	for i := 0; i < length; i++ {
		iterator.Next = &LinkNode{
			Payload: i,
			Next:    nil,
		}

		iterator = iterator.Next
	}
	return head
}
