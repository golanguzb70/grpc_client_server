package gprcclient

import (
	pb "github.com/golanguzb70/grpc_client_server/client/genproto/post_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	PostServiceClient() pb.PostServiceClient
}

type grpcClient struct {
	postService pb.PostServiceClient
}

func New() (ServiceManager, error) {
	postServiceConnection, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClient{
		postService: pb.NewPostServiceClient(postServiceConnection),
	}, nil
}

func (c *grpcClient) PostServiceClient() pb.PostServiceClient {
	return c.postService
}
