package cmap

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLinkedHashMap_Size(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	if linkedHashMap.Size() != 0 {
		t.Error("linkedHashMap size not 0")
	}
	linkedHashMap.Put("aa", "aa")
	linkedHashMap.Put("aa", 11)
	if linkedHashMap.Size() != 1 {
		t.Error("linkedHashMap size not 1")
	}
}

func TestLinkedHashMap_IsEmpty(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	if !linkedHashMap.IsEmpty() {
		t.Error("linkedHashMap is not empty")
	}
	linkedHashMap.Put("aa", "aa")
	if linkedHashMap.IsEmpty() {
		t.Error("linkedHashMap is empty")
	}
}

func TestLinkedHashMap_ContainsKey(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	if linkedHashMap.ContainsKey("aaa") {
		t.Error("linkedHashMap ContainsKey aaa")
	}
	linkedHashMap.Put("aa", "aa")
	if !linkedHashMap.ContainsKey("aa") {
		t.Error("linkedHashMap not ContainsKey  aa")
	}
}

func TestLinkedHashMap_Found(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	res, found := linkedHashMap.Found("aa")
	if found {
		t.Error("linkedHashMap found aa")
	}
	t.Log(res)
	linkedHashMap.Put("aa", "aa")
	res, found = linkedHashMap.Found("aa")
	if !found {
		t.Error("linkedHashMap not found aa")
	}
	t.Log(res)
}

func TestLinkedHashMap_Get(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	res := linkedHashMap.Get("aa")
	if res != nil {
		t.Error("linkedHashMap found aa")
	}
	t.Log(res)
	linkedHashMap.Put("aa", "aa")
	res = linkedHashMap.Get("aa")
	if res != "aa" {
		t.Error("linkedHashMap not found aa")
	}
	t.Log(res)
}

func TestLinkedHashMap_Put(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	linkedHashMap.Put("aa", 1)
	if linkedHashMap.Size() != 1 {
		t.Error("linkedHashMap size not 1")
	}
}

func TestLinkedHashMap_Remove(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	linkedHashMap.Put("aa", 1)
	linkedHashMap.Remove("aa")
	if linkedHashMap.Size() != 0 {
		t.Error("linkedHashMap size not 0")
	}
}

func TestLinkedHashMap_Clear(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	linkedHashMap.Put("aa", 1)
	linkedHashMap.Put("bb", 1)
	linkedHashMap.Clear()
	if linkedHashMap.Size() != 0 {
		t.Error("linkedHashMap size not 0")
	}
}

func TestLinkedHashMap_Keys(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	linkedHashMap.Put("aa", 1)
	linkedHashMap.Put("bb", 1)
	keys := linkedHashMap.Keys()
	if keys[0] != "aa" {
		t.Error("key 0 is not aa")
	}
	if keys[1] != "bb" {
		t.Error("key 1 is not bb")
	}
}
func TestLinkedHashMap_Values(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, any]()
	linkedHashMap.Put("aa", 1)
	linkedHashMap.Put("bb", 2)
	values := linkedHashMap.Values()
	if values[0] != 1 {
		t.Error("values 0 is not 1")
	}
	if values[1] != 2 {
		t.Error("values 1 is not 2")
	}
}

func TestLinkedHashMap_ForEach(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, int]()
	linkedHashMap.Put("aa", 1)
	linkedHashMap.Put("bb", 2)

	linkedHashMap.ForEach(func(k string, v int) bool {
		if k == "aa" && v != 1 {
			t.Error("k v err")
		}
		if k == "bb" && v != 2 {
			t.Error("k v err")
		}
		t.Log(fmt.Sprintf("key is %s value is %v", k, v))
		if k == "aa" {
			return true
		}
		return false
	})
}

func TestLinkedHashMap_Iterator(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, int]()
	linkedHashMap.Put("aa", 1)
	linkedHashMap.Put("bb", 2)

	for k, v := range linkedHashMap.Iterator() {
		if k == "aa" && v != 1 {
			t.Error("k v err")
		}
		if k == "bb" && v != 2 {
			t.Error("k v err")
		}
		t.Log(fmt.Sprintf("key is %s value is %v", k, v))
		if k == "aa" {
			break
		}
	}
}

func TestLinkedHashMap_MarshalJSON(t *testing.T) {
	linkedHashMap := NewLinkedHashMap[string, int]()
	linkedHashMap.Put("aa", 1)
	linkedHashMap.Put("bb", 2)
	res, err := json.Marshal(linkedHashMap)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
}

func TestLinkedHashMap_UnmarshalJSON(t *testing.T) {
	jsonStr := `{ "aa" : "aa", "bb": 22}`
	linkedHashMap := NewLinkedHashMap[string, any]()
	err := json.Unmarshal([]byte(jsonStr), linkedHashMap)
	if err != nil {
		t.Error(err)
	}
	if linkedHashMap.Size() != 2 {
		t.Error("linkedHashMap size is not 2")
	}
	t.Log(linkedHashMap)
}
