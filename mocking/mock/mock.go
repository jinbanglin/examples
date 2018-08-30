package mock

import (
	"context"

	proto "github.com/jinbanglin/examples/helloworld/proto"
	"github.com/jinbanglin/go-micro/client"
)

type mockGreeterService struct {
}

func (m *mockGreeterService) Hello(ctx context.Context, req *proto.HelloRequest, opts ...client.CallOption) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Greeting: "Hello " + req.Name,
	}, nil
}

func NewGreeterService() proto.GreeterService {
	return new(mockGreeterService)
}
