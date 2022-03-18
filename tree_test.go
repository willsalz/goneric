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
			values: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tr := &Tree[int]{}
			for _, val := range tc.values {
				tr.Insert(val)
			}
			t.Fatal(tr)
		})
	}
}
