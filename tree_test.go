package tree

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	testCases := []struct {
		desc   string
		root   int
		values []int
	}{
		{
			desc:   "happy path",
			root:   5,
			values: []int{4, 6, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tr := New[int](tc.root)
			for _, val := range tc.values {
				tr.Insert(val)
			}
		})
	}
}
