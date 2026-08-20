package grpc

import (
	"context"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

type clientStream struct {
	callInfo *callInfo
	cc       *ClientConn

	desc *StreamDesc

	codec baseCodec
	cp    Compressor
	comp  encoding.Compressor

	cancel context.CancelFunc

	sentHeader bool

	onFinishOnce sync.Once
	onFinish     func(error)

	mu       sync.Mutex
	finished bool
	attempt  *clientAttempt
	ctx      context.Context
}

func (cs *clientStream) finish(err error) {
	if cs.onFinish != nil {
		cs.onFinishOnce.Do(func() {
			cs.onFinish(err)
		})
	}
	cs.mu.Lock()
	if cs.finished {
		cs.mu.Unlock()
		return
	}
	cs.finished = true
	cs.mu.Unlock()
	if cs.cancel != nil {
		cs.cancel()
	}
}

func (cs *clientStream) Header() (metadata.MD, error) {
	return nil, nil
}

func (cs *clientStream) Trailer() metadata.MD {
	return nil
}

func (cs *clientStream) CloseSend() error {
	return nil
}

func (cs *clientStream) Context() context.Context {
	return cs.ctx
}

func (cs *clientStream) SendMsg(m interface{}) error {
	return nil
}

func (cs *clientStream) RecvMsg(m interface{}) error {
	return nil
}
