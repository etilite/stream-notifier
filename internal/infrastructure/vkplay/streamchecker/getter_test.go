package streamchecker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	server *httptest.Server
)

func TestMain(m *testing.M) {
	server = setupMockServer()
	defer server.Close()

	m.Run()
}

func TestGetter_Get(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		baseUrl string
		timeout time.Duration
		stream  *Stream
		errMsg  string
	}{
		"build request failed": {
			":foo" + "%s",
			0 * time.Millisecond,
			nil,
			"failed to build request: parse \":foo\": missing protocol scheme",
		},
		"connection error": {
			"no-such-url" + "%s",
			1 * time.Millisecond,
			nil,
			"failed to get API response: Get \"no-such-url\": unsupported protocol scheme \"\"",
		},
		"response code is not 200": {

			fmt.Sprintf("%s/no-such-path/", server.URL) + "%s",
			0 * time.Millisecond,
			nil,
			"response code is 404",
		},
		"json error": {
			fmt.Sprintf("%s/json-error/", server.URL) + "%s",
			0 * time.Millisecond,
			nil,
			"failed to decode JSON response: EOF",
		},
		"stream is offline": {
			server.URL + "%s",
			0 * time.Millisecond,
			&Stream{},
			"",
		},
		"stream is online": {
			fmt.Sprintf("%s/2/", server.URL) + "%s",
			0 * time.Millisecond,
			&Stream{
				IsOnline: true,
			},
			"",
		},
	}

	for name, tt := range tests {
		tt := tt // эта магия нужна, чтобы при параллельном запуске тестов каверадж не искажался
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			checker := New(
				tt.baseUrl,
				&http.Client{
					Timeout: tt.timeout,
				},
			)
			result, err := checker.Get("")
			if !reflect.DeepEqual(tt.stream, result) {
				t.Errorf("resultMap mismatch: want %v, got %v", tt.stream, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != tt.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`", tt.errMsg, errMsg)
			}
		})
	}
}

func setupMockServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimSpace(r.URL.Path) {
		case "/":
			jsonResponse := `{
				"isOnline": false
			}`
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(jsonResponse))
		case "/2/":
			jsonResponse := `{
				"isOnline": true
			}`
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(jsonResponse))
		case "/json-error/":
			jsonResponse := ""
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(jsonResponse))
		case "/no-such-path/":
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))
	return server
}
