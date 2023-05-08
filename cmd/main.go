package main

import (
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"

	gRBC "hex/internal/adapters/framework/left/grpc"
	"log"
	"os"
	//"hex/internal/ports"
)

func main() {
	var err error
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
	gRPCAdapter = gRBC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
	// ports
	// var core ports.ArithmeticPort
	// core = arithmetic.NewAdapter()
	// arithAdapter := arithmetic.NewAdapter()
	// result, err := arithAdapter.Addition(1, 3)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)
}
