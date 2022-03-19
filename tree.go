package goneric

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type RecursiveBinaryTree[T constraints.Ordered] struct {
	Value T
	Left  *RecursiveBinaryTree[T]
	Right *RecursiveBinaryTree[T]
}

func New[T constraints.Ordered](root T) *RecursiveBinaryTree[T] {
	return &RecursiveBinaryTree[T]{Value: root}
}

func (t *RecursiveBinaryTree[T]) Insert(val T) error {

	if t == nil {
		return fmt.Errorf("cannot insert into nil RecursiveBinaryTree")
	}

	// equal
	if val == t.Value {
		// TODO(will): return ErrExists on duplicate?
		return nil
	}

	if val < t.Value {
		if t.Left == nil {
			t.Left = &RecursiveBinaryTree[T]{Value: val}
		} else {
			t.Left.Insert(val)
		}
		return nil
	}

	if t.Right == nil {
		t.Right = &RecursiveBinaryTree[T]{Value: val}
	} else {
		t.Right.Insert(val)
	}

	return nil

}
