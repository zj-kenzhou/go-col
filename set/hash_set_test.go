package set

import "testing"

func TestNewSet(t *testing.T) {
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
}
