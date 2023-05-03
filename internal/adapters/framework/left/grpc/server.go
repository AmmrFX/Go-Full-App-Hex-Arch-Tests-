package rpc

import (
	"hex/internal/adapters/framework/left/grpc/pb"
	"hex/internal/port"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Adapter struct {
	api port.APIPort
}

func NewAdapter(api port.APIPort) *Adapter {
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
