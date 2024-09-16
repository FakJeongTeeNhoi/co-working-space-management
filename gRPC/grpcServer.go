package gRPC

import (
	"fmt"
	"log"
	"net"

	pb "github.com/FakJeongTeeNhoi/co-working-space-management/generated/space"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type SpaceServer struct {
	pb.UnimplementedSpaceServiceServer
	db *gorm.DB
}

func NewServer(db *gorm.DB) *SpaceServer {
	return &SpaceServer{db: db}
}

func StartGRPCServer(db *gorm.DB) {
	grpcServer := grpc.NewServer()
	pb.RegisterSpaceServiceServer(grpcServer, NewServer(db))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
