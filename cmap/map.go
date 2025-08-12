package cmap

import (
	"encoding/json"
	"iter"

	"github.com/zj-kenzhou/go-col/list"
)

type Map[K comparable, V any] interface {
	Size() int
	IsEmpty() bool
	ContainsKey(k K) bool
	Found(k K) (V, bool)
	Get(k K) V
	Put(k K, v V)
	Remove(k K)
	Clear()
	Keys() []K
	Values() []V
	ForEach(f func(k K, v V) bool)
	Iterator() iter.Seq2[K, V]
	json.Marshaler
	json.Unmarshaler
}

func NewLinkedHashMap[K comparable, V any]() Map[K, V] {
	return &linkedHashMap[K, V]{
		table:    make(map[K]V),
		ordering: list.NewLinkedList[K](),
	}
}
