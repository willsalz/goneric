package goneric

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestRecursiveBinaryTree(t *testing.T) {
	testCases := []struct {
		desc   string
		root   int
		values []int

		want *RecursiveBinaryTree[int]
	}{
		{
			desc:   "happy path",
			root:   5,
			values: []int{4, 6, 3},
			want: &RecursiveBinaryTree[int]{
				Value: 5,
				Left: &RecursiveBinaryTree[int]{
					Value: 4,
					Left: &RecursiveBinaryTree[int]{
						Value: 3,
					},
				},
				Right: &RecursiveBinaryTree[int]{Value: 6},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tr := New[int](tc.root)
			for _, val := range tc.values {
				tr.Insert(val)
			}

			if diff := cmp.Diff(tc.want, tr); diff != "" {
				t.Fatalf("error: (-want, +got)\n%s", diff)
			}
		})
	}
}
