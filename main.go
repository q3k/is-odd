package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/q3k/is-odd/proto"
)

var (
	flagListen string
)

type server struct{}

func main() {
	flag.StringVar(&flagListen, "listen", "0.0.0.0:2137", "Address to listen at")
	flag.Parse()

	lis, err := net.Listen("tcp", flagListen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterIsOddServer(s, &server{})

	log.Printf("listening at %v", flagListen)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) IsOdd(ctx context.Context, req *pb.IsOddRequest) (*pb.IsOddResponse, error) {
	res := &pb.IsOddResponse{
		Result: pb.IsOddResponse_RESULT_NON_ODD,
	}
	if req.Number%2 != 0 {
		res.Result = pb.IsOddResponse_RESULT_ODD
	}
	return res, nil
}
