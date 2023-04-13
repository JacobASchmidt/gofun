package streams

type Stream[A any] func() (A, Stream[A])

func Of[A any](vals ...A) Stream[A] {
	if len(vals) == 0 {
		return nil
	}
	return func() (A, Stream[A]) {
		return vals[0], Of(vals[1:]...)
	}
}

func Map[A any, B any](a Stream[A], f func(A) B) Stream[B] {
	if a == nil {
		return nil
	}
	return func() (B, Stream[B]) {
		first, rest := a()
		return f(first), Map(rest, f)
	}
}

func (a Stream[A]) Filter(f func(A) bool) Stream[A] {
	if a == nil {
		return nil
	}
	first, rest := a()
	if f(first) {
		return func() (A, Stream[A]) {
			return first, rest.Filter(f)
		}
	}
	return rest.Filter(f)
}

func Reduce[A any, B any](a Stream[A], init B, f func(B, A) B) B {
	if a == nil {
		return init
	}
	first, rest := a()
	return Reduce(rest, f(init, first), f)
}

func (a Stream[A]) Take(n int) Stream[A] {
	if a == nil || n == 0 {
		return nil
	}
	return func() (A, Stream[A]) {
		first, rest := a()
		return first, rest.Take(n - 1)
	}
}

func (a Stream[A]) TakeWhile(f func(A) bool) Stream[A] {
	if a == nil {
		return nil
	}
	first, rest := a()
	if f(first) {
		return func() (A, Stream[A]) {
			return first, rest.TakeWhile(f)
		}
	}
	return nil
}

func (a Stream[A]) DropWhile(f func(A) bool) Stream[A] {
	if a == nil {
		return nil
	}
	first, rest := a()
	if f(first) {
		return rest.DropWhile(f)
	}
	return a
}
