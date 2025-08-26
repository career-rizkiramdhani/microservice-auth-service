package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/career-rizkiramdhani/microservice-auth-service/internal/service"
	pb "github.com/career-rizkiramdhani/microservice-auth-service/pkg/protobuf"
)

// Server manages both gRPC and HTTP servers
type Server struct {
	grpcPort    string
	httpPort    string
	authService *service.AuthService
}

// NewServer creates a new server instance
func NewServer(grpcPort, httpPort string) *Server {
	return &Server{
		grpcPort:    grpcPort,
		httpPort:    httpPort,
		authService: service.NewAuthService(),
	}
}

// Start begins both gRPC and HTTP servers concurrently
func (s *Server) Start() error {
	var wg sync.WaitGroup

	// Channel to listen for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start gRPC server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startGRPCServer(quit); err != nil {
			log.Printf("gRPC server error: %v", err)
		}
	}()

	// Start HTTP gateway server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startHTTPGateway(quit); err != nil {
			log.Printf("HTTP gateway error: %v", err)
		}
	}()

	fmt.Printf("Auth Service started:\n")
	fmt.Printf("gRPC Server: localhost:%s\n", s.grpcPort)
	fmt.Printf("HTTP Gateway: http://localhost:%s\n", s.httpPort)
	fmt.Printf("Health Check: http://localhost:%s/api/v1/health\n", s.httpPort)
	fmt.Printf("Swagger UI: Upload ./docs/auth_api.swagger.json to https://editor.swagger.io\n")
	fmt.Printf("Press Ctrl+C to stop\n\n")

	// Wait for interrupt signal
	<-quit
	fmt.Println("\n Shutting down servers...")

	wg.Wait()
	fmt.Println("Servers stopped gracefully")

	return nil
}

// startGRPCServer starts the gRPC server
func (s *Server) startGRPCServer(quit chan os.Signal) error {
	lis, err := net.Listen("tcp", ":"+s.grpcPort)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %v", s.grpcPort, err)
	}

	grpcServer := grpc.NewServer()

	// Register auth service
	pb.RegisterAuthServiceServer(grpcServer, s.authService)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("gRPC server serve error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-quit

	// Graceful shutdown
	grpcServer.GracefulStop()
	return nil
}

// startHTTPGateway starts the HTTP gateway server
func (s *Server) startHTTPGateway(quit chan os.Signal) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create gRPC connection to our own server
	conn, err := grpc.DialContext(
		ctx,
		"localhost:"+s.grpcPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create HTTP gateway mux
	mux := runtime.NewServeMux()

	// Register auth service with the gateway
	err = pb.RegisterAuthServiceHandler(ctx, mux, conn)
	if err != nil {
		return fmt.Errorf("failed to register auth service handler: %v", err)
	}

	// Add custom routes for documentation
	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)

	// Serve swagger documentation
	httpMux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/auth_api.swagger.json")
	})

	server := &http.Server{
		Addr:    ":" + s.httpPort,
		Handler: httpMux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP gateway serve error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-quit

	// Graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(shutdownCtx)
}

// GetEnv gets environment variable with default value
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
