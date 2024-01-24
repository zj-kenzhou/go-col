package list

import "github.com/zj-kenzhou/go-col"

type List[E any] interface {
	col.Collection[E]
	Get(index int) (E, bool)
	Set(index int, e E)
	RemoveIndex(index int)
	IndexOf(e E) int
	ForEach(f func(i int, e E) bool)
}

func NewLinkedList[E any](values ...E) List[E] {
	list := &linkedList[E]{}
	list.AddAll(values)
	return list
}
