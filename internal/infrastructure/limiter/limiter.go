package limiter

import (
	"time"
)

type Limiter[T any] struct {
	limit    int
	duration time.Duration
}

func New[T any](limit int, duration time.Duration) *Limiter[T] {
	return &Limiter[T]{
		limit:    limit,
		duration: duration,
	}
}

func (l *Limiter[T]) Limit(in <-chan T) <-chan T {
	out := make(chan T)
	throttle := make(chan time.Time, l.limit)
	// todo use context?
	done := make(chan bool)

	go func() {
		tick := time.NewTicker(l.duration / time.Duration(l.limit))
		defer tick.Stop()
		for t := range tick.C {
			select {
			case throttle <- t:
			case <-done:
				return
			}
		}
	}()

	go func() {
		defer close(out)
		defer close(done)
		for request := range in {
			<-throttle
			out <- request
		}
	}()

	return out
}
