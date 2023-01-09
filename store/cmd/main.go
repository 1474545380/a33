package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"store/config"
	"store/internal/handler"
	"store/internal/repository"
	pb "store/internal/service"
)

const (
	port = ":50051"
)

func main() {
	config.InitConfig()
	repository.InitDB()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	defer s.Stop()
	pb.RegisterStoreServiceServer(s, handler.NewStoreService())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
