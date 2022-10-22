package vrm1migrate

type sortedSet[T comparable] struct {
	s   []T
	dup map[T]bool
}

func newSortedSet[T comparable]() sortedSet[T] {
	return sortedSet[T]{
		dup: map[T]bool{},
	}
}

func (s *sortedSet[T]) add(v T) {
	if !s.dup[v] {
		s.s = append(s.s, v)
		s.dup[v] = true
	}
}

func (s *sortedSet[T]) slice() []T {
	return s.s
}

func (s *sortedSet[T]) reverse() []T {
	rs := make([]T, len(s.s))
	for i, j := 0, len(s.s)-1; i < j; i, j = i+1, j+1 {
		rs[i], rs[j] = s.s[j], s.s[i]
	}
	return rs
}
