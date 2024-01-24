package list

import "testing"

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[string]("aa")
	if list.IsEmpty() {
		t.Error("list is empty")
	}
	if list.Size() != 1 {
		t.Error("list size is not 1")
	}
}
