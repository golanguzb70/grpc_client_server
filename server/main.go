package main

import (
	"log"
	"net"

	pb "github.com/golanguzb70/grpc_client_server/server/genproto/post_service"
	"github.com/golanguzb70/grpc_client_server/server/service"
	"google.golang.org/grpc"
)

func main() {
	postService := service.NewPostService()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Printf("Error while listening: %v\n", err)
		return
	}

	server := grpc.NewServer()
	pb.RegisterPostServiceServer(server, postService)

	log.Println("gRPC server is running on port :9000")

	err = server.Serve(lis)
	if err != nil {
		log.Println("Error while running gRPC server", err)
	}
}
