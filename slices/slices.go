package slices

import (
	"github.com/JacobASchmidt/gofun/streams"
)

func Stream[A any](s []A) streams.Stream[A] {
	if len(s) == 0 {
		return nil
	}
	return func() (A, streams.Stream[A]) {
		return s[0], Stream(s[1:])
	}
}
