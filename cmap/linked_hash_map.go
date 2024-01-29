package cmap

import (
	"bytes"
	"encoding/json"
	"github.com/zj-kenzhou/go-col/list"
	"sort"
)

type linkedHashMap[K comparable, V any] struct {
	table    map[K]V
	ordering list.List[K]
}

func (m *linkedHashMap[K, V]) Size() int {
	return m.ordering.Size()
}

func (m *linkedHashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

func (m *linkedHashMap[K, V]) ContainsKey(k K) bool {
	_, ok := m.table[k]
	return ok
}

func (m *linkedHashMap[K, V]) Found(k K) (V, bool) {
	value, found := m.table[k]
	return value, found
}

func (m *linkedHashMap[K, V]) Get(k K) V {
	value := m.table[k]
	return value
}

func (m *linkedHashMap[K, V]) Put(k K, v V) {
	if _, contains := m.table[k]; !contains {
		m.ordering.Add(k)
	}
	m.table[k] = v
}

func (m *linkedHashMap[K, V]) Remove(k K) {
	if _, contains := m.table[k]; contains {
		delete(m.table, k)
		index := m.ordering.IndexOf(k)
		m.ordering.RemoveIndex(index)
	}
}

func (m *linkedHashMap[K, V]) Clear() {
	m.table = make(map[K]V)
	m.ordering.Clear()
}

func (m *linkedHashMap[K, V]) Keys() []K {
	return m.ordering.ToSlice()
}

func (m *linkedHashMap[K, V]) Values() []V {
	values := make([]V, m.Size())
	count := 0
	for _, value := range m.table {
		values[count] = value
		count++
	}
	return values
}

func (m *linkedHashMap[K, V]) ForEach(f func(k K, v V) bool) {
	for key, value := range m.table {
		if f(key, value) {
			return
		}
	}
}

func (m *linkedHashMap[K, V]) MarshalJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)

	buf.WriteRune('{')

	lastIndex := m.Size() - 1
	for index, key := range m.Keys() {
		km, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		buf.Write(km)

		buf.WriteRune(':')

		vm, err := json.Marshal(m.table[key])
		if err != nil {
			return nil, err
		}
		buf.Write(vm)
		if index != lastIndex {
			buf.WriteRune(',')
		}

		index++
	}
	buf.WriteRune('}')

	return buf.Bytes(), nil
}

func (m *linkedHashMap[K, V]) UnmarshalJSON(data []byte) error {
	elements := make(map[K]V)
	err := json.Unmarshal(data, &elements)
	if err != nil {
		return err
	}
	index := make(map[K]int)
	var keys []K
	for key := range elements {
		keys = append(keys, key)
		esc, _ := json.Marshal(key)
		index[key] = bytes.Index(data, esc)
	}
	byIndex := func(i, j int) bool {
		key1 := keys[i]
		key2 := keys[j]
		index1 := index[key1]
		index2 := index[key2]
		return index1 > index2
	}
	sort.Slice(keys, byIndex)

	m.Clear()

	for _, key := range keys {
		m.Put(key, elements[key])
	}

	return nil
}
