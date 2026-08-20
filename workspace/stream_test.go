package grpc

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestClientStreamOnFinishOnce(t *testing.T) {
	var count int32
	cs := &clientStream{
		desc: &StreamDesc{},
		cc:   &ClientConn{},
		onFinish: func(err error) {
			atomic.AddInt32(&count, 1)
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cs.finish(nil)
		}()
	}
	wg.Wait()

	if count != 1 {
		t.Errorf("onFinish called %d times, want 1", count)
	}
}
