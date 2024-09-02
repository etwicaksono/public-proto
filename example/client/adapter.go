package client

import (
	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Adapter is a wrapper around the generated protobuf client.
type Adapter struct {
	example.ExampleServiceClient
	ClientConn *grpc.ClientConn
}

// NewExampleServiceGrpcClient returns a new gRPC client for the example service.
func NewExampleServiceGrpcClient(address string) *Adapter {
	var err error
	adapter := &Adapter{}
	// Set up a connection to the gRPC server.
	adapter.ClientConn, err = grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	adapter.ExampleServiceClient = example.NewExampleServiceClient(adapter.ClientConn)
	return adapter
}
