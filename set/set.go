package set

import (
	"iter"

	"github.com/zj-kenzhou/go-col"
	"github.com/zj-kenzhou/go-col/list"
)

type Set[E comparable] interface {
	col.Collection[E]
	ForEach(func(E))
	Iterator() iter.Seq[E]
}

func NewSet[T comparable](values ...T) Set[T] {
	res := hashSet[T]{}
	res.AddAll(values)
	return &res
}

func NewLinkedHashSet[T comparable](values ...T) Set[T] {
	res := &linkedHashSet[T]{
		table:    make(map[T]struct{}),
		listData: list.NewLinkedList[T](),
	}
	res.AddAll(values)
	return res
}

func NewSyncSet[T comparable](values ...T) Set[T] {
	res := &syncHashSet[T]{
		uss: &hashSet[T]{},
	}
	res.AddAll(values)
	return res
}
