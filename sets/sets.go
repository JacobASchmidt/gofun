package sets

import (
	"github.com/JacobASchmidt/gofun/slices"
	"github.com/JacobASchmidt/gofun/streams"
)

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
func (s Set[A]) Len() int {
	return len(s)
}

func (s Set[A]) Values() []A {
	ret := make([]A, 0, s.Len())
	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Set[A]) Stream() streams.Stream[A] {
	return slices.Stream(s.Values())
}

type _ streams.Stream[int]

func Collect[A comparable](s streams.Stream[A]) Set[A] {
	return streams.Reduce(s, Set[A]{}, Set[A].Add)
}

func Union[A comparable](a Set[A], b streams.Stream[A]) Set[A] {
	return streams.Reduce(b, a, Set[A].Add)
}

func Intersection[A comparable](a streams.Stream[A], b Set[A]) streams.Stream[A] {
	return a.Filter(func(el A) bool { return b.Contains(el) })
}

func Difference[A comparable](a streams.Stream[A], b Set[A]) streams.Stream[A] {
	return a.Filter(func(el A) bool { return !b.Contains(el) })
}

func SymmetricDifference[A comparable](a Set[A], b Set[A]) streams.Stream[A] {
	a_not_in_b := a.Stream().Filter(func(el A) bool { return !b.Contains(el) })
	b_not_in_a := b.Stream().Filter(func(el A) bool { return !a.Contains(el) })
	return streams.Chain(a_not_in_b, b_not_in_a)
}
