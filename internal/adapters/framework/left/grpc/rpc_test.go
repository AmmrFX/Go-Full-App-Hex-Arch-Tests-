package rpc

import (
	"context"
	"hex/internal/adapters/framework/left/grpc/pb"
	"log"
	"net"
	"os"
	"testing"

	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.ort/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort

	var gRPCAdapter ports.GRPCPorts
	dbaseDriver := os.Getenv("DB_DRIVER")
	dbsourceName := os.Getenv("DS_Name")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dbsourceName)

	if err != nil {
		log.Fatalf("error creating db adapter %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()
	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)
	gRPCAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, gRPCAdapter)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {

			log.Fatalf("failed to serve: %v", err)
		}
	}()

}
func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure)
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	return conn
}
func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}
	answer, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("could not get addition: %v", err)
	}
	require.Equal(t, answer.Value, int32(2))
}
func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}
	answer, err := client.GetSubtraction(ctx, params)
	if err != nil {
		t.Fatalf("could not get addition: %v", err)
	}
	require.Equal(t, answer.Value, int32(0))
}
func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}
	answer, err := client.GetMultiplication(ctx, params)
	if err != nil {
		t.Fatalf("could not get addition: %v", err)
	}
	require.Equal(t, answer.Value, int32(1))
}
func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}
	answer, err := client.GetDivision(ctx, params)
	if err != nil {
		t.Fatalf("could not get addition: %v", err)
	}
	require.Equal(t, answer.Value, int32(1))
}
