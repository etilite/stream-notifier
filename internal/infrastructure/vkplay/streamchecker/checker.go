package streamchecker

import (
	"fmt"
	"log"
	"sync"

	"github.com/etilite/stream-notifier/internal/domain/dto"
)

const streamBaseUrl = "https://vkplay.live/"

type VkPlayStreamGetter interface {
	Get(name string) (*Stream, error)
}

type Checker struct {
	getter  VkPlayStreamGetter
	workers int
}

func NewChecker(getter VkPlayStreamGetter, workers int) *Checker {
	if workers < 0 {
		workers = 1
	}
	return &Checker{
		getter:  getter,
		workers: workers,
	}
}

func (c *Checker) Check(in <-chan string) <-chan dto.CheckResultDTO {
	// todo add limiter
	var wg sync.WaitGroup
	out := make(chan dto.CheckResultDTO)

	for i := 0; i < c.workers; i++ {
		wg.Add(1)
		go func() {
			// todo handle panic
			defer wg.Done()
			for nick := range in {
				stream, err := c.getter.Get(nick)
				if err != nil {
					// todo separate ch for errors
					log.Printf("error getting \"%s\" stream: %s", nick, err)
					continue
				}
				checkResult := dto.NewCheckResult(
					nick,
					stream.DaNick,
					fmt.Sprintf("%s%s", streamBaseUrl, nick),
					stream.PreviewUrl,
					fmt.Sprintf("[%s]%s", stream.Category.Type, stream.Category.Title),
					stream.Title,
					stream.IsOnline,
				)
				out <- checkResult
			}
		}()
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}
