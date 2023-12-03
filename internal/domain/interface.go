package domain

import "context"

type Sender interface {
	Send(follower Follower, dto CheckResultDTO) (string, error)
}

type Runner interface {
	// todo add error return
	Run(ctx context.Context)
}
