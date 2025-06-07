package httptool

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/common/util"
	"io"
	"net/http"
	"time"
)

// 请求方法
// 请求头
// 请求体
// traceId
// TODO 可以加上鉴权

type requestOption struct {
	ctx     context.Context
	header  map[string]string
	data    []byte
	timeout time.Duration
}

type Option func(*requestOption) error

func (f Option) Apply(opts *requestOption) error {
	return f(opts)
}

func defaultRequestOption() *requestOption {
	return &requestOption{
		ctx:     context.Background(),
		header:  map[string]string{},
		timeout: 10 * time.Second,
	}
}

func WithContext(ctx context.Context) Option {
	return Option(func(req *requestOption) error {
		req.ctx = ctx
		return nil
	})
}

func WithHeaders(header map[string]string) Option {
	return Option(func(req *requestOption) error {
		for k, v := range header {
			req.header[k] = v
		}
		return nil
	})
}

func WithData(data []byte) Option {
	return Option(func(req *requestOption) error {
		req.data = data
		return nil
	})
}

func WithTimeout(timeout time.Duration) Option {
	return Option(func(req *requestOption) error {
		req.timeout = timeout
		return nil
	})
}

func WithAuthorization(token string) Option {
	return Option(func(req *requestOption) error {
		req.header["Authorization"] = fmt.Sprintf("Bearer %s", token)
		return nil
	})
}

func Request(method string, url string, opts ...Option) (ResultCode int, body []byte, err error) {
	startTime := time.Now()
	reqOpts := defaultRequestOption()
	for i := range opts {
		if err = opts[i].Apply(reqOpts); err != nil {
			return 0, nil, err
		}
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(reqOpts.data))
	if err != nil {
		return
	}
	if len(reqOpts.header) > 0 {
		for k, v := range reqOpts.header {
			req.Header.Set(k, v)
		}
	}
	traceId, spanId, _ := util.GetTraceIdFromContext(reqOpts.ctx)
	req.Header.Set("traceId", traceId)
	req.Header.Set("spanId", spanId)
	req.WithContext(reqOpts.ctx)
	defer req.Body.Close()
	client := &http.Client{
		Timeout: reqOpts.timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if time.Since(startTime) > 3*time.Second {
		log.New(reqOpts.ctx).Warn("slow_request", "url", url, "costTime(ms)", time.Now().Sub(startTime)/time.Millisecond,
			"traceId", reqOpts.ctx.Value("traceId"))
	} else {
		log.New(reqOpts.ctx).Debug("fast_request", "url", url, "costTime(ms)", time.Now().Sub(startTime)/time.Millisecond,
			"traceId", reqOpts.ctx.Value("traceId"))
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, data, errors.New(fmt.Sprintf("status code is not 200. code:[%d]", resp.StatusCode))
	}
	return resp.StatusCode, data, nil
}

func Get(ctx context.Context, url string, opts ...Option) (ResultCode int, body []byte, err error) {
	opts = append(opts, WithContext(ctx))
	return Request("GET", url, opts...)
}

func Post(ctx context.Context, data []byte, url string, opts ...Option) (ResultCode int, body []byte, err error) {
	opts = append(opts,
		WithContext(ctx),
		WithHeaders(map[string]string{"Content-Type": "application/json"}),
		WithData(data))
	return Request("POST", url, opts...)
}
