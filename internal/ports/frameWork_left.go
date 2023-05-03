package ports

import (
	"context"
	"hex/internal/adapters/framework/left/grpc/pb"
)

type GRPCPorts interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationParameters) *pb.Answer
	GetSubtraction(ctx context.Context, req *pb.OperationParameters) *pb.Answer
	GetMultiplication(ctx context.Context, req *pb.OperationParameters) *pb.Answer
	GetDivision(ctx context.Context, req *pb.OperationParameters) *pb.Answer
}
