package poller

import (
	"context"
	"testing"
	"time"
)

type mockRunner struct {
	count int
}

func (r *mockRunner) Run(context.Context) {
	r.count++
}

func TestPoll(t *testing.T) {
	t.Parallel()

	r := &mockRunner{}
	poller := New(16*time.Millisecond, r)

	ctx, cancel := context.WithTimeout(context.Background(), 64*time.Millisecond)
	defer cancel()

	poller.Poll(ctx)

	if r.count != 4 {
		t.Errorf("count of polls during period mismatch, want %d, got %d", 3, r.count)
	}
}
