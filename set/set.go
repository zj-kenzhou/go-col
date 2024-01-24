package set

import (
	"github.com/zj-kenzhou/go-col"
)

type Set[E comparable] interface {
	col.Collection[E]
	ForEach(func(E) bool)
}

func NewSet[T comparable](values ...T) Set[T] {
	res := hashSet[T]{}
	res.AddAll(values)
	return res
}

func NewSyncSet[T comparable](values ...T) Set[T] {
	res := &syncHashSet[T]{
		uss: hashSet[T]{},
	}
	res.AddAll(values)
	return res
}
