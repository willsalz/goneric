package tree

import (
	"golang.org/x/exp/constraints"
)

type Tree[T constraints.Ordered] struct {
	value T
	Left  *Tree[T]
	Right *Tree[T]
}

func New[T constraints.Ordered](root T) *Tree[T] {
	return &Tree[T]{value: root}
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
		if t.Left == nil {
			t.Left = &Tree[T]{value: val}
			return
		}
		t.Left.Insert(val)
		return
	}

	if t.Right == nil {
		t.Right = &Tree[T]{value: val}
		return
	}
	t.Right.Insert(val)

}
