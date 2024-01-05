package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/anyboards/proto/gen/go/debug"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type srv struct {
	debug.UnimplementedDebugServer
}

func (*srv) Add(_ context.Context, r *debug.AddRequest) (*debug.AddResponse, error) {
	return &debug.AddResponse{Result: r.A + r.B}, nil
}

func (*srv) Greet(_ context.Context, r *debug.GreetRequest) (*debug.GreetResponse, error) {
	return &debug.GreetResponse{Greeting: "Hello " + r.Name + "!"}, nil
}

func (*srv) Stream(r *debug.StreamRequest, dss debug.Debug_StreamServer) error {
	ticker := time.NewTicker(r.GetInterval().AsDuration())
	num := 1

	for {
		select {
		case <-ticker.C:
			dss.Send(&debug.StreamResponse{MessageNum: fmt.Sprintf("%d", num)})
			num += 1
		case <-dss.Context().Done():
			return dss.Context().Err()
		}
	}

	return nil
}

func newServer() *srv {
	s := &srv{}
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
	debug.RegisterDebugServer(grpcServer, newServer())
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
