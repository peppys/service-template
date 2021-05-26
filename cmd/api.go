package main

import (
	"context"
	"github.com/peppys/service-template/gen/go/proto"
	"github.com/peppys/service-template/internal/grpcserver"
	"github.com/peppys/service-template/internal/repositories"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	runtime.NewServeMux()

	todorepository := &repositories.TodoRepository{}
	todo.RegisterTodoServiceServer(server, grpcserver.NewTodoGrpcServer(todorepository))

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gateway := runtime.NewServeMux()
	if err = todo.RegisterTodoServiceHandler(context.Background(), gateway, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	ok := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(httpGrpcRouter(server, gateway), &http2.Server{}),
	}
	log.Println("Listing on port :8080...")
	log.Fatal(ok.ListenAndServe())
}

func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}
