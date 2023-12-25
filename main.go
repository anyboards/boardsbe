package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/anyboards/proto/gen/go/boards"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type boardsServer struct {
	boards.UnimplementedBoardsServer
}

func (*boardsServer) Create(context.Context, *emptypb.Empty) (*boards.CreateBoardResponse, error) {
	return &boards.CreateBoardResponse{Id: "123"}, nil
}

func (*boardsServer) ListBoards(ctx context.Context, _ *emptypb.Empty) (*boards.ListBoardsResponse, error) {
	print("LB")
	// md, ok := metadata.FromIncomingContext(ctx)
	// if ok {
	// 	vv := md.Get("Authorization")
	// 	// vv = "Bearer ..."
	// 	print(vv)
	// }

	return &boards.ListBoardsResponse{Boards: []*boards.ListBoardResponse{
		{
			Id:        "123",
			CreatedAt: timestamppb.New(time.Now()),
		},
		{
			Id:        "4567",
			CreatedAt: timestamppb.New(time.Now()),
		},
	}}, nil
}

func newServer() *boardsServer {
	s := &boardsServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 2222))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening!")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	boards.RegisterBoardsServer(grpcServer, newServer())
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
