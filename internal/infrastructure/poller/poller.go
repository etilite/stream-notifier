package poller

import (
	"context"
	"time"

	"github.com/etilite/stream-notifier/internal/domain"
)

type Poller struct {
	interval time.Duration
	doer     domain.Doer
}

func New(interval time.Duration, doer domain.Doer) *Poller {
	return &Poller{
		interval: interval,
		doer:     doer,
	}
}

func (p *Poller) Poll(ctx context.Context) {
	ticker := time.NewTicker(p.interval)

	p.doer.Do(ctx)

	for {
		select {
		case <-ticker.C:
			p.doer.Do(ctx)
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}
