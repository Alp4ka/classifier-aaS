package grpc

import "context"

func mergeContext(a, b context.Context) (context.Context, context.CancelFunc) {
	mctx, mcancel := context.WithCancel(a) // will cancel if `a` cancels

	go func() {
		select {
		case <-mctx.Done(): // don't leak go-routine on clean gRPC run
		case <-b.Done():
			mcancel() // b canceled, so cancel mctx
		}
	}()

	return mctx, mcancel
}

func stringPointer(s string) *string {
	t := s
	return &t
}
