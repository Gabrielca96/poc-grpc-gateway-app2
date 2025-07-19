package main

import (
	"context"
	"log"
	"net"

	pb "grpc-service-b/gen/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProductServiceServer
}

func (s *server) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{Id: req.Id, Name: "Test Product"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &server{})
	log.Println("Product gRPC service running on :50052")
	grpcServer.Serve(lis)
}
