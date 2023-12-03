package poller

import (
	"context"
	"time"

	"github.com/etilite/stream-notifier/internal/domain"
)

type Poller struct {
	interval time.Duration
	runner   domain.Runner
}

func New(interval time.Duration, runner domain.Runner) *Poller {
	return &Poller{
		interval: interval,
		runner:   runner,
	}
}

func (p *Poller) Poll(ctx context.Context) {
	ticker := time.NewTicker(p.interval)

	for {
		select {
		case <-ticker.C:
			//fmt.Println(time.Now())
			p.runner.Run(ctx)
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}
