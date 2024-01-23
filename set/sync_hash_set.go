package set

import (
	"encoding/json"
	"sync"
)

type syncHashSet[E comparable] struct {
	sync.RWMutex
	uss hashSet[E]
}

func (s *syncHashSet[E]) Size() int {
	return s.uss.Size()
}

func (s *syncHashSet[E]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *syncHashSet[E]) Contains(e E) bool {
	return s.uss.Contains(e)
}

func (s *syncHashSet[E]) ToSlice() []E {
	return s.uss.ToSlice()
}

func (s *syncHashSet[E]) Add(e E) bool {
	s.Lock()
	res := s.uss.Add(e)
	s.Unlock()
	return res
}

func (s *syncHashSet[E]) Remove(e E) bool {
	s.Lock()
	res := s.uss.Remove(e)
	s.Unlock()
	return res
}

func (s *syncHashSet[E]) ContainsAll(col []E) bool {
	return s.uss.ContainsAll(col)
}

func (s *syncHashSet[E]) AddAll(col []E) bool {
	s.Lock()
	defer s.Unlock()
	res := true
	for _, e := range col {
		if !s.uss.Add(e) {
			res = false
		}
	}
	return res
}

func (s *syncHashSet[E]) RemoveAll(col []E) bool {
	s.Lock()
	defer s.Unlock()
	res := true
	for _, e := range col {
		if !s.uss.Remove(e) {
			res = false
		}
	}
	return res
}

func (s *syncHashSet[E]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.uss.Clear()
}

func (s *syncHashSet[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToSlice())
}

func (s *syncHashSet[E]) UnmarshalJSON(bytes []byte) error {
	var slice []E
	err := json.Unmarshal(bytes, &slice)
	if err != nil {
		return err
	}
	s.AddAll(slice)
	return nil
}
