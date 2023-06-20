package document_generator

import (
	"context"
	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/document_generator_api"
	"google.golang.org/grpc"
)

func (c *Client) Generate(ctx context.Context, req *pb.GenerateLetterRequest,
	opts ...grpc.CallOption) (*pb.GenerateLetterResponse, error) {
	return c.grpcClient.Generate(ctx, req)
}
