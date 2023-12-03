package limiter

import (
	"time"
)

type Limiter[T any] struct {
	limit    int
	duration time.Duration
}

func (l Limiter[T]) Limit(in <-chan T) <-chan T {
	out := make(chan T)
	quotas := make(chan time.Time, l.limit)

	go func() {
		tick := time.NewTicker(l.duration / time.Duration(l.limit))
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	go func() {
		defer close(out)
		for request := range in {
			//t := <-quotas
			<-quotas
			//fmt.Println("request", request, t)
			out <- request
		}
	}()

	return out
}
