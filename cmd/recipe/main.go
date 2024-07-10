package main

import (
	"context"
	"log"
	"net"

	recipe "github.com/warmcoldz/microservices/api/recipe"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	recipe.UnimplementedRecipeServiceServer
}

func (s *server) GetRecipe(ctx context.Context, in *recipe.RecipeRequest) (*recipe.RecipeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "GetRecipe is not implemented")
}

func RunServer() {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	recipe.RegisterRecipeServiceServer(s, &server{})

	log.Println("Server is running on port :8080")

	if err := s.Serve(conn); err != nil {
		log.Fatalf("Failed start GRPC server: %v", err)
	}
}

func main() {
	RunServer()
}
