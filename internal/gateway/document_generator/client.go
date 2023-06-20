package document_generator

import (
	"context"
	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/document_generator_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GrpcClient interface {
	Generate(ctx context.Context, in *pb.GenerateLetterRequest, opts ...grpc.CallOption) (*pb.GenerateLetterResponse, error)
}

type Client struct {
	grpcClient GrpcClient
}

func New(ctx context.Context, grpcHost string) (*Client, error) {
	grpcClient, err := newClient(ctx, grpcHost)
	if err != nil {
		return nil, err
	}

	return &Client{grpcClient: grpcClient}, nil
}

func newClient(_ context.Context, grpcHost string) (GrpcClient, error) {
	grpcConn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("can't dial grpc due to:", err)
	}

	return pb.NewGenerateLecInitialLetterClient(grpcConn), nil
}
