package set

import (
	"encoding/json"
	"github.com/zj-kenzhou/go-col/list"
)

type linkedHashSet[E comparable] struct {
	table    map[E]struct{}
	listData list.List[E]
}

func (s *linkedHashSet[E]) Size() int {
	return s.listData.Size()
}

func (s *linkedHashSet[E]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *linkedHashSet[E]) Contains(e E) bool {
	_, ok := s.table[e]
	return ok
}

func (s *linkedHashSet[E]) ToSlice() []E {
	return s.listData.ToSlice()
}

func (s *linkedHashSet[E]) Add(e E) bool {
	if _, contains := s.table[e]; !contains {
		s.listData.Add(e)
		s.table[e] = struct{}{}
		return true
	}
	return false
}

func (s *linkedHashSet[E]) Remove(e E) bool {
	if s.Contains(e) {
		s.listData.Remove(e)
		delete(s.table, e)
		return true
	}
	return false
}

func (s *linkedHashSet[E]) ContainsAll(col []E) bool {
	for _, e := range col {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

func (s *linkedHashSet[E]) AddAll(col []E) bool {
	res := false
	for _, e := range col {
		if s.Add(e) {
			res = true
		}
	}
	return res
}

func (s *linkedHashSet[E]) RemoveAll(col []E) bool {
	res := false
	for _, e := range col {
		if s.Remove(e) {
			res = true
		}
	}
	return res
}

func (s *linkedHashSet[E]) Clear() {
	for a := range s.table {
		delete(s.table, a)
	}
	s.listData.Clear()
}

func (s *linkedHashSet[E]) ForEach(f func(E)) {
	for e := range s.table {
		f(e)
	}
}

func (s *linkedHashSet[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToSlice())
}

func (s *linkedHashSet[E]) UnmarshalJSON(bytes []byte) error {
	var slice []E
	err := json.Unmarshal(bytes, &slice)
	if err != nil {
		return err
	}
	if !s.IsEmpty() {
		s.Clear()
	}
	s.AddAll(slice)
	return nil
}
