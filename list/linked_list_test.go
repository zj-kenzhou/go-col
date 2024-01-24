package list

import (
	"encoding/json"
	"testing"
)

func TestLinkedList_Size(t *testing.T) {
	list := NewLinkedList[string]("aa")
	if list.Size() != 1 {
		t.Error("list size is not 1")
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	list := NewLinkedList[string]("aa")
	if list.IsEmpty() {
		t.Error("list is empty")
	}
	list.Remove("aa")
	if !list.IsEmpty() {
		t.Error("list is not empty")
	}
}

func TestLinkedList_Contains(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb")
	if list.Contains("a") {
		t.Error("list contains a")
	}
	if !list.Contains("aa") {
		t.Error("list not contains aa")
	}
}

func TestLinkedList_ToSlice(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb")
	res := list.ToSlice()
	if len(res) != 2 {
		t.Error("slice len is not 2")
	}
	if res[0] != "aa" {
		t.Error("slice index 0 is not aa")
	}
	if res[1] != "bb" {
		t.Error("slice index 1 is not bb")
	}
	t.Log(list.ToSlice())
}

func TestLinkedList_Add(t *testing.T) {
	list := NewLinkedList[string]()
	addFlag := list.Add("aa")
	if !addFlag {
		t.Error("list add return is  false")
	}
	if list.IsEmpty() {
		t.Error("list is empty")
	}
	res, found := list.Get(0)
	if !found {
		t.Error("list not found")
	}
	if res != "aa" {
		t.Error("list index 0 is not aa")
	}
}

func TestLinkedList_Remove(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb", "cc", "ee", "aa", "aa")
	rmRes := list.Remove("aa")
	if !rmRes {
		t.Error("list remove return is false")
	}
	if list.Size() != 3 {
		t.Error("list size is not 3")
	}
	t.Log(list.ToSlice())
}

func TestLinkedList_RemoveIndex(t *testing.T) {
	list := NewLinkedList[string]("aa", "aa")
	list.RemoveIndex(0)
	if list.Size() != 1 {
		t.Error("list is not empty")
	}
}

func TestLinkedList_ContainsAll(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb")
	if !list.ContainsAll([]string{"aa", "bb"}) {
		t.Error("list not contains  aa bb")
	}
	if list.ContainsAll([]string{"aa", "cc"}) {
		t.Error("list contains  aa cc")
	}
}

func TestLinkedList_AddAll(t *testing.T) {
	list := NewLinkedList[string]()
	addAllRes := list.AddAll([]string{"aa", "bb"})
	if !addAllRes {
		t.Error("add all res is false")
	}
	if list.Size() != 2 {
		t.Error("list size is not 2")
	}
	t.Log(list.ToSlice())
}

func TestLinkedList_RemoveAll(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb", "cc")
	rmAllRes := list.RemoveAll([]string{"aa", "bb"})
	if !rmAllRes {
		t.Error("rm all res is false")
	}
	if list.Size() != 1 {
		t.Error("list size is not 2")
	}
	t.Log(list.ToSlice())
}

func TestLinkedList_Clear(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb", "cc")
	list.Clear()
	if !list.IsEmpty() {
		t.Error("list is not empty")
	}
	t.Log(list.ToSlice())
}

func TestLinkedList_ForEach(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb", "cc")
	var res []string
	list.ForEach(func(i int, e string) bool {
		if i == 1 {
			return true
		}
		res = append(res, e)
		return false
	})
	if len(res) != 1 {
		t.Error("res len not 1")
	}
	t.Log(res)
}

func TestLinkedList_MarshalJSON(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb", "cc", "aa")
	res, err := json.Marshal(list)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
}

func TestLinkedList_UnmarshalJSON(t *testing.T) {
	jsonStr := `["aa","bb", "cc", "aa"]`
	list := NewLinkedList[string]("ee")
	err := json.Unmarshal([]byte(jsonStr), list)
	if err != nil {
		t.Error(err)
	}
	t.Log(list.ToSlice())
}

func TestLinkedList_Get(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb", "cc", "aa")
	res, found := list.Get(4)
	if found {
		t.Error("found err")
	}
	t.Log(res)
	res, found = list.Get(3)
	if !found {
		t.Error("index 3 not found")
	}
	t.Log(res)
}

func TestLinkedList_Set(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb", "cc", "aa")
	list.Set(5, "ee")
	if list.Size() != 4 {
		t.Log("list size not 4")
	}
	list.Set(4, "ee")
	if list.Size() != 5 {
		t.Log("list size not 5")
	}
	t.Log(list.ToSlice())
	list.Set(1, "ff")
	res, found := list.Get(1)
	if !found {
		t.Error("index 1 not found")
	}
	if res != "ff" {
		t.Error("set ff err")
	}
	t.Log(list.ToSlice())
}

func TestLinkedList_IndexOf(t *testing.T) {
	list := NewLinkedList[string]("aa", "bb", "cc", "aa")
	res := list.IndexOf("ee")
	if res != -1 {
		t.Error("ee index not -1")
	}
	res = list.IndexOf("aa")
	if res != 0 {
		t.Error("aa index not 0")
	}
}
