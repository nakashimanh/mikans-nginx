package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/nakashimanh/mikans/mikanpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) Mikan(ctx context.Context, req *mikanpb.MikanRequest) (*mikanpb.MikanResponse, error) {
	fmt.Printf("Mikan function was invoked with %v\n", req)
	name := req.GetMikan().GetName()
	kind := req.GetMikan().GetKind()
	quality := req.GetMikan().GetQuality()
	result := "Response Mikan= " + name + " kind= " + kind + " quality= " + strconv.FormatInt(quality, 10)
	res := &mikanpb.MikanResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Starting Mikan Service")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	mikanpb.RegisterMikanServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
