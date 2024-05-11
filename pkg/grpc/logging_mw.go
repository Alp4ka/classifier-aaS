package grpc

import (
	"context"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func ServerStreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if isReflectionStream(info) {
			return handler(srv, ss)
		}

		ctx := ss.Context()
		startTime := time.Now()

		ctx = field.WithContextFields(ctx, getStreamFields(ctx, info)...)

		wrappedStream := &serverStream{
			ServerStream: ss,
			ctx:          ctx,
		}

		err := handler(srv, wrappedStream)

		st, _ := status.FromError(err)
		fields := []field.Field{
			field.String("grpc.code", st.Code().String()),
			field.String("grpc.call_duration", time.Since(startTime).String()),
		}
		if st.Code() == codes.OK {
			mlogger.L(ctx).Info("Stream completed", fields...)
		} else {
			fields = append(fields, field.String("grpc.error", err.Error()))
			mlogger.L(ctx).Error("Stream completed", fields...)
		}

		return err
	}
}

func getStreamFields(ctx context.Context, info *grpc.StreamServerInfo) []field.Field {
	ret := make([]field.Field, 0, 2)

	ret = append(ret, field.String("grpc.method", info.FullMethod))

	if deadline, ok := ctx.Deadline(); ok {
		ret = append(ret, field.String("grpc.request_deadline", deadline.String()))
	}

	return ret
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss *serverStream) Context() context.Context {
	return ss.ctx
}

func (ss *serverStream) SendMsg(msg interface{}) error {
	mlogger.L(ss.ctx).Info("SendMsg", field.Any("response", msg))

	return ss.ServerStream.SendMsg(msg)
}

func (ss *serverStream) RecvMsg(msg interface{}) error {
	err := ss.ServerStream.RecvMsg(msg)
	mlogger.L(ss.ctx).Info("RecvMsg", field.Any("request", msg))

	return err
}

func isReflectionStream(info *grpc.StreamServerInfo) bool {
	return info.FullMethod == "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if isReflectionUnary(info) {
			return handler(ctx, req)
		}

		startTime := time.Now()
		ctx = field.WithContextFields(ctx, getUnaryFields(ctx, info)...)

		mlogger.L(ctx).Info("Request!", field.Any("request", req))
		ret, err := handler(ctx, req)

		st, _ := status.FromError(err)
		fields := []field.Field{
			field.String("grpc.code", st.Code().String()),
			field.String("grpc.call_duration", time.Since(startTime).String()),
		}
		if st.Code() == codes.OK {
			mlogger.L(ctx).Info("Response!", fields...)
		} else {
			fields = append(fields, field.String("grpc.error", err.Error()))
			mlogger.L(ctx).Error("Response!", fields...)
		}

		return ret, err
	}
}

func getUnaryFields(ctx context.Context, info *grpc.UnaryServerInfo) []field.Field {
	ret := make([]field.Field, 0, 2)

	ret = append(ret, field.String("grpc.method", info.FullMethod))

	if deadline, ok := ctx.Deadline(); ok {
		ret = append(ret, field.String("grpc.request_deadline", deadline.String()))
	}

	return ret
}

func isReflectionUnary(info *grpc.UnaryServerInfo) bool {
	return info.FullMethod == "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
}
