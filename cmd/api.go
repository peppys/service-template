package main

import (
	"context"
	"github.com/peppys/service-template/gen/go/proto"
	"github.com/peppys/service-template/internal/grpcserver"
	"github.com/peppys/service-template/internal/repositories"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc/reflection"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	reflection.Register(server)

	todorepository := &repositories.TodoRepository{}
	proto.RegisterTodoServiceServer(server, grpcserver.NewTodoGrpcServer(todorepository))

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := http.NewServeMux()
	gateway := runtime.NewServeMux()
	if err = proto.RegisterTodoServiceHandler(context.Background(), gateway, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	mux.Handle("/", gateway)
	mux.Handle("/openapiv2/", openapiFileHandler())
	mux.Handle("/swagger-ui/", swaggerUIHandler())

	ok := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(httpGrpcRouter(server, mux), &http2.Server{}),
	}
	log.Println("Listing on port :8080...")
	log.Fatal(ok.ListenAndServe())
}

func openapiFileHandler() http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			log.Printf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		log.Printf("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/openapiv2/")
		p = path.Join("gen/openapiv2/", p)
		http.ServeFile(w, r, p)
	})
}

func swaggerUIHandler() http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/swagger-ui/")
		p = path.Join("swagger-ui/", p)
		http.ServeFile(w, r, p)
	})
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
