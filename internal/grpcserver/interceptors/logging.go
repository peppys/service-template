package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		h, err := handler(ctx, req)

		log.Printf("Request - Method:%s\tDuration:%f seconds\tError:%v", info.FullMethod, time.Since(start).Seconds(), err)

		return h, err
	}
}
