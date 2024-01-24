package cmap

import "testing"

func TestNewLinkedHashMap(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	if !linkedHashMap.IsEmpty() {
		t.Error("linkedHashMap is not empty")
	}
}
