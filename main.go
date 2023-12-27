package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/anyboards/proto/gen/go/boards"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type boardsServer struct {
	boards.UnimplementedBoardsServer
}

func (*boardsServer) Create(context.Context, *boards.CreateBoardRequest) (*boards.CreateBoardResponse, error) {
	return &boards.CreateBoardResponse{Board: &boards.Board{Id: "123", Name: "test"}}, nil
}

func (*boardsServer) ListBoards(context.Context, *emptypb.Empty) (*boards.ListBoardResponse, error) {
	return &boards.ListBoardResponse{Item: []*boards.ListBoardResponseItem{
		{Id: "123", Name: "test 1"},
		{Id: "124", Name: "test 2"},
	}}, nil
}

func newServer() *boardsServer {
	s := &boardsServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 4444))
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
