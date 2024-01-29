package set

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLinkedHashSet_Size(t *testing.T) {
	set := NewLinkedHashSet[string]()
	if set.Size() != 0 {
		t.Error("set is not empty")
	}
	set.Add("aa")
	set.Add("bb")
	if set.Size() != 2 {
		t.Error("set is not empty")
	}
}

func TestLinkedHashSet_IsEmpty(t *testing.T) {
	set := NewLinkedHashSet[string]()
	if !set.IsEmpty() {
		t.Error("set is not empty")
	}
	set.Add("aa")
	set.Add("bb")
	if set.IsEmpty() {
		t.Error("set is not empty")
	}
}

func TestLinkedHashSet_Contains(t *testing.T) {
	set := NewLinkedHashSet[string]("aaa")
	if !set.Contains("aaa") {
		t.Error("aaa is not Contains")
	}
	if set.Contains("aa") {
		t.Error("aaa is not Contains")
	}
}

func TestLinkedHashSet_ToSlice(t *testing.T) {
	set := NewLinkedHashSet[string]("aaa")
	for i := 0; i < 10; i++ {
		set.Add(fmt.Sprintf("%v", i))
	}
	t.Log(set.ToSlice())
}

func TestLinkedHashSet_Add(t *testing.T) {
	set := NewLinkedHashSet[string]()
	aaAdd := set.Add("aa")
	if !aaAdd {
		t.Error("add return is false")
	}
	aaAddAgain := set.Add("aa")
	if aaAddAgain {
		t.Error("aa again return is true")
	}
	if !set.Contains("aa") {
		t.Error("add err")
	}

	if set.Contains("cc") {
		t.Error("add err")
	}
	t.Log(set.ToSlice())
}

func TestLinkedHashSet_Remove(t *testing.T) {
	set := NewLinkedHashSet[string]("aa")
	rmFlag := set.Remove("aa")
	if !rmFlag {
		t.Error("rm return is false")
	}
	if set.Remove("aa") {
		t.Error("rm return is true")
	}
	if !set.IsEmpty() {
		t.Error("set is not empty")
	}
}

func TestLinkedHashSet_ContainsAll(t *testing.T) {
	set := NewLinkedHashSet[string]("aa", "bb")
	if !set.ContainsAll([]string{"aa", "bb"}) {
		t.Error("ContainsAll is false")
	}
	if set.ContainsAll([]string{"aa", "cc"}) {
		t.Error("ContainsAll is true")
	}
}

func TestLinkedHashSet_AddAll(t *testing.T) {
	set := NewLinkedHashSet[string]()
	if !set.AddAll([]string{"aa", "bb"}) {
		t.Error("addAll return is false")
	}
	if !set.AddAll([]string{"aa", "cc"}) {
		t.Error("addAll return is false")
	}
	if set.AddAll([]string{"aa"}) {
		t.Error("addAll return is true")
	}
	if set.Size() != 3 {
		t.Error("size is not 2")
	}
	t.Log(set.ToSlice())
}

func TestLinkedHashSet_RemoveAll(t *testing.T) {
	set := NewLinkedHashSet[string]("aa", "bb")
	if !set.RemoveAll([]string{"aa"}) {
		t.Error("addAll return is false")
	}
	if set.RemoveAll([]string{"cc"}) {
		t.Error("addAll return is true")
	}
	if set.Size() != 1 {
		t.Error("size is not 2")
	}
	t.Log(set.ToSlice())
}

func TestLinkedHashSet_Clear(t *testing.T) {
	set := NewLinkedHashSet[string]("aa", "bb")
	set.Clear()
	if !set.IsEmpty() {
		t.Error("set is not empty")
	}
}

func TestLinkedHashSet_ForEach(t *testing.T) {
	set := NewLinkedHashSet[string]("aa", "bb", "cc")
	set.ForEach(func(s string) {
		t.Log(s)
	})
}

func TestLinkedHashSet_MarshalJSON(t *testing.T) {
	set := NewLinkedHashSet[string]("aa", "bb", "cc")
	bytes, err := json.Marshal(set)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bytes))
}

func TestLinkedHashSet_UnmarshalJSON(t *testing.T) {
	jsonStr := "[\"aa\", \"bb\"]"
	set := NewLinkedHashSet[string]("cc")
	err := json.Unmarshal([]byte(jsonStr), set)
	if err != nil {
		t.Error(err)
	}
	if set.Contains("cc") {
		t.Error("element not remove")
	}
	t.Log(set.ToSlice())
}
