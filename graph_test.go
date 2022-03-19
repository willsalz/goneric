package goneric

import (
	"testing"
	// "github.com/google/go-cmp/cmp"
)

func TestUndirectedEdgeListGraph(t *testing.T) {
	testCases := []struct {
		desc string

		edges map[float64]Tuple[string, string]
		from  string
		to    string

		shouldHavePath bool
	}{
		{
			desc: "happy path",
			edges: map[float64]Tuple[string, string]{
				0: {"foo", "bar"},
				1: {"bar", "baz"},
			},
			from:           "foo",
			to:             "baz",
			shouldHavePath: true,
		},
		{
			desc: "islands",
			edges: map[float64]Tuple[string, string]{
				0: {"foo", "bar"},
				1: {"bar", "baz"},
				2: {"baz", "foo"},
				4: {"qux", "mux"},
				5: {"mux", "qux"},
			},
			from:           "foo",
			to:             "qux",
			shouldHavePath: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			g := NewUndirectedGraph[float64, string]()

			for e, pair := range tc.edges {
				e := e
				pair := pair
				g.AddEdge(e, pair.Left, pair.Right)
			}

			hasPath := g.PathExists(tc.from, tc.to)
			if hasPath != tc.shouldHavePath {
				t.Fatal(tc)
			}
		})
	}
}
