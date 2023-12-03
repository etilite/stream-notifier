package limiter

import (
	"testing"
	"time"
)

func TestLimiter_Limit(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		limit    int
		duration time.Duration
	}{
		"10 requests in 160ms": {
			limit:    10,
			duration: 160 * time.Millisecond,
		},
		"2 requests in 50ms": {
			limit:    2,
			duration: 50 * time.Millisecond,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			requests := make(chan int, tc.limit)
			for i := 0; i < tc.limit; i++ {
				requests <- i
			}
			close(requests)

			l := Limiter[int]{
				limit:    tc.limit,
				duration: tc.duration,
			}
			limitedRequests := l.Limit(requests)

			var results []int
			start := time.Now()
			for s := range limitedRequests {
				results = append(results, s)
			}
			end := time.Now()
			elapsedTime := end.Sub(start)

			//fmt.Println(elapsedTime, len(results))

			if elapsedTime < tc.duration {
				t.Errorf(
					"error: limit bypassed, want %d requests in %v, got %d requests in %v",
					tc.limit,
					tc.duration,
					len(results),
					elapsedTime,
				)
			}
		})
	}
}
