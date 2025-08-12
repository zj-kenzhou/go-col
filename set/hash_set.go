package set

import (
	"encoding/json"
	"iter"
)

type hashSet[E comparable] map[E]struct{}

func (s *hashSet[E]) Size() int {
	return len(*s)
}

func (s *hashSet[E]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *hashSet[E]) Contains(e E) bool {
	_, ok := (*s)[e]
	return ok
}

func (s *hashSet[E]) ToSlice() []E {
	res := make([]E, 0)
	for e := range *s {
		res = append(res, e)
	}
	return res
}

func (s *hashSet[E]) Add(e E) bool {
	if s.Contains(e) {
		return false
	}
	(*s)[e] = struct{}{}
	return true
}

func (s *hashSet[E]) Remove(e E) bool {
	res := false
	if s.Contains(e) {
		res = true
	}
	delete(*s, e)
	return res
}

func (s *hashSet[E]) ContainsAll(col []E) bool {
	for _, e := range col {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

func (s *hashSet[E]) AddAll(col []E) bool {
	res := false
	for _, e := range col {
		if s.Add(e) {
			res = true
		}
	}
	return res
}

func (s *hashSet[E]) RemoveAll(col []E) bool {
	res := false
	for _, e := range col {
		if s.Remove(e) {
			res = true
		}
	}
	return res
}

func (s *hashSet[E]) Clear() {
	for a := range *s {
		delete(*s, a)
	}
}

func (s *hashSet[E]) ForEach(f func(E)) {
	for e := range *s {
		f(e)
	}
}

func (s *hashSet[E]) Iterator() iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range *s {
			if !yield(e) {
				return
			}
		}
	}
}

func (s *hashSet[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToSlice())
}

func (s *hashSet[E]) UnmarshalJSON(bytes []byte) error {
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
