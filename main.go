package main

import (
	"log"

	"github.com/career-rizkiramdhani/microservice-auth-service/internal/server"
)

func main() {
	grpcPort := server.GetEnv("GRPC_PORT", "9090")
	httpPort := server.GetEnv("HTTP_PORT", "8080")

	srv := server.NewServer(grpcPort, httpPort)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
