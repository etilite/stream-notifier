package domain

import (
	"context"

	"github.com/etilite/stream-notifier/internal/domain/dto"
)

type StreamChecker interface {
	Check(in <-chan string) <-chan dto.CheckResultDTO
}

type Splitter interface {
	Split(in <-chan dto.CheckResultDTO) (<-chan dto.PersonalCheckResultDTO, <-chan *Notification)
}

type OnlineResultsHandler interface {
	Handle(ctx context.Context, in <-chan dto.PersonalCheckResultDTO)
}

type OfflineResultsHandler interface {
	Handle(ctx context.Context, in <-chan Notification)
}

type Sender interface {
	Send(follower Follower, dto dto.CheckResultDTO) (string, error)
}

type Doer interface {
	// todo add error return?
	Do(ctx context.Context)
}
