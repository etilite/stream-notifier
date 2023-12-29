package streamchecker

import (
	"fmt"
	"testing"

	"github.com/etilite/stream-notifier/internal/domain/dto"
	"github.com/google/go-cmp/cmp"
)

type mockGetter struct {
}

func (g mockGetter) Get(nick string) (*Stream, error) {
	switch nick {
	case "offline":
		return &Stream{DaNick: nick}, nil
	case "error":
		return nil, fmt.Errorf("some error happened")
	default:
		return &Stream{
			DaNick:     nick,
			PreviewUrl: "url",
			Category:   Category{Type: "category-type", Title: "category-title"},
			Title:      "title",
			IsOnline:   true,
		}, nil
	}
}

func TestChecker_Check(t *testing.T) {
	t.Parallel()
	type fields struct {
		getter VkPlayStreamGetter
	}
	type args struct {
		nicks []string
	}

	tests := map[string]struct {
		fields fields
		args   args
		want   map[string]dto.CheckResultDTO
	}{
		"all online": {
			fields: fields{mockGetter{}},
			args:   args{[]string{"id1", "error", "id2"}},
			want: map[string]dto.CheckResultDTO{
				"id1": dto.NewCheckResult(
					"id1",
					"id1",
					"https://vkplay.live/id1",
					"url",
					"[category-type]category-title",
					"title",
					true,
				),
				"id2": dto.NewCheckResult(
					"id2",
					"id2",
					"https://vkplay.live/id2",
					"url",
					"[category-type]category-title",
					"title",
					true,
				),
			},
		},
		"all errors": {
			fields: fields{mockGetter{}},
			args:   args{[]string{"error", "error"}},
			want:   map[string]dto.CheckResultDTO{},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			nicks := make(chan string)
			go func() {
				defer close(nicks)
				for _, nick := range tt.args.nicks {
					nicks <- nick
				}
			}()

			c := NewChecker(tt.fields.getter, 4)
			results := c.Check(nicks)

			result := make(map[string]dto.CheckResultDTO)
			for s := range results {
				result[s.ID()] = s
				//fmt.Println(s)
			}
			//fmt.Println(result)

			opt := cmp.AllowUnexported(dto.CheckResultDTO{})
			if diff := cmp.Diff(tt.want, result, opt); diff != "" {
				t.Errorf("Check() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
