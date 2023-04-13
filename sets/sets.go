package sets

type Set[A comparable] map[A]struct{}

func (s Set[A]) Add(a A) Set[A] {
	s[a] = struct{}{}
	return s
}

func (s Set[A]) Remove(a A) Set[A] {
	delete(s, a)
	return s
}

func (s Set[A]) Contains(a A) bool {
	_, ok := s[a]
	return ok
}
