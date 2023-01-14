package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"schedul_rule/config"
	"schedul_rule/internal/handler"
	"schedul_rule/internal/repository"
	pb "schedul_rule/internal/service"
)

const (
	port = ":50053"
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
	pb.RegisterSchedulRuleServiceServer(s, handler.NewStaffService())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
