package rpc

import (
	"hex/internal/adapters/framework/left/grpc/pb"
	"hex/internal/ports"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}
func (grpcA Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	arithmeticServiceServer := grpcA
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
