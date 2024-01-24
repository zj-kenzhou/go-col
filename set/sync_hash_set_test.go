package set

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

func TestSyncHashSet_Size(t *testing.T) {
	set := NewSyncSet[string]()
	if set.Size() != 0 {
		t.Error("set is not empty")
	}
	set.Add("aa")
	set.Add("bb")
	if set.Size() != 2 {
		t.Error("set is not empty")
	}
}

func TestSyncHashSet_IsEmpty(t *testing.T) {
	set := NewSyncSet[string]()
	if !set.IsEmpty() {
		t.Error("set is not empty")
	}
	set.Add("aa")
	set.Add("bb")
	if set.IsEmpty() {
		t.Error("set is not empty")
	}
}

func TestSyncHashSet_Contains(t *testing.T) {
	set := NewSyncSet[string]("aaa")
	if !set.Contains("aaa") {
		t.Error("aaa is not Contains")
	}
	if set.Contains("aa") {
		t.Error("aaa is not Contains")
	}
}

func TestSyncHashSet_ToSlice(t *testing.T) {
	set := NewSyncSet[string]("aaa")
	for i := 0; i < 10; i++ {
		set.Add(fmt.Sprintf("%v", i))
	}
	t.Log(set.ToSlice())
}

func TestSyncHashSet_Add(t *testing.T) {
	set := NewSyncSet[string]()
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

func TestSyncHashSet_Remove(t *testing.T) {
	set := NewSyncSet[string]("aa")
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

func TestSyncHashSet_Remove_Concurrent(t *testing.T) {
	set := NewSyncSet[string]()
	for i := 0; i < 10000; i++ {
		set.Add(fmt.Sprintf("%v", i))
	}
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(index int) {
			set.Remove(fmt.Sprintf("%v", index))
			wg.Done()
		}(i)
	}
	wg.Wait()
	if !set.IsEmpty() {
		t.Error("set is not empty")
	}
}

func TestSyncHashSet_ContainsAll(t *testing.T) {
	set := NewSyncSet[string]("aa", "bb")
	if !set.ContainsAll([]string{"aa", "bb"}) {
		t.Error("ContainsAll is false")
	}
	if set.ContainsAll([]string{"aa", "cc"}) {
		t.Error("ContainsAll is true")
	}
}

func TestSyncHashSet_AddAll(t *testing.T) {
	set := NewSyncSet[string]()
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

func TestSyncHashSet_Add_Concurrent(t *testing.T) {
	set := NewSyncSet[string]()
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(index int) {
			set.Add(fmt.Sprintf("%v", index))
			wg.Done()
		}(i)
	}
	wg.Wait()
	if set.Size() != 10000 {
		t.Error("set size is not 1000")
	}
}

func TestSyncHashSet_RemoveAll(t *testing.T) {
	set := NewSyncSet[string]("aa", "bb")
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

func TestSyncHashSet_Clear(t *testing.T) {
	set := NewSyncSet[string]("aa", "bb")
	set.Clear()
	if !set.IsEmpty() {
		t.Error("set is not empty")
	}
}

func TestSyncHashSet_ForEach(t *testing.T) {
	set := NewSyncSet[string]("aa", "bb", "cc")
	set.ForEach(func(s string) {
		t.Log(s)
	})
}

func TestSyncHashSet_MarshalJSON(t *testing.T) {
	set := NewSyncSet[string]("aa", "bb", "cc")
	bytes, err := json.Marshal(set)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bytes))
}

func TestSyncHashSet_UnmarshalJSON(t *testing.T) {
	jsonStr := "[\"aa\", \"bb\"]"
	set := NewSyncSet[string]("cc")
	err := json.Unmarshal([]byte(jsonStr), set)
	if err != nil {
		t.Error(err)
	}
	if set.Contains("cc") {
		t.Error("element not remove")
	}
	t.Log(set.ToSlice())
}
