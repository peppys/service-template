package main

import (
	"context"
	"fmt"
	"github.com/peppys/service-template/gen/go/proto"
	"github.com/peppys/service-template/internal/config"
	"github.com/peppys/service-template/internal/grpcservers"
	"github.com/peppys/service-template/internal/grpcservers/interceptors"
	"github.com/peppys/service-template/internal/repositories"
	"github.com/peppys/service-template/internal/services"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var appConfig *config.AppConfig

func main() {
	appConfig = config.NewAppConfig()
	db := initDB()

	// repositories
	todoRepository := repositories.NewTodoRepository(db)
	userRepository := repositories.NewUserRepository(db)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(db)

	// services
	authservice := services.NewAuthService(userRepository, refreshTokenRepository)

	// grpc servers
	server := grpc.NewServer(buildServerOpts(authservice)...)
	reflection.Register(server)

	healthserver := grpcservers.NewHealthGrpcServer(db)
	proto.RegisterAuthServiceServer(server, grpcservers.NewAuthGrpcServer(authservice))
	proto.RegisterTodoServiceServer(server, grpcservers.NewTodoGrpcServer(todoRepository))
	proto.RegisterHealthServiceServer(server, healthserver)
	grpc_health_v1.RegisterHealthServer(server, healthserver)

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := http.NewServeMux()
	gateway := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{UseProtoNames: true},
	}))
	if err = proto.RegisterAuthServiceHandler(context.Background(), gateway, conn); err != nil {
		log.Fatalln("Failed to register auth gateway:", err)
	}
	if err = proto.RegisterTodoServiceHandler(context.Background(), gateway, conn); err != nil {
		log.Fatalln("Failed to register todo gateway:", err)
	}
	if err = proto.RegisterHealthServiceHandler(context.Background(), gateway, conn); err != nil {
		log.Fatalln("Failed to register health gateway:", err)
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

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", appConfig.DB.User, appConfig.DB.Pass, appConfig.DB.Host, appConfig.DB.Port, appConfig.DB.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Errorf("error initializing db connection: %v", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(fmt.Errorf("error connecting to db: %v", err))
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func openapiFileHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			log.Printf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		log.Printf("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/openapiv2/")
		p = path.Join("static/openapiv2/", p)
		http.ServeFile(w, r, p)
	})
}

func buildServerOpts(authService *services.AuthService) []grpc.ServerOption {
	var opts []grpc.ServerOption

	opts = append(opts, grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			interceptors.Logging(),
			interceptors.Authorization(authService),
		),
	))

	return opts
}

func swaggerUIHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/swagger-ui/")
		p = path.Join("static/swagger-ui/", p)
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
