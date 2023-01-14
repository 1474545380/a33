package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"schedu_date/config"
	"schedu_date/internal/handler"
	"schedu_date/internal/repository"
	pb "schedu_date/internal/service"
)

const (
	port = ":50054"
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
	pb.RegisterSchedulDateServiceServer(s, handler.NewStaffService())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
