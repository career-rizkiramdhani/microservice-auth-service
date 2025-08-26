package service

import (
	"context"

	pb "github.com/career-rizkiramdhani/microservice-auth-service/pkg/protobuf"
)

// AuthService implements all business logic for authentication
type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

// NewAuthService creates a new authentication service
func NewAuthService() *AuthService {
	return &AuthService{}
}

// Health implements the Health RPC method
func (s *AuthService) Health(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	// Business logic for health check
	response := &pb.HealthResponse{
		Status:  "OK",
		Message: "Auth service is running and healthy",
	}

	return response, nil
}
