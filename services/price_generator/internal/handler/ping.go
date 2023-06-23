package handler

import (
	"context"

	pb "github.com/bearatol/interview_golang_task/proto/price_generator"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) Ping(ctx context.Context, in *emptypb.Empty) (*pb.Pong, error) {
	return &pb.Pong{
		Pong: "pong",
	}, nil
}
