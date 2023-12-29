package poller

import (
	"context"
	"testing"
	"time"
)

type mockDoer struct {
	count int
}

func (r *mockDoer) Do(context.Context) {
	r.count++
}

func TestPoll(t *testing.T) {
	t.Parallel()

	r := &mockDoer{}
	poller := New(16*time.Millisecond, r)

	ctx, cancel := context.WithTimeout(context.Background(), 64*time.Millisecond)
	defer cancel()

	poller.Poll(ctx)

	if r.count != 5 {
		t.Errorf("count of polls during period mismatch, want %d, got %d", 5, r.count)
	}
}
