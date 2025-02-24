package list

import (
	"encoding/json"
	"reflect"
)

type linkedList[E any] struct {
	first *element[E]
	last  *element[E]
	size  int
}

type element[E any] struct {
	value E
	prev  *element[E]
	next  *element[E]
}

func (l *linkedList[E]) Size() int {
	return l.size
}

func (l *linkedList[E]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *linkedList[E]) Contains(e E) bool {
	for element := l.first; element != nil; element = element.next {
		if reflect.DeepEqual(element.value, e) {
			return true
		}
	}
	return false
}

func (l *linkedList[E]) ToSlice() []E {
	res := make([]E, 0)
	for element := l.first; element != nil; element = element.next {
		res = append(res, element.value)
	}
	return res
}

func (l *linkedList[E]) Add(e E) bool {
	newElement := &element[E]{value: e, prev: l.last}
	if l.size == 0 {
		l.first = newElement
		l.last = newElement
	} else {
		l.last.next = newElement
		l.last = newElement
	}
	l.size++
	return true
}

func (l *linkedList[E]) Remove(e E) bool {
	res := false
	var removeItem *element[E]
	for element := l.first; element != nil; element = element.next {
		if reflect.DeepEqual(element.value, e) {
			if element.prev == nil && element.next == nil {
				l.first = nil
				l.last = nil
			} else if element.prev == nil {
				element.next.prev = nil
				l.first = element.next
			} else if element.next == nil {
				element.prev.next = nil
				l.last = element.prev
			} else {
				element.prev.next = element.next
				element.next.prev = element.prev
			}
			removeItem = element
			l.size--
			res = true
		}
		if removeItem != nil {
			removeItem = nil
		}
	}
	return res
}

func (l *linkedList[E]) withinRange(index int) bool {
	return index >= 0 && index < l.size
}

func (l *linkedList[E]) RemoveIndex(index int) {
	if !l.withinRange(index) {
		return
	}

	if l.size == 1 {
		l.Clear()
		return
	}
	var element *element[E]
	if l.size-index < index {
		element = l.last
		for e := l.size - 1; e != index; e, element = e-1, element.prev {
		}
	} else {
		element = l.first
		for e := 0; e != index; e, element = e+1, element.next {
		}
	}
	if element == l.first {
		l.first = element.next
	}
	if element == l.last {
		l.last = element.prev
	}
	if element.prev != nil {
		element.prev.next = element.next
	}
	if element.next != nil {
		element.next.prev = element.prev
	}
	element = nil
	l.size--
}

func (l *linkedList[E]) ContainsAll(col []E) bool {
	for _, e := range col {
		if !l.Contains(e) {
			return false
		}
	}
	return true
}

func (l *linkedList[E]) AddAll(col []E) bool {
	res := false
	for _, e := range col {
		if l.Add(e) {
			res = true
		}
	}
	return res
}

func (l *linkedList[E]) RemoveAll(col []E) bool {
	res := false
	for _, e := range col {
		if l.Remove(e) {
			res = true
		}
	}
	return res
}

func (l *linkedList[E]) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
}

func (l *linkedList[E]) ForEach(f func(i int, e E) bool) {
	index := 0
	for element := l.first; element != nil; element = element.next {
		if f(index, element.value) {
			return
		}
		index++
	}
}

func (l *linkedList[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToSlice())
}

func (l *linkedList[E]) UnmarshalJSON(bytes []byte) error {
	elements := make([]E, 0)
	err := json.Unmarshal(bytes, &elements)
	if err == nil {
		l.Clear()
		l.AddAll(elements)
	}
	return err
}

func (l *linkedList[E]) Get(index int) (E, bool) {
	if !l.withinRange(index) {
		tempMap := make(map[int]E)
		value, found := tempMap[index]
		return value, found
	}
	if l.size-index < index {
		element := l.last
		for e := l.size - 1; e != index; e, element = e-1, element.prev {
		}
		return element.value, true
	}
	element := l.first
	for e := 0; e != index; e, element = e+1, element.next {
	}
	return element.value, true
}

func (l *linkedList[E]) Set(index int, e E) {
	if !l.withinRange(index) {
		// Append
		if index == l.size {
			l.Add(e)
		}
		return
	}
	var foundElement *element[E]
	// determine traversal direction, last to first or first to last
	if l.size-index < index {
		foundElement = l.last
		for e := l.size - 1; e != index; {
			e, foundElement = e-1, foundElement.prev
		}
	} else {
		foundElement = l.first
		for e := 0; e != index; {
			e, foundElement = e+1, foundElement.next
		}
	}
	foundElement.value = e
}

func (l *linkedList[E]) IndexOf(e E) int {
	if l.size == 0 {
		return -1
	}
	for index, element := range l.ToSlice() {
		if reflect.DeepEqual(element, e) {
			return index
		}
	}
	return -1
}
