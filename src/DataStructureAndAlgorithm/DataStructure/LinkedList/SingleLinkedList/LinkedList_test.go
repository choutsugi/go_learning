package SingleLinkedList

import "testing"

func TestLinkNode_Add(t *testing.T) {
	head := NewLinkNode(10)
	for head.Next != nil {
		t.Logf("%s", head.Payload)
		head = head.Next
	}
}
