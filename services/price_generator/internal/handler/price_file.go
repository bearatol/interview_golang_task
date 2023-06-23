package handler

import (
	"context"
	"fmt"

	pb "github.com/bearatol/interview_golang_task/proto/price_generator"
	"github.com/bearatol/interview_golang_task/sevices/price_generator/internal/mapping"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

const MaxBarcodesCount = 100

type ServicePriceFile interface {
	SetFile(barcode, title string, cost int32) error
	GetFilesByBarcodes(barcodes []string) ([]*mapping.PriceFile, error)
	DeleteFilesByBarcodes(barcodes []string) error
}

func (h *Handler) Set(ctx context.Context, in *pb.PriceFileSetRequest) (*empty.Empty, error) {
	if len(in.Barcode) == 0 {
		return &emptypb.Empty{}, h.errResp(fmt.Errorf("incorect inner data: barcode"))
	}
	if len(in.Title) == 0 {
		return &emptypb.Empty{}, h.errResp(fmt.Errorf("incorect inner data: title"))
	}
	if in.Cost == 0 {
		return &emptypb.Empty{}, h.errResp(fmt.Errorf("incorect inner data: cost"))
	}

	return &empty.Empty{}, h.servPriceFile.SetFile(in.Barcode, in.Title, in.Cost)
}

func (h *Handler) Get(ctx context.Context, in *pb.PriceFilesRequest) (*pb.PriceFilesResponse, error) {
	if len(in.Barcodes) > MaxBarcodesCount {
		return &pb.PriceFilesResponse{}, h.errResp(fmt.Errorf("too many barcodes, max: %d", MaxBarcodesCount))
	}

	res, err := h.servPriceFile.GetFilesByBarcodes(in.Barcodes)
	if err != nil {
		return &pb.PriceFilesResponse{}, h.errResp(err)
	}

	files := make([]*pb.File, len(res))
	for k, file := range res {
		files[k].Name = file.Name
		files[k].Content = file.Content
	}

	return &pb.PriceFilesResponse{Files: files}, nil
}

func (h *Handler) Delete(ctx context.Context, in *pb.PriceFilesRequest) (*empty.Empty, error) {
	if len(in.Barcodes) > MaxBarcodesCount {
		return &empty.Empty{}, h.errResp(fmt.Errorf("too many barcodes, max: %d", MaxBarcodesCount))
	}

	return &empty.Empty{}, h.servPriceFile.DeleteFilesByBarcodes(in.Barcodes)
}
