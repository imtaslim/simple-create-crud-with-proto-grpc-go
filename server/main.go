package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	cpb "grpc/proto/category"
	"grpc/server/categories"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	grpcServer := grpc.NewServer()
	s := categories.Server{}

	cpb.RegisterCategoryServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to listen: %s", err)
	}


}