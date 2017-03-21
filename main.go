package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw loggingMiddleware) YoBro(s string) (output string) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "yobro",
			"input", s,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.next.YoBro(s)
	return
}

// StringService provides operations on a string
type StringService interface {
	YoBro(string) string
}

type stringService struct{}

func (stringService) YoBro(s string) string {
	if s == "" {
		return "woogidi woogidi woogidi..."
	}
	return fmt.Sprintf("let's crank %s", s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

func makeWoogidiEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(woogidiRequest)
		response := svc.YoBro(req.Request)
		return woogidiResponse{response}, nil
	}
}

func decodeWoogidiRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request woogidiRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type woogidiRequest struct {
	Request string `json:"gnarly"`
}

type woogidiResponse struct {
	Response string `json:"response"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	logger := log.NewLogfmtLogger(os.Stderr)

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}

	woogidiHandler := httptransport.NewServer(
		makeWoogidiEndpoint(svc),
		decodeWoogidiRequest,
		encodeResponse,
	)

	http.Handle("/", woogidiHandler)
	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(port, nil))
}
