package col

import "encoding/json"

type Collection[E any] interface {
	Size() int
	IsEmpty() bool
	Contains(e E) bool
	ToSlice() []E
	Add(e E) bool
	Remove(e E) bool
	ContainsAll(col []E) bool
	AddAll(col []E) bool
	RemoveAll(col []E) bool
	Clear()
	json.Marshaler
	json.Unmarshaler
}
