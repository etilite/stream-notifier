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

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			requests := make(chan int, tt.limit)
			for i := 0; i < tt.limit; i++ {
				requests <- i
			}
			close(requests)

			l := New[int](tt.limit, tt.duration)
			limitedRequests := l.Limit(requests)

			var results []int
			start := time.Now()
			for s := range limitedRequests {
				results = append(results, s)
			}
			end := time.Now()
			elapsedTime := end.Sub(start)

			if elapsedTime < tt.duration {
				t.Errorf(
					"error: limit bypassed, want %d requests in %v, got %d requests in %v",
					tt.limit,
					tt.duration,
					len(results),
					elapsedTime,
				)
			}
		})
	}
}
