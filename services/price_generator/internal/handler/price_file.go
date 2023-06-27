package handler

import (
	"context"
	"fmt"

	pb "github.com/bearatol/interview_golang_task/proto/price_generator"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

//go:generate mockgen -source=price_file.go -destination=mocks/price_file_mock.go

type ServicePriceFile interface {
	SetFile(fileName, barcode, title string, cost int32) error
	GetFileByName(fileName string) ([]byte, error)
	DeleteFileByName(fileName string) error
}

func (h *Handler) Set(ctx context.Context, in *pb.PriceFileSetRequest) (*empty.Empty, error) {
	if len(in.FileName) == 0 {
		return &emptypb.Empty{}, h.errResp(fmt.Errorf("incorect inner data: file name"))
	}
	if len(in.Barcode) == 0 {
		return &emptypb.Empty{}, h.errResp(fmt.Errorf("incorect inner data: barcode"))
	}
	if len(in.Title) == 0 {
		return &emptypb.Empty{}, h.errResp(fmt.Errorf("incorect inner data: title"))
	}
	if in.Cost == 0 {
		return &emptypb.Empty{}, h.errResp(fmt.Errorf("incorect inner data: cost"))
	}

	return &empty.Empty{}, h.servPriceFile.SetFile(in.FileName, in.Barcode, in.Title, in.Cost)
}

func (h *Handler) Get(ctx context.Context, in *pb.PriceFileRequest) (*pb.PriceFileResponse, error) {
	res, err := h.servPriceFile.GetFileByName(in.FileName)
	if err != nil {
		return &pb.PriceFileResponse{}, h.errResp(err)
	}
	if res == nil {
		return &pb.PriceFileResponse{}, h.errResp(fmt.Errorf("cannot get price file"))
	}

	return &pb.PriceFileResponse{File: res}, nil
}

func (h *Handler) Delete(ctx context.Context, in *pb.PriceFileRequest) (*empty.Empty, error) {
	return &empty.Empty{}, h.servPriceFile.DeleteFileByName(in.FileName)
}
