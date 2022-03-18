package tree

import (
	"golang.org/x/exp/constraints"
)

type Tree[T constraints.Ordered] struct {
	value T
	Left  *Tree[T]
	Right *Tree[T]
}

func (t *Tree[T]) Insert(val T) {
	if t == nil {
		(*t) = Tree[T]{value: val}
		return
	}

	if val == t.value {
		return
	}

	if val < t.value {
		return t.Left.Insert(val)	
	}

	return t.Right.Insert(val)

}
