package goneric

///////////
// Tuple //
///////////

type Tuple[T comparable, U comparable] struct {
	Left  T
	Right U
}

//////////////
// Optional //
//////////////

type Optional[T comparable] struct {
	element *T
}

func (o Optional[T]) Or(other T) T {
	if o.element == nil {
		return other
	}
	return *o.element
}

func (o Optional[T]) UnsafeUnwrap() T {
	return *o.element
}

func Some[T comparable](e *T) Optional[T] {
	return Optional[T]{element: e}
}

func None[T comparable]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) IsNone() bool {
	if o.element == nil {
		return true
	}
	return false
}

///////////
// Queue //
///////////

type Queue[T comparable] struct {
	entries []T
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		entries: []T{},
	}
}

func (q *Queue[T]) PushBack(e T) {
	q.entries = append(q.entries, e)
}

func (q *Queue[T]) PopFront() Optional[T] {
	if len(q.entries) > 0 {
		var head T
		head, q.entries = q.entries[0], q.entries[1:]
		return Some(&head)
	}

	return None[T]()
}

/////////
// Set //
/////////

type Set[E comparable] struct {
	elements map[E]bool
}

func NewSet[E comparable]() *Set[E] {
	return &Set[E]{
		elements: map[E]bool{},
	}
}

func (s Set[E]) Add(e E) {
	s.elements[e] = true
}

func (s *Set[E]) Remove(e E) {
	delete(s.elements, e)
}

func (s *Set[E]) Contains(e E) bool {
	_, ok := s.elements[e]
	return ok
}

func (s *Set[_]) Empty() bool {
	return len(s.elements) == 0
}

//////////////////////
// Undirected Graph //
//////////////////////

type UndirectedGraph[E comparable, N comparable] struct {
	g map[N]map[N]E
}

func NewUndirectedGraph[E comparable, N comparable]() *UndirectedGraph[E, N] {
	return &UndirectedGraph[E, N]{
		g: map[N]map[N]E{},
	}
}

func (g *UndirectedGraph[E, N]) AddNode(node N) {
	if _, ok := g.g[node]; ok {
		return
	}
	g.g[node] = map[N]E{}
}

func (g *UndirectedGraph[E, N]) AddEdge(e E, from, to N) {
	g.AddNode(from)
	g.AddNode(to)
	g.g[from][to] = e
	g.g[to][from] = e
}

func (g *UndirectedGraph[_, N]) Neighbors(node N) []N {
	m, ok := g.g[node]
	if !ok {
		return nil
	}

	neighbors := []N{}
	for n, _ := range m {
		neighbors = append(neighbors, n)
	}

	return neighbors
}

func (g *UndirectedGraph[_, N]) PathExists(from, to N) bool {
	visited := NewSet[N]()
	toVisit := NewQueue[N]()
	toVisit.PushBack(from)

	for currOpt := toVisit.PopFront(); !currOpt.IsNone(); currOpt = toVisit.PopFront() {
		curr := currOpt.UnsafeUnwrap()
		visited.Add(curr)

		neighbors := g.Neighbors(curr)
		for _, neighbor := range neighbors {
			// we've arrived
			if neighbor == to {
				return true
			}

			// we've been here
			if visited.Contains(neighbor) {
				continue
			}

			// we need to explore
			toVisit.PushBack(neighbor)
		}
	}

	return false
}
