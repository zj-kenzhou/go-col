package set

import "testing"

func TestHashSetNew(t *testing.T) {
	set := NewSet[string]("aa", "bb", "aa")
	if set.IsEmpty() {
		t.Error("set is empty")
	}
	if set.Size() != 2 {
		t.Error("set size is not 2")
	}
	set.Add("cc")
	if set.Size() != 3 {
		t.Error("set size is not 3")
	}
	t.Log(set.ToSlice())
	intSet := NewSet[int](1, 2)
	t.Log(intSet.ToSlice())
}
func TestNewSyncSet(t *testing.T) {
	set := NewSyncSet[string]("aa", "bb", "aa")
	if set.IsEmpty() {
		t.Error("set is empty")
	}
	if set.Size() != 2 {
		t.Error("set size is not 2")
	}
	set.Add("cc")
	if set.Size() != 3 {
		t.Error("set size is not 3")
	}
	t.Log(set.ToSlice())
	intSet := NewSet[int](1, 2)
	t.Log(intSet.ToSlice())
}

func TestNewLinkedHashSet(t *testing.T) {
	set := NewLinkedHashSet[string]("aa", "bb", "aa")
	if set.IsEmpty() {
		t.Error("set is empty")
	}
	if set.Size() != 2 {
		t.Error("set size is not 2")
	}
	set.Add("cc")
	if set.Size() != 3 {
		t.Error("set size is not 3")
	}
	t.Log(set.ToSlice())
	intSet := NewSet[int](1, 2)
	t.Log(intSet.ToSlice())
}
