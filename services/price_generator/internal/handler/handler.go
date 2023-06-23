package handler

import (
	"context"

	pb "github.com/bearatol/interview_golang_task/proto/price_generator"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	ctx           context.Context
	log           *zap.SugaredLogger
	servPriceFile ServicePriceFile
	pb.UnimplementedPriceGeneratorServer
}

func NewHandler(ctx context.Context, log *zap.SugaredLogger, serv ServicePriceFile) *Handler {
	return &Handler{
		ctx:           ctx,
		log:           log,
		servPriceFile: serv,
	}
}

func (h *Handler) errResp(err error) error {
	h.log.Error(err)
	return status.Errorf(codes.FailedPrecondition, "error: [%v]", err)
}
