package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/nosarthur/telep/messages"
	"google.golang.org/grpc"
)

type server struct{}

const port = ":9527"

func (s *server) Run(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	fmt.Println("in Run")
	input := strings.Join(req.Cmds, " ")
	reply := &pb.Reply{Stdout: input}
	return reply, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommandServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
